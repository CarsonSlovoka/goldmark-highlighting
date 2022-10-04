package highlighting

import (
	"encoding/json"
	"github.com/alecthomas/chroma/v2"
	chromaHtml "github.com/alecthomas/chroma/v2/formatters/html"
)

type RendererConfig struct {
	// Style is a highlighting style.
	// Supported styles are defined under https://github.com/alecthomas/chroma/tree/3f86ac7/styles
	// 使用chroma所提供的主題
	Style string

	// Pass in a custom Chroma style. If this is not nil, the Style string will be ignored
	// 自定義Chroma.Style
	// https://xyproto.github.io/splash/docs/
	// https://github.com/alecthomas/chroma/tree/3f86ac7/styles
	CustomStyle *chroma.Style

	// See https://github.com/alecthomas/chroma#the-html-formatter for details.
	// permalink: https://github.com/alecthomas/chroma/blob/d38b87110b078027006bc34aa27a065fa22295a1/README.md?plain=1#L175-L190
	FormatOptions []chromaHtml.Option

	NoHighlight   bool
	GuessLanguage bool
}

func (c *RendererConfig) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}
