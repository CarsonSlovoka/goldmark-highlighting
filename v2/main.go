package highlighting

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Option interface {
	renderer.Option // 要實作SetConfig(c *renderer.Config) // 其實不一定要設定，但是它可以讓renderer.Config更能知道總共有哪些內容
	SetHighlightingOption(*RendererConfig)
}

type highlightingExtender struct {
	options []Option // 慣用方法，通常Extender如果有需要其他可選項，都是用類似的方法去添加
}

func NewExtender(opts ...Option) goldmark.Extender {
	return &highlightingExtender{opts}
}

// Extend adds a hashtag parser to a Goldmark parser
func (e *highlightingExtender) Extend(m goldmark.Markdown) {
	// m.Parser().AddOptions() // 如果您的擴展不需要再額外自己定義ast，那就不需要再定義Parser。 Parser只是在Parse的時候，會把相關屬性塞給該ast罷了，而codeblock運用的ast用已定義的ast.KindFencedCodeBlock即可，所以不需額外定義

	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewHTMLRenderer(e.options...), 200),
		),
	)
}
