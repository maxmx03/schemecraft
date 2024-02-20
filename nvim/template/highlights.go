package template

func Highlights() string {
	return `---@class {{mustRegexFind "^[a-z]+" .Name}}.highlights
{{- range $index, $group := .Highlights.Editor}}
---@field {{$group.name}} table
{{- end}}
{{- range $index, $group := .Highlights.Syntax}}
---@field {{$group.name}} table
{{- end}}
{{- range $x, $groups := .Highlights.Plugins}}
{{- range $y, $group_vals := $groups}}
{{- range $z, $group := $group_vals}}
{{- if not (hasPrefix "@" $group.name)}}
---@field {{$group.name}} table
{{- end}}
{{- end}}
{{- end}}
{{- end}}

local M = {}

local function nvim_get_hl(opts)
  local hl = vim.api.nvim_get_hl(0, opts)

  if hl.link then
    return nvim_get_hl({ name = hl.link })
  end

  return hl
end

local nvim_set_hl = function(group_name, group_val, config)
  local val = { fg = "NONE", bg = "NONE" }

  if not group_val.link then
    if config and config.transparent then
      group_val.bg = "NONE"
    end
    val = vim.tbl_extend("force", val, group_val)
    vim.api.nvim_set_hl(0, group_name, val)
  else
    vim.api.nvim_set_hl(0, group_name, group_val)
  end
end

M.set_highlight = function(colors, config)
	if config.on_colors then
		colors = vim.tbl_extend("force", colors, config.on_colors(colors))
	end
  -- EDITOR :h highlight-groups
  {{- range $index, $group := .Highlights.Editor -}}
  {{if get $group "link"}}
	nvim_set_hl("{{$group.name}}", { link = "{{$group.link}}" })
  {{else}}
	nvim_set_hl("{{$group.name}}", {
    {{- if get $group "fg"}}fg = colors.{{$group.fg}},{{end}}
    {{- if get $group "bg"}}bg = colors.{{$group.bg}},{{end}}
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
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}}
    {{- if get $group "transparent" }},config{{end}})
  {{- end -}}
  {{- end -}}
  -- SYNTAX :h group-name
  {{- range $index, $group := .Highlights.Syntax -}}
  {{if get $group "link"}}
  nvim_set_hl("{{$group.name}}", { link = "{{$group.link}}" })
  {{else}}
  nvim_set_hl("{{$group.name}}", {
    {{- if get $group "fg"}}fg = colors.{{$group.fg}},{{else}}{{end}}
    {{- if get $group "bg"}}bg = colors.{{$group.bg}},{{else}}{{end}}
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
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}}
    {{- if get $group "transparent" }},config{{end}})
  {{- end -}}
  {{- end -}}
   -- PLUGINS
  {{- range $x, $groups := .Highlights.Plugins}}
  {{- range $y, $group_vals := $groups}}
  if config.plugins["{{$y}}"] then
  {{- range $z, $group := $group_vals}}
  {{- if get $group "link" -}}
    nvim_set_hl("{{$group.name}}", { link = "{{$group.link}}" })
  {{else}}
    nvim_set_hl("{{$group.name}}", {
    {{- if get $group "fg"}}fg = colors.{{$group.fg}},{{else}}{{end}}
    {{- if get $group "bg"}}bg = colors.{{$group.bg}},{{else}}{{end}}
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
    {{- if get $group "nocombine"}}nocombine = {{$group.nocombine}},{{end}}}
    {{- if get $group "transparent" }},config{{end}})
  {{- end -}}
  {{end}}
  end
  {{end -}}
  {{end -}}

	if config.on_highlight then
		local color = require("{{.Name}}.color")
		local highlights = config.on_highlight(colors, color)

		for group_name, group_val in pairs(highlights) do
			local hl = nvim_get_hl({ name = group_name, link = true })
			local val = vim.tbl_extend("force", hl, group_val)
			vim.api.nvim_set_hl(0, group_name, val)
		end
	end
end

return M`
}
