package highlighting_test

import (
	"bytes"
	"errors"
	"fmt"
	highlighting "github.com/CarsonSlovoka/goldmark-highlighting/v2"
	hOpts "github.com/CarsonSlovoka/goldmark-highlighting/v2/options"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/yuin/goldmark"
	"log"
	"os"
	"strings"
	"testing"
)

func startTest(markdown goldmark.Markdown, testDir string) error {
	content, err := os.ReadFile(fmt.Sprintf("testData/%s/input.md", testDir))
	if err != nil {
		return err
	}
	got := bytes.NewBuffer(make([]byte, 0))
	if err = markdown.Convert(content, got); err != nil {
		return err
	}
	want, err := os.ReadFile(fmt.Sprintf("testData/%s/expected.html", testDir))
	if err != nil {
		return err
	}
	if got.String() != strings.TrimRight(string(want), "\n") { // 移除結尾多出的空行，避免editorconfig中的insert_final_newline與其衝突
		log.Printf("[error]\ngot:\n%s\nwant:\n%s", got, want)
		return errors.New("got != want")
	}
	return nil
}

func Test_NewHighlightingExtender(t *testing.T) {
	mdCustomStyle := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewExtender(
				hOpts.WithCustomStyle(styles.Monokai), // 這個其實是讓您自己去定義chroma.styles，不過如果您不想指定，也可以用這種方式，好處是不會怕打錯字
			),
		),
	)

	mdStyle := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewExtender(hOpts.WithStyle("monokai")), // 不建議用這種名稱來指定style，怕打錯，如果名稱於chroma.styles.Get(style)找不到，會使用chroma.Fallback當作主題樣式
		),
	)

	for _, markdown := range []goldmark.Markdown{mdCustomStyle, mdStyle} {
		if err := startTest(markdown, "basic"); err != nil {
			t.Fatal(err)
		}
	}
}

func ExampleWithCustomStyle() {
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewExtender(
				hOpts.WithCustomStyle(
					styles.Register(chroma.MustNewStyle("my-style", chroma.StyleEntries{
						chroma.Comment:        "italic #888",
						chroma.CommentSpecial: "#888",
						chroma.Keyword:        "#00f",
						chroma.OperatorWord:   "#00f",
						chroma.Name:           "#000",
						chroma.LiteralNumber:  "#3af",
						chroma.LiteralString:  "#5a2",
						chroma.Error:          "#F00",
						chroma.Background:     " bg:#ffffff",
					}),
					),
				))))
	content := `🐬🐬🐬go
package main
import "fmt"

func main() {
	fmt.Println("Hello World!") // comment
}
🐬🐬🐬`
	result := bytes.NewBuffer(make([]byte, 0))
	if err := md.Convert([]byte(strings.ReplaceAll(content, "🐬", "`")), result); err != nil {
		panic(err)
	}

	fmt.Println(result)
	// Output:
	// <div class="highlight Go">
	// <pre tabindex="0" style="background-color:#fff;"><code><span style="display:flex;"><span><span style="color:#00f">package</span> <span style="color:#000">main</span>
	// </span></span><span style="display:flex;"><span><span style="color:#00f">import</span> <span style="color:#5a2">&#34;fmt&#34;</span>
	// </span></span><span style="display:flex;"><span>
	// </span></span><span style="display:flex;"><span><span style="color:#00f">func</span> <span style="color:#000">main</span>() {
	// </span></span><span style="display:flex;"><span>	<span style="color:#000">fmt</span>.<span style="color:#000">Println</span>(<span style="color:#5a2">&#34;Hello World!&#34;</span>) <span style="color:#888;font-style:italic">// comment
	// </span></span></span><span style="display:flex;"><span><span style="color:#888;font-style:italic"></span>}
	// </span></span></code></pre>
	// </div>
}

func ExampleHTMLRenderer_options() {
	renderDefault := highlighting.NewHTMLRenderer().(*highlighting.HTMLRenderer)
	fmt.Println("Default")
	fmt.Println(renderDefault.NoHighlight)
	fmt.Println(renderDefault.GuessLanguage)
	fmt.Println(renderDefault.Style)
	fmt.Println(renderDefault.CustomStyle)

	renderCustomize := highlighting.NewHTMLRenderer(
		hOpts.WithGuessLanguage(true),
		hOpts.WithNoHighlight(true),
		hOpts.WithStyle("github"), // Ignore the 'Style:github' since the 'CustomStyle:vim' have been set.
		hOpts.WithCustomStyle(styles.Vim),
	).(*highlighting.HTMLRenderer)
	fmt.Println("Customize")

	fmt.Println(renderCustomize.NoHighlight)
	fmt.Println(renderCustomize.GuessLanguage)
	fmt.Println(renderCustomize.CustomStyle.Name)

	// Output:
	// Default
	// false
	// false
	// github
	// <nil>
	// Customize
	// true
	// true
	// vim
}
