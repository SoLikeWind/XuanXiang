package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	ERROR_CONVERT_ARTICLE = errors.New(1021, "ERROR_CONVERT_Article", "转换Article错误")
	ERROT_CREATE_ARTICLE  = errors.New(1022, "ERROT_CREATE_ARTICLE", "创建Article错误")
	ERROT_LIST_ARTICLE    = errors.New(1023, "ERROT_LIST_ARTICLE", "获取Article列表错误")
	ERROT_GET_ARTICLE     = errors.New(1023, "ERROT_GET_ARTICLE", "获取Article错误")
	ERROT_UPDATE_ARTICLE  = errors.New(1024, "ERROT_UPDATE_ARTICLE", "更新Article错误")
	ERROT_DELETE_ARTICLE  = errors.New(1025, "ERROT_DELETE_ARTICLE", "删除Article错误")

	ERROT_CREATE_TAG = errors.New(1031, "ERROT_CREATE_TAG", "创建Tag错误")
	ERROT_LIST_TAG   = errors.New(1032, "ERROT_LIST_TAG", "获取Tag列表错误")
	ERROT_COUNT_TAG  = errors.New(1033, "ERROT_COUNT_TAG", "统计Tag数量错误")
	ERROT_GET_TAG    = errors.New(1034, "ERROT_GET_TAG", "获取Tag错误")
)
