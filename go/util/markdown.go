package util

import (
	"Campus-forum-system/logs"
	"strings"
	"sync"

	"github.com/88250/lute"
	"github.com/PuerkitoBio/goquery"
)

var (
	engine *lute.Lute
	once   sync.Once
)

func getEngine() *lute.Lute {
	// 单例模式 Do 调用函数 f 当且仅当 Do 被调用 第一次使用此 Once 实例。
	once.Do(func() {
		// 创建一个 Lute 引擎
		engine = lute.New(func(lute *lute.Lute) {
			//用于启用或禁用自动生成文档目录（Table of Contents，简称 TOC）的功能
			lute.SetToC(true) // 当设置为 true 时，Lute 将会在 Markdown 文档中自动生成目录。
			// 用于启用或禁用 GitHub Flavored Markdown（GFM）中的删除线（strikethrough）语法
			lute.SetGFMStrikethrough(true) // 当设置为 true 时，Lute 将会解析并渲染 Markdown 中使用的删除线语法
		})
	})
	return engine
}

// MarkdownToHTML 将 Markdown 文本转换为 HTML 文本
func MarkdownToHTML(markdownStr string) string {
	if IsBlank(markdownStr) {
		return ""
	}
	// func (lute *Lute) MarkdownStr(name, markdown string) (html string)  MarkdownStr 接受 string 类型的 markdown 后直接调用 Markdown 进行处理。
	return getEngine().MarkdownStr("", markdownStr)
}

func GetHTMLText(html string) string {
	txt, err := goquery.NewDocumentFromReader(strings.NewReader(html)) // NewReader 返回一个新的 Reader 读取 s。 它类似于字节。NewBufferString，但效率更高且不可写。
	if err != nil {
		logs.Logger.Errorf("从html读取文本出错", txt)
	}
	return txt.Text() // Text 获取匹配集合中每个元素的组合文本内容 元素，包括它们的后代。
}
