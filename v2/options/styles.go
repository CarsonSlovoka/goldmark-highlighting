package options

import (
	"fmt"
	. "github.com/CarsonSlovoka/goldmark-highlighting/v2"
	"github.com/alecthomas/chroma/v2"
	"github.com/yuin/goldmark/renderer"
	"os"
)

type withStyle struct {
	value string
}

func (o *withStyle) SetConfig(c *renderer.Config) {
	c.Options[optStyle] = o.value
}

func (o *withStyle) SetHighlightingOption(config *RendererConfig) {
	config.Style = o.value
}

// WithStyle chroma.styles.Get(style) == nil =>  styles = chroma.Fallback
func WithStyle(style string) Option {
	return &withStyle{style}
}

type withCustomStyle struct {
	value *chroma.Style
}

func (o *withCustomStyle) SetConfig(c *renderer.Config) {
	c.Options[optCustomStyle] = o.value
}

func (o *withCustomStyle) SetHighlightingOption(config *RendererConfig) {
	config.CustomStyle = o.value
	if config.Style != "" {
		_, _ = fmt.Fprintf(os.Stderr, "Ignore the 'Style:%s' since the 'CustomStyle:%s' have been set.\n", config.Style, config.CustomStyle.Name)
	}
}

// WithCustomStyle define the style by yourself
// https://xyproto.github.io/splash/docs/
// https://github.com/alecthomas/chroma/tree/3f86ac7/styles
// 只是定義關鍵字的顏色，如果您需要對文本內容進行解析要去擴展lexer
// https://github.com/alecthomas/chroma/blob/master/lexers/html.go
// https://github.com/alecthomas/chroma/pull/276/files
func WithCustomStyle(style *chroma.Style) Option {
	return &withCustomStyle{style}
}
