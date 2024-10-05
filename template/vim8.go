package template

import "fmt"

func Vim8Readme() string {
	var readme = `# {{title (mustRegexFind "^[a-z]+" .Name)}}

![schemecraft](https://github.com/maxmx03/schemecraft/assets/50273941/ee682aae-00cb-4282-ba24-3d9621a430a3)

## Installation

To install {{title .Name}}, you need a plugin manager.
In the example, bellow we are going to use vim-plug.

%v

## Transparency

let g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency = 0

## Plugins

Bellow are the {{title .Name}} supported plugins.
Enable the plugins you want.

%v
`
	var installBlockCode = "```vim"
	installBlockCode += `
Plug '{{.Repo}}'

colorscheme {{.Name}}
`
	installBlockCode += "```"

	var plugins = "```vim"
	plugins += `
{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
let g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} = 1
{{end -}}
{{end -}}
`
	plugins += "```"

	return fmt.Sprintf(readme, installBlockCode, plugins)
}

func Vim8Colors() string {
	return `
" {{.Name}} Colorscheme for vim
" Url: {{.Repo}}
" Maintainer: {{.Author}} <{{.Contact}}>
" License: {{.License}}

if exists("g:loaded_{{mustRegexFind "^[a-z]+" .Name}}")
  finish
endif

let g:loaded_{{mustRegexFind "^[a-z]+" .Name}} = 1

hi clear

if exists('syntax_on')
  syntax reset
endif

set termguicolors
let g:colors_name = '{{.Name}}'
let g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency = get(g:, '{{mustRegexFind "^[a-z]+" $.Name}}_transparency', 0)

{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
let g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} = get(g:, '{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}}', 0)
{{end -}}
{{end -}}

{{ $counter := 0 }}
if (has('termguicolors') && &termguicolors) || has('gui_running')
  let g:terminal_ansi_colors = [
    {{ range $index, $color := .Palette -}}
    {{if lt $counter 16 -}}
    \ "{{ upper $color }}",
    {{ $counter = add $counter 1 }}
    {{- end -}}
    {{- end -}}
  \]
endif

{{ $transparent_groups := list "Normal" "NormalFloat" "Pmenu" }}
{{range $index, $group := .Highlights.Editor}}
  {{- if get $group "link"}}
hi! link {{$group.name}} {{$group.link}}
  {{- else}}
hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
    {{- if has $group.name $transparent_groups }}
if g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency == 1
  hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{"NONE"}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
endif
    {{- end -}}
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
if g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$y}} == 1
  {{- range $z, $group := $group_vals}}
  {{- if get $group "link"}}
  hi! link {{$group.name}} {{$group.link}}
  {{- else}}
  hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
  {{- end -}}
{{end}}
endif
{{end -}}
{{end -}}

unlet g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency
{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
unlet g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}}
{{end -}}
{{end -}}`
}
