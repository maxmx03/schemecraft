init_template="""
local config = require('{{ colorscheme.name }}.config')
local colors = require('{{ colorscheme.name }}.palette')
local color = require('{{ colorscheme.name }}.color')
local highlights = require('{{ colorscheme.name }}.highlights')

local M = {}

M.config = vim.deepcopy(config)
M.colors = vim.deepcopy(colors)

M.setup = function(opts)
  opts = opts or {}

  if not vim.tbl_isempty(opts) then
    M.config = vim.tbl_deep_extend('force', M.config, opts)

    if type(opts.on_colors) == 'function' then
      M.colors = vim.tbl_deep_extend('force', M.colors, opts.on_colors(colors, color))
    end
  end
end

M.load = function()
  if vim.g.colors_name then
    vim.cmd('hi clear')
  end

  if vim.fn.exists('syntax_on') then
    vim.cmd('syntax reset')
  end

  vim.o.termguicolors = true
  vim.g.colors_name = '{{ colorscheme.name }}'

  highlights.load(M.colors, M.config)
end

return M
"""
