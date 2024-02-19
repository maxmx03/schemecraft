package template

func Config() string {
	return `---@class {{mustRegexFind "^[a-z]+" .Name}}.config
---@field transparent? boolean
---@field on_highlights? fun(colors: {{mustRegexFind "^[a-z]+" .Name}}.palette, color: table): {{mustRegexFind "^[a-z]+" .Name}}.highlights
---@field on_colors? fun(colors: {{mustRegexFind "^[a-z]+" .Name}}.palette, color: table): {{mustRegexFind "^[a-z]+" .Name}}.palette
return {
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

func Root() string {
	return `---@class {{mustRegexFind "^[a-z]+" .Name}}
---@field config {{mustRegexFind "^[a-z]+" .Name}}.config
---@field setup fun(config: {{mustRegexFind "^[a-z]+" .Name}}.config)
---@field load fun(theme: string)
local M = {}

M.config = require("{{mustRegexFind "^[a-z]+" .Name}}.config")

M.setup = function(config)
	M.config = vim.tbl_extend("force", M.config, config)
end

M.load = function(theme)
  theme = theme or ""
	if vim.g.colors_name then
		vim.cmd("hi clear")
	end

	if vim.fn.exists("syntax_on") then
		vim.cmd("syntax reset")
	end

	vim.g.colors_name = theme
  local ok, highlights = pcall(require, "{{mustRegexFind "^[a-z]+" .Name}}.highlights." .. theme)
  if not ok then
    highlights = require("{{regexFind "^[a-z]+" .Name}}.highlights")
  end

  local colors = {}
  ok, colors = pcall(require, "{{mustRegexFind "^[a-z]+" .Name}}.palette" .. theme)
  if not ok then
    colors = require("{{mustRegexFind "^[a-z]+" .Name}}.palette")
  end

  highlights.set_highlight(colors, M.config)
end

return M`
}

func Utils() string {
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

func Color() string {
	return `---@class {{mustRegexFind "^[a-z]+" .Name}}.color
---@field hex_to_rgb fun(hex: string): number, number, number
---@field rgb_to_hex fun(red: number, green: number, blue: number): string
---@field darken fun(hex: string, percentage: number): string
---@field lighten fun(hex: string, percentage: number): string
---@field shade fun(hex: string, i: number): string
---@field tint fun(hex: string, i: number): string
---@field blend fun(fg: string, bg: string, alpha: number): string
local M = {}

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

---@param fg number
---@param bg number
---@param alpha number
---@return number
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
