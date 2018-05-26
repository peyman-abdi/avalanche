package config

import (
	gConfig "github.com/peyman-abdi/conf"
	"github.com/peyman-abdi/avalanche/app/core/app"
	"time"
	"github.com/peyman-abdi/avalanche/app/interfaces"
)

/**
	Storage Path: storage(path)
 */
type StoragePathEvaluator struct {
	app interfaces.Application
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
func NewStoragePathEvaluator(app interfaces.Application) *StoragePathEvaluator {
	return &StoragePathEvaluator{
		app: app,
	}
}

/**
	Resource Path: resource(path)
 */
type ResourcesPathEvaluator struct {
	app interfaces.Application
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
func NewResourcesPathEvaluator(app interfaces.Application) *ResourcesPathEvaluator {
	return &ResourcesPathEvaluator{
		app: app,
	}
}

/**
	Root Path: root(path)
 */
type RootPathEvaluator struct {
	app interfaces.Application
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
func NewRootPathEvaluator(app interfaces.Application) *RootPathEvaluator {
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