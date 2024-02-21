package scheme

type Config interface {
	GetTransparent() bool
	GetOnHighlights() string
	GetOnColors() string
}

type Highlights interface {
	GetEditor() []interface{}
	GetSyntax() []interface{}
	GetPlugins() []interface{}
}

type Scheme interface {
	GetName() string
	GetVersion() string
	GetAuthor() string
	GetRepo() string
	GetLicense() string
	GetContact() string
	GetConfig() Config
	GetPalette() map[string]interface{}
	GetHighlights() Highlights
}
