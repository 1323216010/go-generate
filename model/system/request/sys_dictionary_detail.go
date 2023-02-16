package request

import (
	"go-generate/model/common/request"
	"go-generate/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
