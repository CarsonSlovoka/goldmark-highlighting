package options

import (
	. "github.com/CarsonSlovoka/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/renderer"
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

func WithGuessLanguage(enable bool) Option {
	return &withGuessLanguage{enable}
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
