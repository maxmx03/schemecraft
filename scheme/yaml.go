package scheme

type ConfigYaml struct {
	Transparent  bool   `yaml:"transparent"`
	OnHighlights string `yaml:"on_highlights"`
	OnColors     string `yaml:"on_colors"`
}

type HighlightsYaml struct {
	Editor  []interface{} `yaml:"editor"`
	Syntax  []interface{} `yaml:"syntax"`
	Plugins []interface{} `yaml:"plugins"`
}

type SchemeYaml struct {
	Name       string                 `yaml:"name"`
	Author     string                 `yaml:"author"`
	Repo       string                 `yaml:"repo"`
	License    string                 `yaml:"license"`
	Contact    string                 `yaml:"contact"`
	Config     ConfigYaml             `yaml:"config"`
	Palette    map[string]interface{} `yaml:"palette"`
	Highlights HighlightsYaml         `yaml:"highlights"`
}

func (s SchemeYaml) GetName() string {
	return s.Name
}

func (s SchemeYaml) GetAuthor() string {
	return s.Author
}

func (s SchemeYaml) GetRepo() string {
	return s.Repo
}

func (s SchemeYaml) GetLicense() string {
	return s.License
}

func (s SchemeYaml) GetContact() string {
	return s.Contact
}

func (s SchemeYaml) GetConfig() Config {
	return s.Config
}

func (s SchemeYaml) GetPalette() map[string]interface{} {
	return s.Palette
}

func (s SchemeYaml) GetHighlights() Highlights {
	return s.Highlights
}

func (c ConfigYaml) GetTransparent() bool {
	return c.Transparent
}

func (c ConfigYaml) GetOnHighlights() string {
	return c.OnHighlights
}

func (c ConfigYaml) GetOnColors() string {
	return c.OnColors
}

func (h HighlightsYaml) GetEditor() []interface{} {
	return h.Editor
}

func (h HighlightsYaml) GetSyntax() []interface{} {
	return h.Syntax
}

func (h HighlightsYaml) GetPlugins() []interface{} {
	return h.Plugins
}
