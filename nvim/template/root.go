package template

func NixShell() string {
	return `{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  nativeBuildInputs = with pkgs.buildPackages; [luajitPackages.vusted];
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

func Docs() string {
	return `*{{.Name}}.nvim.txt*         For NVIM v0.9.1         Last change: {{now | date "2006-01-02"}}

==============================================================================
Table of Contents                               *{{.Name}}.nvim-table-of-contents*

1. {{title .Name}}                                                    |{{lower .Name}}.nvim-{{lower .Name}}|
  - Installation                               |{{lower .Name}}.nvim-{{lower .Name}}-installation|
  - Configuration                             |{{lower .Name}}.nvim-{{lower .Name}}-configuration|
  - Commands                                       |{{lower .Name}}.nvim-{{lower .Name}}-commands|
  - Docs                                               |{{lower .Name}}.nvim-{{lower .Name}}-docs|
  - Api                                                 |{{lower .Name}}.nvim-{{lower .Name}}-api|

==============================================================================
1. {{title .Name}}                                                    *{{lower .Name}}.nvim-{{lower .Name}}*


- |{{lower .Name}}.nvim-{{lower .Name}}-colorscheme|
    - |{{lower .Name}}.nvim-installation|
    - |{{lower .Name}}.nvim-configuration|
    - |{{lower .Name}}.nvim-commands|
    - |{{lower .Name}}.nvim-docs|
    - |{{lower .Name}}.nvim-api|

    {{title .Name}} colorscheme is based on onedarkpro <olimorris/onedarkpro.nvim> and
github-theme <projekt0n/github-nvim-theme>


INSTALLATION                                   *{{lower .Name}}.nvim-{{lower .Name}}-installation*

Neovim users

>lua
    return {
        {
            '{{.Repo}}',
            lazy = false,
            priority = 1000,
            config = function ()
                vim.cmd.colorscheme '{{.Name}}'
            end
        },
        {
            'nvim-lualine/lualine.nvim',
            config = function()
              require('lualine').setup {
                options = {
                  theme = '{{.Name}}',
                },
              }
            end,
        },
    }
<

Vim9 Users

>vim
    Plug '{{.Repo}}', { 'branch': 'vim' }
<


CONFIGURATION                                 *{{lower .Name}}.nvim-{{lower .Name}}-configuration*

>lua
    local {{lower .Name}} = require '{{.Name}}'
    {{lower .Name}}.setup {
        transparent = false,
        plugins = {
	{{range $index, $value := .Config.Plugins -}}
	{{"  "}}["{{$index}}"] = {{default true $value}},
	{{end -}}
        },
        on_colors = function(colors, color)
            return {}
        end
        on_highlights = function (colors, color)
            local darken = color.darken
            local lighten = color.lighten
            local shade = color.shade
            local tint = color.tint
            local blend = color.blend
    
            return {
                ---@type vim.api.keyset.highlight
                Function = { italic = false },
                Identifier = { bold = true },
                LineNr = {
                    fg = colors.pink,
                    bg = darken(colors.pink, 50)
                },
                LineNr = {
                    fg = colors.pink,
                    bg = lighten(colors.pink, 50)
                },
                LineNr = {
                    fg = colors.pink,
                    bg = shade(colors.pink, 5)
                },
                LineNr = {
                    fg = colors.pink,
                    bg = tint(colors.pink, 5)
                },
                LineNr = {
                    fg = colors.pink,
                    bg = blend(colors.pink, colors.base03, 0.2)
                },
            }
        end
    }
<


COMMANDS                                           *{{lower .Name}}.nvim-{{lower .Name}}-commands*

{{title .Name}} colors


DOCS                                                   *{{lower .Name}}.nvim-{{lower .Name}}-docs*

>lua
    h {{lower .Name}}
<


API                                                     *{{lower .Name}}.nvim-{{lower .Name}}-api*

>lua
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
<

Generated by panvimdoc <https://github.com/kdheepak/panvimdoc>

vim:tw=78:ts=8:noet:ft=help:norl:
`
}
