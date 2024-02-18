
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
	nvim_set_hl("Normal", {fg = colors.base0,bg = colors.base03,},config)
	nvim_set_hl("NormalNC", { link = "Normal" })
  -- SYNTAX :h group-name
  nvim_set_hl("Function", { link = "Constant" })
  
  nvim_set_hl("Comment", {fg = colors.base01,bg = colors.base03,})
  nvim_set_hl("Constant", {fg = colors.purple,bold = true,italic = true,})
  nvim_set_hl("Variable", { link = "Constant" })
  -- PLUGINS
  if config.plugins["treesitter"] then
    nvim_set_hl("@variable.parameter", {fg = colors.orange,})
    nvim_set_hl("@comment", {fg = colors.base01,})
  end
  
  if config.plugins["tree"] then
    nvim_set_hl("directory", {fg = colors.green,})
    nvim_set_hl("directory", {fg = colors.green,})
  end
  if config.on_highlight then
		local color = require("dracula.color")
		local highlights = config.on_highlight(colors, color)

		for group_name, group_val in pairs(highlights) do
			local hl = nvim_get_hl({ name = group_name, link = true })
			local val = vim.tbl_extend("force", hl, group_val)
			vim.api.nvim_set_hl(0, group_name, val)
		end
	end
end

return M