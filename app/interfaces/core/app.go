package core

type Application interface {
	ConfigPath(path string) string
	StoragePath(path string) string
	RootPath(path string) string
	ModulesPath(path string) string
	ResourcesPath(path string) string
	InitAvalanchePlugins(path string, services Services) []AvalanchePlugin

	Version() string
	Build() string
	BuildCode() int
	Platform() string
	Variant() string
	BuildTime() string
	IsDebugMode() bool
}
