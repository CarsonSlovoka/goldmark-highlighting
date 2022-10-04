// 這是一個renderer最基本應該實現那些方法
// 如果有其他的方法要新增請加在renderer-method

package highlighting

import (
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

type HTMLRenderer struct {
	*RendererConfig
}

func NewHTMLRenderer(opts ...Option) renderer.NodeRenderer {
	r := &HTMLRenderer{&RendererConfig{
		// GuessLanguage: true, // 預設不啟用
	}}
	for _, o := range opts {
		o.SetHighlightingOption(r.RendererConfig) // 把所有opt的結果寫入到r.config之中
	}

	if r.Style == "" && r.CustomStyle == nil {
		r.Style = styles.GitHub.Name
	}

	return r
}

func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.nodeRendererFunc) // codeBlock不需要再定義額外的ast去描述，基於原始的KindFencedCodeBlock增加額外的判斷即可
}
