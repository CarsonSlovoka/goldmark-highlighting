package highlighting

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2"
	chromaHtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Attributes https://github.com/alecthomas/chroma/blob/d38b87110b078027006bc34aa27a065fa22295a1/README.md?plain=1#L175-L190
type Attributes struct {
	Style   string // 大小寫有差異，如果用小寫reflect無法CanSet
	Hls     string // highlight numbers
	Base    string // baseline number
	Linenos string // true, false, table, inline
	TW      string // tab width (default = 4) chroma預設值為8我改成4
	LPrefix string // LinkableLineNumbers
}

func (attr *Attributes) Set(name, val string) {
	pointToStruct := reflect.ValueOf(attr)
	curStruct := pointToStruct.Elem()

	curField := curStruct.FieldByName(name) // type: Value
	// 注意如果name找不到會引發:  call of reflect.Value.SetString on zero Value
	curField.SetString(val)
}

// getAttributes
// parse ```go {style="vim" hls=[1, 3-5] base=3 linenos=table tw=2}
func (r *HTMLRenderer) getAttributes(node *ast.FencedCodeBlock, source []byte) []chromaHtml.Option {

	if node.Info == nil { // Info為一個*ast.Text
		return nil
	}

	info := node.Info.Segment.Value(source) // 取得 ```go{style="git" hls=[2-3, 8, 98-99] base=5 linenos=inline} 的內容
	if info == nil {
		return nil // 無任何屬性資料
	}
	// /S+排除大括號: [^\r\n\t\f\v {}]*
	re := regexp.MustCompile(`(?i)LPrefix=(?P<LPrefix>\S*)|tw=(?P<TW>\d*)|style="(?P<Style>\S*)"|hls=\[(?P<Hls>[\d\-, ]*)]*|base=(?P<Base>\d*)|linenos=(?P<Linenos>inline|table|true|false|1|0)`)
	ms := re.FindAllStringSubmatch(strings.TrimRight(string(info), "}"), -1) // 修剪右側}正規式就可不必再多做判斷
	var attr *Attributes
	for _, m := range ms {
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" && m[i] != "" {
				if attr == nil {
					attr = new(Attributes)
				}
				attr.Set(name, m[i])
				break // // 理論上我們的匹配模式，一個match只會有一個項目匹配，所以如果匹配到了，其他的項目就不用再考慮
			}
		}
	}

	if attr == nil {
		return nil
	}

	chromaFormatterOptions := make([]chromaHtml.Option, len(r.FormatOptions)) // 複製一份，避免更動原始資料
	copy(chromaFormatterOptions, r.FormatOptions)

	if linkPrefix := attr.LPrefix; linkPrefix != "" {
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.LinkableLineNumbers(true, linkPrefix))
		if attr.Linenos == "" {
			attr.Linenos = "true" // 在有指定linkPrefix時，如果沒指定打開列號，會自動幫忙開啟
		}
	}

	if attr.TW != "" {
		if tabWidth, err := strconv.Atoi(attr.TW); err == nil {
			chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.TabWidth(tabWidth))
		}
	} else {
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.TabWidth(4))
	}

	baselineNumber := 1
	if attr.Base != "" {
		if n, err := strconv.Atoi(attr.Base); err == nil {
			baselineNumber = n
			chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.BaseLineNumber(baselineNumber))
		}
	}

	switch attr.Linenos {
	case "1":
		fallthrough
	case "true":
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.WithLineNumbers(true))
	case "0":
		fallthrough
	case "false":
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.WithLineNumbers(false))
	case "table":
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.WithLineNumbers(true))
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.LineNumbersInTable(true))
	case "inline":
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.WithLineNumbers(true))
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.LineNumbersInTable(false))
	}

	if attr.Hls != "" {
		var hlRanges [][2]int // 每一個元素都使一個[begin end]，如果只有單列，用[begin begin]
		var (
			begin string
			end   string
		)
		var cur *string
		cur = &begin
		for _, ch := range attr.Hls { // 2, 8-19 ...
			if !unicode.IsDigit(ch) && ch != ',' && ch != ' ' && ch != '-' {
				fmt.Fprintf(os.Stderr, "Hls contains an illegal character. %s\n", attr.Hls)
				break
			}

			if ch == ',' || ch == ' ' {
				if begin == "" {
					continue
				}
				beginI, _ := strconv.Atoi(begin)
				if end == "" {
					hlRanges = append(hlRanges, [2]int{baselineNumber + beginI - 1, baselineNumber + beginI - 1})
				} else {
					endI, _ := strconv.Atoi(end)
					hlRanges = append(hlRanges, [2]int{baselineNumber + beginI - 1, baselineNumber + endI - 1})
				}
				begin = ""
				end = ""
				cur = &begin
				continue
			}

			if ch == '-' {
				cur = &end
				continue
			}

			*cur += string(ch)
			continue
			// if ch != " "
		}

		if begin != "" { // 處理結尾資料
			beginI, _ := strconv.Atoi(begin)
			if end == "" {
				hlRanges = append(hlRanges, [2]int{baselineNumber + beginI - 1, baselineNumber + beginI - 1})
			} else {
				endI, _ := strconv.Atoi(end)
				hlRanges = append(hlRanges, [2]int{baselineNumber + beginI - 1, baselineNumber + endI - 1})
			}
		}
		chromaFormatterOptions = append(chromaFormatterOptions, chromaHtml.HighlightLines(hlRanges))
	}

	if style := attr.Style; style != "" {
		r.Style = style
	}
	return chromaFormatterOptions
}

func (r *HTMLRenderer) writeCodeBlock(
	w util.BufWriter, lexer chroma.Lexer, codeBlockContent *bytes.Buffer,
	chromaOptions []chromaHtml.Option,
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

	formatter := chromaHtml.New(chromaOptions...)
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

// nodeRendererFunc 此函數可以把結果直接寫到writer之中，就能對輸出產生影響
func (r *HTMLRenderer) nodeRendererFunc(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.FencedCodeBlock) // 轉成對應的ast.Node

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
	formatter := r.getAttributes(n, source)
	if language != nil {
		lexer = lexers.Get(string(language))
		if lexer != nil {
			return r.writeCodeBlock(
				w, lexer, &codeBlockContent,
				formatter,
				defaultCodeBlockHandlerFunc, // 如果錯誤就用預設的code-block取代
			)
		}
	}

	if !r.NoHighlight && r.GuessLanguage {
		if lexer = lexers.Analyse(codeBlockContent.String()); lexer != nil {
			return r.writeCodeBlock(w, lexer, &codeBlockContent, formatter, defaultCodeBlockHandlerFunc)
		}
	}

	return defaultCodeBlockHandlerFunc()
}
