package scheme

type Config struct {
	Transparent bool `yaml:"transparent"`
}

type Highlights struct {
	Editor  []any `yaml:"editor"`
	Syntax  []any `yaml:"syntax"`
	Plugins []any `yaml:"plugins"`
}

type Scheme struct {
	Name       string         `yaml:"name"`
	Version    string         `yaml:"version"`
	Author     string         `yaml:"author"`
	Repo       string         `yaml:"repo"`
	License    string         `yaml:"license"`
	Contact    string         `yaml:"contact"`
	Config     Config         `yaml:"config"`
	Palette    map[string]any `yaml:"palette"`
	Highlights Highlights     `yaml:"highlights"`
}
