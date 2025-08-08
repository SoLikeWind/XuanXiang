package errors

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ERROR_CONVERT_ARTICLE = errors.New(1021, "ERROR_CONVERT_Article", "转换Article错误")
	ERROR_CREATE_ARTICLE  = errors.New(1022, "ERROT_CREATE_ARTICLE", "创建Article错误")
	ERROR_LIST_ARTICLE    = errors.New(1023, "ERROT_LIST_ARTICLE", "获取Articles错误")
	ERROR_GET_ARTICLE     = errors.New(1023, "ERROT_GET_ARTICLE", "获取单个Article错误")
	ERROR_UPDATE_ARTICLE  = errors.New(1024, "ERROT_UPDATE_ARTICLE", "更新Article错误")
	ERROR_DELETE_ARTICLE  = errors.New(1025, "ERROT_DELETE_ARTICLE", "删除Article错误")

	ERROR_CREATE_TAG = errors.New(1031, "ERROT_CREATE_TAG", "创建Tag错误")
	ERROR_LIST_TAG   = errors.New(1032, "ERROT_LIST_TAG", "获取Tag列表错误")
	ERROR_COUNT_TAG  = errors.New(1033, "ERROT_COUNT_TAG", "统计Tag数量错误")
	ERROR_GET_TAG    = errors.New(1034, "ERROT_GET_TAG", "获取Tag错误")
)

// Error 包装错误，将原生错误信息添加到自定义错误中
func Error(customErr *errors.Error, originalErr error) error {
	if originalErr == nil {
		return customErr
	}

	// 创建一个新的错误，包含原始错误信息
	return errors.New(
		int(customErr.Code),
		customErr.Reason,
		fmt.Sprintf("%s: %v", customErr.Message, originalErr),
	)
}

// // WrapWithContext 包装错误并添加上下文信息
// func WrapWithContext(customErr *errors.Error, originalErr error, context string) error {
// 	if originalErr == nil {
// 		return customErr
// 	}

// 	return errors.New(
// 		customErr.Code,
// 		customErr.Reason,
// 		fmt.Sprintf("%s [%s]: %v", customErr.Message(), context, originalErr),
// 	)
// }
