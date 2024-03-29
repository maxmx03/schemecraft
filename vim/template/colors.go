package template

func Colors() string {
	return `vim9script

# {{.Name}} Colorscheme for vim
# Url: {{.Repo}}
# Maintainer: {{.Author}} <{{.Contact}}>
# License: {{.License}}

if exists("g:loaded_{{mustRegexFind "^[a-z]+" .Name}}")
  finish
endif

g:loaded_{{mustRegexFind "^[a-z]+" .Name}} = 1

hi clear

if exists('syntax_on')
  syntax reset
endif

set termguicolors
g:colors_name = '{{.Name}}'
g:{{mustRegexFind "^[a-z]+" .Name}} = {
  plugins: {},
}
{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
g:{{mustRegexFind "^[a-z]+" $.Name}}.plugins["{{$pluginName}}"] = false
{{end -}}
{{end -}}

{{ $counter := 0 }}
if (has('termguicolors') && &termguicolors) || has('gui_running')
  g:terminal_ansi_colors = [
    {{ range $index, $color := .Palette -}}
    {{if lt $counter 16 -}}
    "{{ upper $color }}",
    {{ $counter = add $counter 1 }}
    {{- end -}}
    {{- end -}}
  ]
endif

{{range $index, $group := .Highlights.Editor}}
  {{- if get $group "link"}}
hi! link {{$group.name}} {{$group.link}}
  {{- else}}
hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
  {{- end -}}
{{- end -}}

{{range $index, $group := .Highlights.Syntax}}
  {{- if get $group "link"}}
hi! link {{$group.name}} {{$group.link}}
  {{- else}}
{{- if not (hasPrefix "@" $group.name)}}
hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- range $x, $groups := .Highlights.Plugins}}
  {{- range $y, $group_vals := $groups}}
if g:{{$.Name}}.plugins["{{$y}}"]
  {{- range $z, $group := $group_vals}}
  {{- if not (hasPrefix "@" $group.name)}}
  {{- if get $group "link"}}
  hi! link {{$group.name}} {{$group.link}}
  {{- else}}
  hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
  {{- end -}}
  {{- end -}}
{{end}}
endif
{{end -}}
{{end -}}`
}
