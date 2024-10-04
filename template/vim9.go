package template

import "fmt"

func Vim9Readme() string {
	var readme = `# {{title (mustRegexFind "^[a-z]+" .Name)}}

![schemecraft](https://github.com/maxmx03/schemecraft/assets/50273941/ee682aae-00cb-4282-ba24-3d9621a430a3)

## Installation

To install {{title .Name}}, you need a plugin manager.
In the example, bellow we are going to use vim-plug.

## Transparency

g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency = true

%v

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
g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} = true
{{end -}}
{{end -}}
`
	plugins += "```"

	return fmt.Sprintf(readme, installBlockCode, plugins)
}

func Vim9Colors() string {
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
g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency = true

{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} = get(g:, '{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}}', false)
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

{{ $transparent_groups := list "Normal" "NormalFloat" "Pmenu" }}
{{range $index, $group := .Highlights.Editor}}
  {{- if get $group "link"}}
hi! link {{$group.name}} {{$group.link}}
  {{- else}}
hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{default "NONE" (get $.Palette (get $group "bg"))}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
    {{- if has $group.name $transparent_groups }}
if g:{{mustRegexFind "^[a-z]+" $.Name}}_transparency
  hi {{$group.name}} guifg={{default "NONE" (get $.Palette (get $group "fg"))}} guibg={{"NONE"}} gui={{default "NONE" $group.gui}} cterm={{default "NONE" $group.gui}}
endif
    {{- end -}}
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
if g:{{mustRegexFind "^[a-z]+" $.Name}}_{{$y}}
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
