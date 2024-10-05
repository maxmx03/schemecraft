![schemecraft](https://github.com/maxmx03/schemecraft/assets/50273941/ee682aae-00cb-4282-ba24-3d9621a430a3)

<div align="center"> 
    
![Neovim](https://img.shields.io/badge/Neovim-v0.10.1+-blue?NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white)
![Vim](https://img.shields.io/badge/Vim-green?NeoVim-%2357A143.svg?&style=for-the-badge&logo=vim&logoColor=white)
![LICENSE](https://shields.io/badge/LICENSE-MIT-orange?style=for-the-badge)

</div>

Schemecraft is a versatile colorscheme generator designed for Neovim and Vim.
It provides support for multiple variants, allowing users to customize and
create unique color schemes tailored to their preferences.

## Getting Started

## Installation

```bash
go install github.com/maxmx03/schemecraft@latest
```

## Usage

Generate a colorscheme by providing a base Yaml file

### Create

#### example 1

```bash
schemecraft scheme.json
```

> output: `build/`

#### example 2

```bash
schemecraft scheme.yml scheme-light.yml
```

### Template YAML

Check the example folder for a reference.


```yaml
specification:
  identifier: &identifier
    fg: fg

name: Solarized
author: Max
repo: https://github.com/maxmx03/solarized
contact: ""
config: # you keep this as default
  transparent: false
  plugins: ~
palette:
  fg: "#FEECE2"
  bg: "#374259"
  comment: "#545B77"
  blue: "#B5C0D0"
highlights:
  editor: # :h highlihts-groups
    - name: Normal
      fg: fg
      bg: bg
      transparent: true
  syntax: # :h group-name
    - name: Comment
      fg: comment
      italic: true
      gui: italic,bold
plugins:
  treesitter:
    - name: "@variable"
      <<: *identifier
    - name: "@variable.builtin"
      link: Constant
```

## Schemecraft Colorschemes

- [dracula](https://github.com/maxmx03/dracula.nvim)

## Contributing

Feel free to contribute to Schemecraft by opening issues or pull requests.
Your feedback and suggestions are highly appreciated.
