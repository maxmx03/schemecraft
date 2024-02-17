package template

func Colors() string {
	return `vim9script

# {{.Name}} Colorscheme for vim
# Url: {{.Repo}}
# Maintainer: {{.Author}} <{{.Contact}}>
# License: {{.License}}

if exists("g:loaded_{{.Name}}")
  finish
endif

g:loaded_{{.Name}} = 1

hi clear

if exists('syntax_on')
  syntax reset
endif

set termguicolors
g:colors_name = '{{.Name}}'
g:{{.Name}} = {
  plugins: {},
}
{{range $index, $value := .Config.Plugins -}}
g:{{$.Name}}.plugins["{{$index}}"] = {{default false $value}}
{{end -}}

if (has('termguicolors') && &termguicolors) || has('gui_running')
  g:terminal_ansi_colors = [
    {{ range $index, $color := .Palette -}}
    "{{upper $color}}",
    {{end -}}
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
hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
{{- end -}}
{{- end -}}

{{- range $x, $groups := .Highlights.Plugins}}
  {{- range $y, $group_vals := $groups}}
if g:{{$.Name}}.plugins["{{$y}}"]
  {{- range $z, $group := $group_vals}}
  {{- if get $group "link" -}}
  hi! link {{$group.name}} {{$group.link}}
  {{else}}
  hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
{{- end -}}
{{end}}
endif
{{end -}}
{{end -}}`
}