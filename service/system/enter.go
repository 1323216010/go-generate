package system

type ServiceGroup struct {
	ApiService
	InitDBService
	AutoCodeService
	DictionaryService
	AutoCodeHistoryService
	DictionaryDetailService
}
