package scheme

type ConfigJson struct {
	Transparent  bool   `json:"transparent"`
	OnHighlights string `json:"on_highlights"`
	OnColors     string `json:"on_colors"`
}

type HighlightsJson struct {
	Editor  []interface{} `json:"editor"`
	Syntax  []interface{} `json:"syntax"`
	Plugins []interface{} `json:"plugins"`
}

type SchemeJson struct {
	Name       string                 `json:"name"`
	Author     string                 `json:"author"`
	Repo       string                 `json:"repo"`
	License    string                 `json:"license"`
	Contact    string                 `json:"contact"`
	Config     ConfigJson             `json:"config"`
	Palette    map[string]interface{} `json:"palette"`
	Highlights HighlightsJson         `json:"highlights"`
}

func (s SchemeJson) GetName() string {
	return s.Name
}

func (s SchemeJson) GetAuthor() string {
	return s.Author
}

func (s SchemeJson) GetRepo() string {
	return s.Repo
}

func (s SchemeJson) GetLicense() string {
	return s.License
}

func (s SchemeJson) GetContact() string {
	return s.Contact
}

func (s SchemeJson) GetConfig() Config {
	return s.Config
}

func (s SchemeJson) GetPalette() map[string]interface{} {
	return s.Palette
}

func (s SchemeJson) GetHighlights() Highlights {
	return s.Highlights
}

func (c ConfigJson) GetTransparent() bool {
	return c.Transparent
}

func (c ConfigJson) GetOnHighlights() string {
	return c.OnHighlights
}

func (c ConfigJson) GetOnColors() string {
	return c.OnColors
}

func (h HighlightsJson) GetEditor() []interface{} {
	return h.Editor
}

func (h HighlightsJson) GetSyntax() []interface{} {
	return h.Syntax
}

func (h HighlightsJson) GetPlugins() []interface{} {
	return h.Plugins
}
