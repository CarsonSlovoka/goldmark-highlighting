package highlighting_test

import (
	"bytes"
	"errors"
	"fmt"
	highlighting "github.com/CarsonSlovoka/goldmark-highlighting/v2"
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
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewExtender(),
		),
	)

	if err := startTest(markdown, "basic"); err != nil {
		t.Fatal(err)
	}
}

func ExampleHTMLRenderer_options() {
	renderDefault := highlighting.NewHTMLRenderer().(*highlighting.HTMLRenderer)
	fmt.Println("Default")
	fmt.Printf("%+v\n", renderDefault)

	renderCustomize := highlighting.NewHTMLRenderer(
		highlighting.WithGuessLanguage(true),
		highlighting.WithNoHighlight(true),
	).(*highlighting.HTMLRenderer)
	fmt.Println("Customize")
	fmt.Printf("%+v\n", renderCustomize)

	// Output:
	// Default
	// {"NoHighlight":false,"GuessLanguage":false}
	// Customize
	// {"NoHighlight":true,"GuessLanguage":true}
}
