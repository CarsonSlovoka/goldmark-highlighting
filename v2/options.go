package highlighting

import (
	"fmt"
	"github.com/alecthomas/chroma/v2"
	"github.com/yuin/goldmark/renderer"
	"os"
)

const (
	optGuessLanguage renderer.OptionName = "HighlightingGuessLanguage"
	optNoHighlight   renderer.OptionName = "HighlightingNoHighlight"
	optStyle         renderer.OptionName = "HighlightingStyle"
	optCustomStyle   renderer.OptionName = "HighlightingCustomStyle"
)

type withGuessLanguage struct {
	value bool
}

func (o *withGuessLanguage) SetConfig(c *renderer.Config) {
	c.Options[optGuessLanguage] = o.value
}

func (o *withGuessLanguage) SetHighlightingOption(config *RendererConfig) {
	config.GuessLanguage = o.value
}

func WithGuessLanguage(val bool) Option {
	return &withGuessLanguage{val}
}

type withNoHighlight struct {
	value bool
}

func (o *withNoHighlight) SetConfig(c *renderer.Config) {
	c.Options[optNoHighlight] = o.value
}

func (o *withNoHighlight) SetHighlightingOption(config *RendererConfig) {
	config.NoHighlight = o.value
}

func WithNoHighlight(val bool) Option {
	return &withNoHighlight{val}
}

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
func WithCustomStyle(style *chroma.Style) Option {
	return &withCustomStyle{style}
}
