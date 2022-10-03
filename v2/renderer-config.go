package highlighting

import "encoding/json"

type RendererConfig struct {
	NoHighlight   bool
	GuessLanguage bool
}

func (c *RendererConfig) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}
