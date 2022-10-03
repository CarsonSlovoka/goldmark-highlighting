package highlighting

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2"
	chromaHtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
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

// nodeRendererFunc 此函數可以把結果直接寫到writer之中，就能對輸出產生影響
func (r *HTMLRenderer) nodeRendererFunc(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.FencedCodeBlock) // 轉成對應的ast.Node

	// Attributes, // parse ```go{hl_lines=["2-3",5], linenostart=5}
	{
		// TODO
	}

	var codeBlockContent bytes.Buffer
	l := n.Lines().Len()
	for i := 0; i < l; i++ { // 把code-block的內容，一列一列的寫進去
		line := n.Lines().At(i)
		codeBlockContent.Write(line.Value(source)) // source是一個比較大的項目，可以包含code-block以外的內容
	}

	language := n.Language(source) // 當前code-block所用的語言 ```myLang

	// 定義預設的code-block函數
	defaultCodeBlockHandlerFunc := func() (ast.WalkStatus, error) {
		// Header
		{
			if language != nil {
				_, _ = w.WriteString(fmt.Sprintf("<pre>\n<code class=\"language-%s\">", language))
			} else {
				_, _ = w.WriteString("<pre>\n<code>")
			}
		}
		// Body
		{
			_, _ = w.WriteString(codeBlockContent.String())
		}
		// Tail
		{
			_, _ = w.WriteString("</code>\n</pre>\n")
		}
		return ast.WalkContinue, nil
	}

	var lexer chroma.Lexer
	if language != nil {
		lexer = lexers.Get(string(language))
		if lexer != nil {
			return r.writeCodeBlock(
				w, lexer, &codeBlockContent,
				defaultCodeBlockHandlerFunc, // 如果錯誤就用預設的code-block取代
			)
		}
	}

	if !r.NoHighlight && r.GuessLanguage {
		if lexer = lexers.Analyse(codeBlockContent.String()); lexer != nil {
			return r.writeCodeBlock(w, lexer, &codeBlockContent, defaultCodeBlockHandlerFunc)
		}
	}

	return defaultCodeBlockHandlerFunc()
}

func (r *HTMLRenderer) writeCodeBlock(
	w util.BufWriter, lexer chroma.Lexer, codeBlockContent *bytes.Buffer,
	defaultHandlerFunc func() (ast.WalkStatus, error),
) (ast.WalkStatus, error) {

	if lexer == nil {
		return ast.WalkContinue, nil
	}
	language := []byte(lexer.Config().Name)
	lexer = chroma.Coalesce(lexer) // 常態化lexer。所得到的結果還是一個Laxer型別

	iterator, err := lexer.Tokenise(nil, codeBlockContent.String()) // 這裡開始準備解析code-block的文本，他會把所有文本解析出個別的關鍵字，整理出一份token清單，可以被疊代
	if err != nil {
		return defaultHandlerFunc()
	}

	formatter := chromaHtml.New()
	style := r.CustomStyle
	if style == nil { // 無自定義主題，則嘗試用chroma所註冊的主題列表中搜尋匹配的主題
		style = styles.Get(r.Style) // 注意，如果名稱匹配找不到返回的是chroma.Fallback的主題
	}

	// Head
	_, _ = w.WriteString(fmt.Sprintf("<div class=\"highlight %s\">\n", language))

	// Body
	_ = formatter.Format(w, style, iterator) // chroma的核心，可以把該token渲染成指定的樣式

	// Tail
	_, _ = w.WriteString("\n</div>")

	return ast.WalkContinue, nil
}
