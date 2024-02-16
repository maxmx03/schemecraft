package scheme

type Config struct {
	Transparent  bool            `yaml:"transparent"`
	OnHighlights string          `yaml:"on_highlights"`
	OnColors     string          `yaml:"on_colors"`
	Plugins      map[string]bool `yaml:"plugins"`
}

type Highlights struct {
	Editor  []interface{} `yaml:"editor"`
	Syntax  []interface{} `yaml:"syntax"`
	Plugins []interface{} `yaml:"plugins"`
}

type Scheme struct {
	Name       string            `yaml:"name"`
	Author     string            `yaml:"author"`
	Repo       string            `yaml:"repo"`
	License    string            `yaml:"license"`
	Config     Config            `yaml:"config"`
	Palette    map[string]string `yaml:"palette"`
	Highlights Highlights        `yaml:"highlights"`
}
