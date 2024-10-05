package template

import "fmt"

func LuaReadme() string {
	var readme = `# {{title (mustRegexFind "^[a-z]+" .Name)}}

![schemecraft](https://github.com/maxmx03/schemecraft/assets/50273941/ee682aae-00cb-4282-ba24-3d9621a430a3)

## Installation

To install {{title .Name}}, you need a plugin manager.
In the example, bellow we are going to use vim-plug.

%v

## Transparency

vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_transparency = true

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
vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} = true
{{end -}}
{{end -}}
`
	plugins += "```"

	return fmt.Sprintf(readme, installBlockCode, plugins)
}

func LuaColors() string {
	return `
-- {{.Name}} Colorscheme for vim
-- Url: {{.Repo}}
-- Maintainer: {{.Author}} <{{.Contact}}>
-- License: {{.License}}

if vim.g.loaded_{{mustRegexFind "^[a-z]+" .Name}} then
  return
end

vim.g.loaded_{{mustRegexFind "^[a-z]+" .Name}} = true

vim.cmd.hi 'clear'

if vim.fn.exists 'syntax_on' then
  vim.cmd.syntax 'reset'
end

vim.o.termguicolors = true
vim.g.colors_name = '{{.Name}}'
vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_transparency = vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_transparency or false

{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} = vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}} or false
{{end -}}
{{end -}}

{{ $counter := 0 }}
vim.cmd [[
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
]]

{{ $transparent_groups := list "Normal" "NormalFloat" "Pmenu" }}
{{range $index, $group := .Highlights.Editor}}
  {{- if get $group "link"}}
vim.api.nvim_set_hl(0, "{{$group.name}}", { link = "{{$group.link}}" })
  {{- else}}

	vim.api.nvim_set_hl(0,"{{$group.name}}", {
    {{- if get $group "fg"}}fg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "fg")))}},{{end}}
    {{- if get $group "bg"}}bg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "bg")))}},{{end}}
    {{- if get $group "blend"}}blend = {{$group.blend}},{{end}}
    {{- if get $group "bold"}}bold = {{$group.bold}},{{end}}
    {{- if get $group "standout"}}standout = {{$group.standout}},{{end}}
    {{- if get $group "underline"}}underline = {{$group.underline}},{{end}}
    {{- if get $group "undercurl"}}undercurl = {{$group.undercurl}},{{end}}
    {{- if get $group "underdouble"}}underdouble = {{$group.underdouble}},{{end}}
    {{- if get $group "underdotted"}}underdotted = {{$group.underdotted}},{{end}}
    {{- if get $group "underdashed"}}underdashed = {{$group.underdashed}},{{end}}
    {{- if get $group "strikethrough"}}strikethrough = {{$group.strikethrough}},{{end}}
    {{- if get $group "italic"}}italic = {{$group.italic}},{{end}}
    {{- if get $group "reverse"}}reverse = {{$group.reverse}},{{end}}
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}})
    {{- if has $group.name $transparent_groups }}
if vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_transparency then
	vim.api.nvim_set_hl(0,"{{$group.name}}", {
    {{- if get $group "fg"}}fg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "fg")))}},{{end}}
    {{- if get $group "bg"}}bg = "NONE",{{end}}
    {{- if get $group "blend"}}blend = {{$group.blend}},{{end}}
    {{- if get $group "bold"}}bold = {{$group.bold}},{{end}}
    {{- if get $group "standout"}}standout = {{$group.standout}},{{end}}
    {{- if get $group "underline"}}underline = {{$group.underline}},{{end}}
    {{- if get $group "undercurl"}}undercurl = {{$group.undercurl}},{{end}}
    {{- if get $group "underdouble"}}underdouble = {{$group.underdouble}},{{end}}
    {{- if get $group "underdotted"}}underdotted = {{$group.underdotted}},{{end}}
    {{- if get $group "underdashed"}}underdashed = {{$group.underdashed}},{{end}}
    {{- if get $group "strikethrough"}}strikethrough = {{$group.strikethrough}},{{end}}
    {{- if get $group "italic"}}italic = {{$group.italic}},{{end}}
    {{- if get $group "reverse"}}reverse = {{$group.reverse}},{{end}}
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}})
end
    {{- end -}}
  {{- end -}}
{{- end -}}

{{range $index, $group := .Highlights.Syntax}}
  {{- if get $group "link"}}
vim.api.nvim_set_hl(0, "{{$group.name}}", { link = "{{$group.link}}" })
  {{- else}}
	vim.api.nvim_set_hl(0,"{{$group.name}}", {
    {{- if get $group "fg"}}fg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "fg")))}},{{end}}
    {{- if get $group "bg"}}bg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "bg")))}},{{end}}
    {{- if get $group "blend"}}blend = {{$group.blend}},{{end}}
    {{- if get $group "bold"}}bold = {{$group.bold}},{{end}}
    {{- if get $group "standout"}}standout = {{$group.standout}},{{end}}
    {{- if get $group "underline"}}underline = {{$group.underline}},{{end}}
    {{- if get $group "undercurl"}}undercurl = {{$group.undercurl}},{{end}}
    {{- if get $group "underdouble"}}underdouble = {{$group.underdouble}},{{end}}
    {{- if get $group "underdotted"}}underdotted = {{$group.underdotted}},{{end}}
    {{- if get $group "underdashed"}}underdashed = {{$group.underdashed}},{{end}}
    {{- if get $group "strikethrough"}}strikethrough = {{$group.strikethrough}},{{end}}
    {{- if get $group "italic"}}italic = {{$group.italic}},{{end}}
    {{- if get $group "reverse"}}reverse = {{$group.reverse}},{{end}}
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}})
{{- end -}}
{{- end -}}

{{- range $x, $groups := .Highlights.Plugins}}
  {{- range $y, $group_vals := $groups}}
if vim.g.{{mustRegexFind "^[a-z]+" $.Name}}_{{$y}} then
  {{- range $z, $group := $group_vals}}
  {{- if get $group "link"}}
vim.api.nvim_set_hl(0, "{{$group.name}}", { link = "{{$group.link}}" })
  {{- else}}
	vim.api.nvim_set_hl(0,"{{$group.name}}", {
    {{- if get $group "fg"}}fg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "fg")))}},{{end}}
    {{- if get $group "bg"}}bg = {{printf "'%s'" (default "NONE" (get $.Palette (get $group "bg")))}},{{end}}
    {{- if get $group "blend"}}blend = {{$group.blend}},{{end}}
    {{- if get $group "bold"}}bold = {{$group.bold}},{{end}}
    {{- if get $group "standout"}}standout = {{$group.standout}},{{end}}
    {{- if get $group "underline"}}underline = {{$group.underline}},{{end}}
    {{- if get $group "undercurl"}}undercurl = {{$group.undercurl}},{{end}}
    {{- if get $group "underdouble"}}underdouble = {{$group.underdouble}},{{end}}
    {{- if get $group "underdotted"}}underdotted = {{$group.underdotted}},{{end}}
    {{- if get $group "underdashed"}}underdashed = {{$group.underdashed}},{{end}}
    {{- if get $group "strikethrough"}}strikethrough = {{$group.strikethrough}},{{end}}
    {{- if get $group "italic"}}italic = {{$group.italic}},{{end}}
    {{- if get $group "reverse"}}reverse = {{$group.reverse}},{{end}}
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}})
  {{- end -}}
{{end}}
end
{{end -}}
{{end -}}
vim.defer_fn(function()
vim.api.nvim_del_var("{{mustRegexFind "^[a-z]+" $.Name}}_transparency")
{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
vim.api.nvim_del_var("{{mustRegexFind "^[a-z]+" $.Name}}_{{$pluginName}}")
{{end -}}
{{end -}}
end, 800)`
}
