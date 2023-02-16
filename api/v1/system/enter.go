package system

import "go-generate/service"

type ApiGroup struct {
	SystemApiApi
	DBApi
	AutoCodeApi
	DictionaryApi
	AutoCodeHistoryApi
	DictionaryDetailApi
}

var (
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	autoCodeHistoryService  = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
)
