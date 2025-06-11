package convert

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func MdToHtml(md string) string {
	// TODO: Convert Markdown to HTML
	var buf bytes.Buffer                                       //创建缓冲区用于接收 HTML 输出
	if err := goldmark.Convert([]byte(md), &buf); err != nil { //转换md为html并写入buf中
		return "" //返回空字符串，代表转换失败
	}
	return buf.String() //将buf的html作为字符串返回
}
