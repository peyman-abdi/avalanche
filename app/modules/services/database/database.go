package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/peyman-abdi/avalanche/app/interfaces/services"
	"time"
)

var (
	appConnection     *gorm.DB
	runtimeConnection *gorm.DB
	connections       = make(map[string]*gorm.DB)
	connectionPrefix  = make(map[*gorm.DB]string)

	migrationsManager *MigrationManager

	migrator    services.Migrator
	repoManager *RepositoryManager
	loggerRef   services.Logger
)

func Initialize(config services.Config, logger services.Logger) (services.Repository, services.Migrator) {
	loggerRef = logger
	appConnectionName := config.GetString("rdbms.app", "sqlite3")
	runtimeConnectionName := config.GetString("rdbms.runtime.connection", "sqlite3")

	gorm.LogFormatter = func(values ...interface{}) (messages []interface{}) {
		return values
	}

	connectionDefs := config.GetMap("rdbms.connections", map[string]interface{}{})
	for connName, connParams := range connectionDefs {
		connMap := connParams.(map[string]interface{})
		var connection *gorm.DB = nil
		if connName == appConnectionName || connName == runtimeConnectionName {
			connection = initDatabase(config, "rdbms.connections."+connName)
			if connName == appConnectionName {
				appConnection = connection
				connections["app"] = connection
			}
			if connName == runtimeConnectionName {
				runtimeConnection = connection
				connections["runtime"] = connection
			}
		}

		if connMap["active"] == true {
			if connection == nil {
				connection = initDatabase(config, "rdbms.connections."+connName)
			}

			connections[connName] = connection
		}
	}

	if config.GetBoolean("app.debug", false) {
		for _, conn := range connections {
			conn.LogMode(true)
		}
		appConnection.LogMode(true)
		runtimeConnection.LogMode(true)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if prefix, ok := connectionPrefix[db]; ok {
			return prefix + defaultTableName
		}
		return defaultTableName
	}

	repoManager = new(RepositoryManager)
	repoManager.log = logger

	migrationsManager = new(MigrationManager)
	migrationsManager.migrationsTableName = config.GetString("rdbms.runtime.migrations", "migrations")
	migrationsManager.log = logger

	migrator = migrationsManager.Connection("runtime")

	migrationsManager.setup(runtimeConnection)

	return repoManager, migrator
}
func Close() {
	appConnection.Close()
}

func initDatabase(config services.Config, keyPrefix string) *gorm.DB {
	var connection *gorm.DB
	if connectionDriver := config.GetString(keyPrefix+".driver", ""); connectionDriver != "" {
		switch connectionDriver {
		case "mysql":
			connection = openMySQL(config, keyPrefix)
		case "sqlite3":
			connection = openSqlite3(config, keyPrefix)
		}
	} else {
		panic("Unknown database connection: " + keyPrefix)
	}

	if config.IsSet(keyPrefix + ".options") {
		maxIdleConnections := config.GetInt(keyPrefix+".options.maxIdleConnections", 1)
		connection.DB().SetMaxIdleConns(maxIdleConnections)

		maxOpenConnections := config.GetInt(keyPrefix+".options.maxOpenConnections", 1)
		connection.DB().SetMaxOpenConns(maxOpenConnections)

		maxConnectionLifetime := config.Get(keyPrefix+".options.maxConnectionLifetime", time.Hour).(time.Duration)
		connection.DB().SetConnMaxLifetime(maxConnectionLifetime)

		if config.IsSet(keyPrefix + ".options.prefix") {
			connectionPrefix[connection] = config.GetString(keyPrefix+".options.prefix", "")
		}
	} else {
		connection.DB().SetMaxIdleConns(1)
		connection.DB().SetMaxOpenConns(1)
		connection.DB().SetConnMaxLifetime(time.Hour)
	}

	connection.SetLogger(gorm.Logger{
		LogWriter: new(AvalancheDBLogWriter),
	})

	return connection
}

func openSqlite3(config services.Config, configPath string) *gorm.DB {
	filename := config.GetString(configPath+".file", "")
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	return db
}

func openMySQL(config services.Config, configPath string) *gorm.DB {

	//engine := config.GetString(keyPrefix + ".options.engine", "InnoDB")
	//connection = appConnection.Set("gorm:table_options", "ENGINE=" + engine)
	return nil
}
