package config

import (
	"github.com/peyman-abdi/avalanche/app/modules/core/app"
	gConfig "github.com/peyman-abdi/conf"
	"time"
	"github.com/peyman-abdi/avalanche/app/interfaces/core"
)

/**
Storage Path: storage(path)
*/
type StoragePathEvaluator struct {
	app core.Application
}

func (r *StoragePathEvaluator) GetFunctionName() string {
	return "storage"
}
func (r *StoragePathEvaluator) Eval(params []string, def interface{}) interface{} {
	if len(params) > 0 {
		return r.app.StoragePath(params[0])
	} else {
		return r.app.StoragePath("")
	}
}

var _ gConfig.EvaluatorFunction = (*StoragePathEvaluator)(nil)

func NewStoragePathEvaluator(app core.Application) *StoragePathEvaluator {
	return &StoragePathEvaluator{
		app: app,
	}
}

/**
Resource Path: resource(path)
*/
type ResourcesPathEvaluator struct {
	app core.Application
}

func (r *ResourcesPathEvaluator) GetFunctionName() string {
	return "resources"
}
func (r *ResourcesPathEvaluator) Eval(params []string, def interface{}) interface{} {
	if len(params) > 0 {
		return r.app.ResourcesPath(params[0])
	} else {
		return r.app.ResourcesPath("")
	}
}

var _ gConfig.EvaluatorFunction = (*ResourcesPathEvaluator)(nil)

func NewResourcesPathEvaluator(app core.Application) *ResourcesPathEvaluator {
	return &ResourcesPathEvaluator{
		app: app,
	}
}

/**
Root Path: root(path)
*/
type RootPathEvaluator struct {
	app core.Application
}

func (r *RootPathEvaluator) GetFunctionName() string {
	return "root"
}
func (r *RootPathEvaluator) Eval(params []string, def interface{}) interface{} {
	if len(params) > 0 {
		return r.app.RootPath(params[0])
	} else {
		return r.app.RootPath("")
	}
}

var _ gConfig.EvaluatorFunction = (*RootPathEvaluator)(nil)

func NewRootPathEvaluator(app core.Application) *RootPathEvaluator {
	return &RootPathEvaluator{
		app: app,
	}
}

/**
System parameters: system(parameter, default)
*/
type SystemParameterEvaluator struct {
}

func (s *SystemParameterEvaluator) GetFunctionName() string {
	return "system"
}
func (s *SystemParameterEvaluator) Eval(params []string, def interface{}) interface{} {
	if len(params) > 0 {
		switch params[0] {
		case "os":
		case "platform":
			return app.Platform
		case "variant":
			return app.Variant
		}
	}
	return def
}

var _ gConfig.EvaluatorFunction = (*SystemParameterEvaluator)(nil)

/**
Time : time(parameter)
parameter:
	hour, minute, second, now
*/
type TimeEvaluator struct {
}

func (t *TimeEvaluator) GetFunctionName() string {
	return "time"
}
func (t *TimeEvaluator) Eval(params []string, def interface{}) interface{} {
	if len(params) > 0 {
		switch params[0] {
		case "hour":
			return time.Hour
		case "minute":
			return time.Minute
		case "second":
			return time.Second
		case "now":
			return time.Now()
		}
	}
	return def
}

var _ gConfig.EvaluatorFunction = (*TimeEvaluator)(nil)