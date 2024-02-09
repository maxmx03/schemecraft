template="""
function to_hex(decimal)
  if type(decimal) ~= 'number' then
    return decimal
  end
  local hex = string.format('#%06x', decimal)
  return hex
end

function nvim_get_hl(name)
  local output = vim.api.nvim_get_hl(0, { name = name })

  if output and output.link then
    return output
  end

  local t = {}

  for key, value in pairs(output) do
    t[key] = to_hex(value)
  end

  return t
end

describe('{{ colorscheme.name }}', function()
  setup(function()
    require('{{ colorscheme.name }}').setup({})
    vim.cmd.colorscheme('{{ colorscheme.name }}')
  end)

  test('colors_name', function()
    assert.equal('{{ colorscheme.name }}', vim.g.colors_name)
  end)

  test('syntax highlights', function()
    local output = nvim_get_hl('String')
    local colors = require('{{ colorscheme.name }}.palette')
    assert.equal(colors.yellow, output.fg)
  end)
end)
"""
