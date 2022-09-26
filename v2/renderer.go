package highlighting

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type highlightingHTMLRenderer struct{}

func NewHighlightingHTMLRenderer() renderer.NodeRenderer {
	return &highlightingHTMLRenderer{}
}

func (r *highlightingHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.nodeRendererFunc) // codeblock不需要再定義額外的ast去描述，基於原始的KindFencedCodeBlock增加額外的判斷即可
}

// nodeRendererFunc 此函數可以把結果直接寫到writer之中，就能對輸出產生影響
func (r *highlightingHTMLRenderer) nodeRendererFunc(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return 0, nil
}
