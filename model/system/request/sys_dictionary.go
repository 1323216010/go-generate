package request

import (
	"go-generate/model/common/request"
	"go-generate/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
