![schemecraft](https://github.com/maxmx03/schemecraft/assets/50273941/ee682aae-00cb-4282-ba24-3d9621a430a3)

<div align="center"> 
    
![Neovim](https://img.shields.io/badge/Neovim-v0.9.1+-blue?NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white)
![Vim](https://img.shields.io/badge/Vim-9-blue?NeoVim-%2357A143.svg?&style=for-the-badge&logo=vim&logoColor=white)
![LICENSE](https://shields.io/badge/LICENSE-MIT-orange?style=for-the-badge)

</div>

Schemecraft is a versatile colorscheme generator designed for Neovim and Vim9.
It provides support for multiple variants, allowing users to customize and
create unique color schemes tailored to their preferences.

## Getting Started

## Build

To use Schemecraft, start by building the executable:

```bash
go build
```

## Usage

Generate a colorscheme by providing a base JSON file (or YAML)

### Create

#### example 1

```bash
schemecraft create scheme.json
```

> output: `build/`

#### example 2

```bash
schemecraft create scheme.yml scheme-light.yml
```

### Update

Update your colorscheme

```bash
schemecraft update nvim scheme.json
schemecraft update vim scheme.json
```

> output: current directory

## Contributing

Feel free to contribute to Schemecraft by opening issues or pull requests.
Your feedback and suggestions are highly appreciated.
