package nvim

func configTemplate() string {
	return `return {
  transparent = {{.Config.Transparent}},
  on_highlights = {{.Config.OnHighlights}},
  on_colors = {{.Config.OnColors}},
  plugins = {
  {{range $index, $value := .Config.Plugins -}}
  {{"  "}}["{{$index}}"] = {{$value}},
  {{end -}}
  }
}
`
}
