package scheme

type Config struct {
	Transparent  bool            `yaml:"transparent"`
	OnHighlights string          `yaml:"on_highlights"`
	OnColors     string          `yaml:"on_colors"`
	Plugins      map[string]bool `yaml:"plugins"`
}

type Scheme struct {
	Name    string                 `yaml:"name"`
	Author  string                 `yaml:"author"`
	Repo    string                 `yaml:"repo"`
	License string                 `yaml:"license"`
	Config  Config                 `yaml:"config"`
	Palette map[string]interface{} `yaml:"palette"`
}
