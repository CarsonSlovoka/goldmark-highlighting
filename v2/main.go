package highlighting

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type highlightingExtender struct {
}

func NewHighlightingExtender() goldmark.Extender {
	return &highlightingExtender{}
}

// Extend adds a hashtag parser to a Goldmark parser
func (e *highlightingExtender) Extend(m goldmark.Markdown) {
	// m.Parser().AddOptions() // 如果您的擴展不需要再額外自己定義ast，那就不需要再定義Parser。 Parser只是在Parse的時候，會把相關屬性塞給該ast罷了，而codeblock運用的ast用已定義的ast.KindFencedCodeBlock即可，所以不需額外定義

	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewHighlightingHTMLRenderer(), 200),
		),
	)
}
