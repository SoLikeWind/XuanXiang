package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	ERROR_CONVERT_ARTICLE = errors.New(1021, "ERROR_CONVERT_Article", "转换Article错误")
	ERROT_CREATE_ARTICLE  = errors.New(1022, "ERROT_CREATE_ARTICLE", "创建Article错误")
)
