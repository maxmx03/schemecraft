template = """vim9script
# {{ colorscheme.name }} colorscheme
# Repo: {{ colorscheme.repo }}
# Maintainer: {{ colorscheme.author }} <"{{ colorscheme.contact }}">
# License: {{ colorscheme.license }}

if exists("g:loaded_{{ colorscheme.name }}")
  finish
endif

g:loaded_{{ colorscheme.name }} = 1

hi clear

if exists('syntax_on')
  syntax reset
endif

set termguicolors
g:colors_name = '{{ colorscheme.name }}'

import "../lib/config.vim"
import "../lib/editor.vim"
import "../lib/terminal.vim"
import "../lib/syntax.vim"
import "../lib/modules.vim"
"""
