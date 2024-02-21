package template

import "fmt"

func NixShell() string {
	return `{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  nativeBuildInputs = with pkgs.buildPackages; [luajitPackages.vusted stylua];
}`
}

func DockerFile() string {
	return `FROM archlinux
WORKDIR /{{lower .Name}}
COPY . .
RUN pacman -Syu --noconfirm
RUN pacman -S luarocks neovim gcc lua51 --noconfirm
RUN luarocks install luacheck && luarocks --lua-version=5.1 install vusted
CMD [ "vusted", "./tests" ]`
}

func DockerCompose() string {
	return `services:
  {{lower .Name}}:
    build: .
    volumes:
      - .:/{{lower .Name}}`
}

func Readme() string {
	var readme = `# {{.Name}}

Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim
labore culpa sint ad nisi Lorem pariatur.

## Installation

%v

%v

## Api

%v
`

	var luarc = "```json"
	luarc += `
{
  "workspace": {
    "library": [
      "~/.local/share/nvim/lazy/{{.Name}}"
    ],
  },
}
`
	luarc += "```"

	var installBlockCode = "```lua"
	installBlockCode += `
return {
  {
    '{{.Repo}}',
    lazy = false,
    priority = 1000,
    config = function ()
      ---@type {{.Name}}
      require "{{.Name}}".setup({
      transparent = {{.Config.Transparent}},
      on_colors = function (colors, color)
        ---@type {{.Name}}.colors
        return {
          -- override or create new colors
          mycolor = "#ffffff",
        }
      end,
      on_highlights = function (colors, color)
        ---@type {{.Name}}.highlights
        return {
          ---@type vim.api.keyset.highlight
          Normal = { fg = colors.mycolor }
        }
      end,
      plugins = {
      {{range $index, $plugins := .Highlights.Plugins -}}
      {{range $pluginName, $pluginConfigs:= $plugins -}}
      {{"  "}}["{{$pluginName}}"] = true,
      {{end -}}
      {{end -}}
      }
      })
      vim.cmd.colorscheme '{{.Name}}'
    end
  }
}
`
	installBlockCode += "```"

	var api = "```lua"
	api += `
local colors = require '{{.Name}}.colors'
local color = require '{{.Name}}.color'
local darken = color.darken
local lighten = color.lighten
local blend = color.blend
local shade = color.shade
local tint = color.tint

-- example 1: get shades
for i = 1, 10, 1 do
    print(shade(colors.yellow, i))
end

for i = 1, 100, 10 do
    print(darken(colors.yellow, i))
end

-- example 2: get tints
for i = 1, 10, 1 do
    print(tint(colors.yellow, i))
end

for i = 1, 100, 10 do
    print(lighten(colors.yellow, i))
end

-- example 3: blend color
local new_color = blend(colors.yellow, colors.base03, 0.2)
`
	api += "```"

	return fmt.Sprintf(readme, luarc, installBlockCode, api)
}
