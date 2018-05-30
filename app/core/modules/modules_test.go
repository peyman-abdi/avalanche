package modules

import (
	"testing"
	"github.com/peyman-abdi/avalanche/app/interfaces"
	"os"
	"github.com/peyman-abdi/testil"
	"github.com/peyman-abdi/avalanche/app/core/config"
	"github.com/peyman-abdi/avalanche/app/core/logger"
	"github.com/peyman-abdi/avalanche/app/core/database"
	application "github.com/peyman-abdi/avalanche/app/core/app"
	"github.com/peyman-abdi/avalanche/app/core/router"
)

var app interfaces.Application
var conf interfaces.Config
var log interfaces.Logger
var repo interfaces.Repository
var mig interfaces.Migrator
var mm interfaces.ModuleManager
var r interfaces.Router
var s = new(ServicesMock)

type ServicesMock struct {
}
func (s *ServicesMock) Repository() interfaces.Repository { return repo }
func (s *ServicesMock) Migrator() interfaces.Migrator { return mig }
func (s *ServicesMock) Localization() interfaces.Localization { return nil }
func (s *ServicesMock) Config() interfaces.Config { return conf }
func (s *ServicesMock) Logger() interfaces.Logger { return log }
func (s *ServicesMock) Modules() interfaces.ModuleManager { return mm }
func (s *ServicesMock) App() interfaces.Application { return app }
func (s *ServicesMock) Router() interfaces.Router { return r }

var envs = map[string]string {
}
var configs = map[string]interface{} {
	"app.hjson": map[string]interface{} {
		"debug": true,
	},
	"database.hjson": map[string]interface{} {
		"app": "sqlite3",
		"runtime": map[string]interface{} {
			"migrations": "migrations",
			"connection": "sqlite3",
		},
		"connections": map[string]interface{} {
			"sqlite3": map[string]interface{} {
				"driver": "sqlite3",
				"file": "storage(\"test.db\")",
			},
		},
	},
}
func init()  {
	app = application.Initialize(0)
	os.MkdirAll(app.StoragePath(""), 0700)

	testil.CreateConfigFiles(app, configs)

	conf = config.Initialize(app)
	log = logger.Initialize(conf)
	log.LoadConsole()

	repo, mig = database.Initialize(conf, log)

	r = router.Initialize(conf, log)

	mm = Initialize(conf, mig)
	mm.LoadModules(s)
}

func TestModuleStatus(t *testing.T) {
	testModule := new(testil.TestMigrationModule)
	mma := mm.(*moduleManagerImpl)
	mma.AvailableModules = append(mma.AvailableModules, testModule)

	mm.Install(testModule)
	mm.Activate(testModule)

	var err error
	var migrations []*database.MigrationModel
	if err = repo.Query(&database.MigrationModel{}).Get(&migrations); err != nil {
		t.Error(err)
	}
	if len(migrations) != 1 {
		t.Errorf("Migrations not happend")
	}

	if !mm.IsInstalled(testModule) {
		t.Errorf("Module not installed!")
	}
	if !mm.IsActive(testModule) {
		t.Errorf("Module not active")
	}

	if c := len(mm.Activated()); c != 1 {
		t.Errorf("Activated list faled: %d", c)
	}

	mm.Deactivate(testModule)

	if mm.IsActive(testModule) {
		t.Errorf("Module not deactive")
	}

	if len(mm.Activated()) != 0 {
		t.Errorf("Activated module faled")
	}
	if len(mm.Deactivated()) != 1 {
		t.Errorf("Deactivated list faled")
	}
	if len(mm.Installed()) != 1 {
		t.Errorf("Installed list failed")
	}

	mm.Purge(testModule)
	if len(mm.NotInstalled()) != 1 {
		t.Errorf("Not installed list failed")
	}

	if err = repo.Query(&database.MigrationModel{}).Get(&migrations); err != nil {
		t.Error(err)
	}
	if len(migrations) != 0 {
		t.Errorf("Migrations not rolledbacked")
	}


	if err = repo.Insert(&testil.TestMigrationModel{MyTest:"text"}); err != nil {
		t.Error(err)
	}
}

