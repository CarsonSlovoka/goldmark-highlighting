package highlighting

import (
	"encoding/json"
	"github.com/alecthomas/chroma/v2"
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

	NoHighlight   bool
	GuessLanguage bool
}

func (c *RendererConfig) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}
