package app

import (
	"github.com/peyman-abdi/avalanche/app/interfaces/services"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

var (
	Version     string
	Code        string
	VersionCode = 1
	Platform    string
	Variant     string
	BuildTime   string
	BuildMode   string
)

type appImpl struct {
	appRoot string
}

func Initialize(roots int, buildMode string) services.Application {
	root, err := os.Executable()
	if err != nil {
		panic(err)
	}

	root = filepath.Dir(root)
	for i := roots; i > 0; i-- {
		root = filepath.Join(root, "..")
	}

	app := &appImpl{
		appRoot: root,
	}
	BuildMode = buildMode

	return app
}

func (a *appImpl) Version() string {
	return Version
}
func (a *appImpl) Build() string {
	return Code
}
func (a *appImpl) BuildCode() int {
	return VersionCode
}
func (a *appImpl) Platform() string {
	return Platform
}
func (a *appImpl) Variant() string {
	return Variant
}
func (a *appImpl) BuildTime() string {
	return BuildTime
}
func (a *appImpl) Mode() string {
	return BuildMode
}
func (a *appImpl) IsDebugMode() bool {
	return Variant == "DEBUG"
}

func (a *appImpl) StoragePath(path string) string {
	return filepath.Join(a.appRoot, "storage", path)
}

func (a *appImpl) ConfigPath(path string) string {
	return filepath.Join(a.appRoot, "config", path)
}

func (a *appImpl) RootPath(path string) string {
	return filepath.Join(a.appRoot, path)
}

func (a *appImpl) ModulesPath(path string) string {
	return filepath.Join(a.appRoot, "bin/platforms/", Platform, Variant, "modules", path)
}

func (a *appImpl) ResourcesPath(path string) string {
	return filepath.Join(a.appRoot, "resources", path)
}

func (a *appImpl) TemplatesPath(path string) string {
	return filepath.Join(a.appRoot, "resources/views/templates", path)
}

func (a *appImpl) InitAvalanchePlugins(path string, references services.Services) []services.AvalanchePlugin {
	var moduleFiles []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		moduleFiles = append(moduleFiles, path)
		return nil
	})

	var modules []services.AvalanchePlugin
	for _, moduleFile := range moduleFiles {
		if !strings.HasSuffix(moduleFile, ".so") {
			continue
		}

		pl, err := plugin.Open(moduleFile)
		if err != nil {
			panic(err)
		}

		pluginInstanceRef, err := pl.Lookup("PluginInstance")
		if err != nil {
			panic(err)
		}

		pluginInstance := *pluginInstanceRef.(*services.AvalanchePlugin)
		if pluginInstance.Initialize(references) {
			modules = append(modules, pluginInstance)
		}
	}

	return modules
}
