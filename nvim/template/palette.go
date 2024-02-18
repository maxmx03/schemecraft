package template

func Palette() string {
	return `return {
  {{range $index, $value := .Palette -}}
  {{"  "}}{{$index}} = "{{upper $value}}",
  {{end -}}
}`
}
