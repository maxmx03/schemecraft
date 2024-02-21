package template

func Palette() string {
	return `---@class {{mustRegexFind "^[a-z]+" .Name}}.palette
{{range $index, $value := .Palette -}}
---@field {{$index}}? string
{{end -}}
return {
  {{range $index, $value := .Palette -}}
  {{"  "}}{{$index}} = "{{upper $value}}",
  {{end -}}
}`
}
