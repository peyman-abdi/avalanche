package interfaces

type Module interface {
	Title() string
	Description() string
	Version() string

	Migrations() []Migratable
	Routes() []*Route
	MiddleWares() map[string]RequestHandler
	GroupsHandlers() map[string]RequestHandler

	Activated() bool
	Installed() bool
	Deactivated()
	Purged()
}

type ModuleManager interface {
	LoadModules(Services)
	List() []Module
	Activated() []Module
	Deactivated() []Module
	Installed() []Module
	NotInstalled() []Module

	IsActive(Module) bool
	IsInstalled(Module) bool
	Install(Module) error
	Activate(Module) error
	Purge(Module) error
	Deactivate(Module) error
}