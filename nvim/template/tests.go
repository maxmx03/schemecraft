package template

func LoadSpec() string {
	return `local nvim_get_hl = require("{{.Name}}.utils").nvim_get_hl

describe("{{.Name}}.load", function()
	setup(function()
		vim.cmd.colorscheme("{{.Name}}")
	end)

	test("name", function()
		local expected = "{{.Name}}"
		assert.equal(expected, vim.g.colors_name)
	end)

	test("palette", function()
		local colors = require("{{.Name}}.palette")
		assert.True(type(colors) == "table")
	end)

  test("default config", function()
		local config = require("{{.Name}}.config")
		assert.True(type(config) == "table")
	end)

  test("set_highlight", function()
  {{- range $index, $group := .Highlights.Syntax}}
   local {{$group.name}} = nvim_get_hl("{{$group.name}}")
   {{- if $group.fg }}
   assert.True("{{upper (get $.Palette (print $group.fg))}}" == {{$group.name}}.fg)
   {{end}}
   {{- if $group.bg }}
   assert.True("{{upper (get $.Palette (print $group.bg))}}" == {{$group.name}}.bg)
   {{end}}
   {{- if $group.link }}
   assert.True("{{$group.link}}" == {{$group.name}}.link)
   {{end}}
  {{- end -}}
  end)
end)`
}

func SetupSpec() string {
	return `local nvim_get_hl = require("{{.Name}}.utils").nvim_get_hl

describe("{{.Name}}.setup", function()
	setup(function()
		require("{{.Name}}").setup({
			transparent = true,
			on_colors = function()
				return {
					mycolor = "#ffffff",
				}
			end,
			on_highlights = function(colors)
				return {
					CustomHighlight = { fg = colors.mycolor },
				}
			end,
		})
		vim.cmd("colorscheme {{.Name}}")
	end)

	test("transparent", function()
		local normal = nvim_get_hl("Normal")
		assert.falsy(normal.bg)
	end)

	test("on_colors and on_highlight", function()
		local user_group = nvim_get_hl("CustomHighlight")
		local expected = { fg = "#FFFFFF" }
		assert.are.same(expected, user_group)
	end)
end)`
}
