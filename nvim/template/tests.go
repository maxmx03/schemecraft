package template

func LoadSpec() string {
	return `local nvim_get_hl = require("{{mustRegexFind "^[a-z]+" .Name}}.utils").nvim_get_hl

describe("{{.Name}}.load", function()
	setup(function()
		vim.cmd.colorscheme("{{mustRegexFind "^[a-z]+" .Name}}")
	end)

	test("name", function()
		local expected = "{{mustRegexFind "^[a-z]+" .Name}}"
		assert.equal(expected, vim.g.colors_name)
	end)

	test("palette", function()
		local colors = require("{{mustRegexFind "^[a-z]+" .Name}}.palette")
		assert.True(type(colors) == "table")
	end)

	test("set highlight", function()
		local string = nvim_get_hl("String")
		local colors = require("{{mustRegexFind "^[a-z]+" .Name}}.palette")
		local expected = colors.orange
		assert.equal(expected, string.fg)
	end)

	test("default config", function()
		local config = require("{{mustRegexFind "^[a-z]+" .Name}}.config")
		assert.True(type(config) == "table")
	end)
end)`
}

func SetupSpec() string {
	return `local nvim_get_hl = require("{{mustRegexFind "^[a-z]+" .Name}}.utils").nvim_get_hl
local colors = require("{{mustRegexFind "^[a-z]+" .Name}}.palette")

describe("{{mustRegexFind "^[a-z]+" .Name}}.setup", function()
	setup(function()
		require("{{mustRegexFind "^[a-z]+" .Name}}").setup({
			transparent = true,
			on_colors = function()
				return {
					yellow = "#ffffff",
				}
			end,
			on_highlight = function()
				return {
					Type = { bg = "#ffffff" },
				}
			end,
			plugins = {
				["nvim-treesitter"] = true,
			},
		})
		vim.cmd("colorscheme {{mustRegexFind "^[a-z]+" .Name}}")
	end)

	test("transparent", function()
		local normal = nvim_get_hl("Normal")
		assert.falsy(normal.bg)
	end)

	test("on_colors", function()
		local expected = nvim_get_hl("Function").fg
		assert.True(expected ~= colors.yellow)
	end)

	test("on_highlight", function()
		local user_group = nvim_get_hl("Type")
		local expected = { fg = colors.green, bg = "#ffffff" }
		assert.are.same(expected, user_group)
	end)

	test("plugins", function()
		local treesitter = nvim_get_hl("@variable")
		assert.True(colors.base0 == treesitter.fg)
	end)
end)`
}
