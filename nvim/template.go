package nvim

func configTemplate() string {
	return `return {
  transparent = {{.Config.Transparent}},
  on_highlights = {{default "nil" .Config.OnHighlights}},
  on_colors = {{default "nil" .Config.OnColors}},
  plugins = {
  {{range $index, $value := .Config.Plugins -}}
  {{"  "}}["{{$index}}"] = {{default true $value}},
  {{end -}}
  }
}`
}

func paletteTemplate() string {
	return `return {
  {{range $index, $value := .Palette -}}
  {{"  "}}{{$index}} = "{{upper $value}}",
  {{end -}}
}`
}

func initTemplate() string {
	return `local M = {}
M.config = require("{{.Name}}.config")
M.colors = require("{{.Name}}.palette")

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

M.setup = function(config)
	M.config = vim.tbl_extend("force", M.config, config)
end

M.load = function()
	if vim.g.colors_name then
		vim.cmd("hi clear")
	end

	if vim.fn.exists("syntax_on") then
		vim.cmd("syntax reset")
	end

	vim.o.termguicolors = true
	vim.g.colors_name = "{{.Name}}"
	M.set_highlight(M.colors, M.config)
end

M.set_highlight = function(colors, config)
	if config.on_colors then
		colors = vim.tbl_extend("force", colors, config.on_colors(colors))
	end
  {{- range $index, $group := .Highlights.Editor -}}
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
		local highlights = M.config.on_highlight(colors, color)

		for group_name, group_val in pairs(highlights) do
			local hl = nvim_get_hl({ name = group_name, link = true })
			local val = vim.tbl_extend("force", hl, group_val)
			vim.api.nvim_set_hl(0, group_name, val)
		end
	end
end

return M
  `
}

func utilsTemplate() string {
	return `local function to_hex(decimal)
  if type(decimal) ~= "number" then
    return decimal
  end
  local hex = string.format("#%06x", decimal)
  return hex
end

local function nvim_get_hl(name)
  local output = vim.api.nvim_get_hl(0, { name = name, link = false })

  local t = {}

  for key, value in pairs(output) do
    t[key] = to_hex(value)
  end

  return t
end

return {
  nvim_get_hl = nvim_get_hl,
}`
}

func colorTemplate() string {
	return `local M = {}

function M.hex_to_rgb(hex)
  local red = tonumber(hex:sub(2, 3), 16)
  local green = tonumber(hex:sub(4, 5), 16)
  local blue = tonumber(hex:sub(6, 7), 16)
  return red, green, blue
end

function M.rgb_to_hex(red, green, blue)
  return string.format('#%02x%02x%02x', red, green, blue)
end

function M.darken(hex, percentage)
  local red, green, blue = M.hex_to_rgb(hex)
  local i = 1
  local result = {
    red = red * (1 - (percentage / 100) * i),
    green = green * (1 - (percentage / 100) * i),
    blue = blue * (1 - (percentage / 100) * i),
  }
  return M.rgb_to_hex(result.red, result.green, result.blue)
end

function M.lighten(color, percentage)
  local red, green, blue = M.hex_to_rgb(color)
  local i = 1
  local result = {
    red = red + (255 - red) * i * (percentage / 100),
    green = green + (255 - green) * i * (percentage / 100),
    blue = blue + (255 - blue) * i * (percentage / 100),
  }
  return M.rgb_to_hex(result.red, result.green, result.blue)
end

function M.shade(hex, i)
  local red, green, blue = M.hex_to_rgb(hex)
  local result = {
    red = red * (1 - 0.1 * i),
    green = green * (1 - 0.1 * i),
    blue = blue * (1 - 0.1 * i),
  }
  return M.rgb_to_hex(result.red, result.green, result.blue)
end

function M.tint(hex, i)
  local red, green, blue = M.hex_to_rgb(hex)
  local result = {
    red = red + (255 - red) * i * 0.1,
    green = green + (255 - green) * i * 0.1,
    blue = blue + (255 - blue) * i * 0.1,
  }
  return M.rgb_to_hex(result.red, result.green, result.blue)
end

local function blend_channel(fg, bg, alpha)
  local ret = alpha * fg + ((1 - alpha) * bg)
  local min = math.min
  local max = math.max
  local floor = math.floor
  return floor(min(max(0, ret), 255) + 0.5)
end

function M.blend(fg, bg, alpha)
  local red_bg, green_bg, blue_bg = M.hex_to_rgb(bg)
  local red_fg, green_fg, blue_fg = M.hex_to_rgb(fg)
  return M.rgb_to_hex(
    blend_channel(red_fg, red_bg, alpha),
    blend_channel(green_fg, green_bg, alpha),
    blend_channel(blue_fg, blue_bg, alpha)
  )
end

return M`
}
