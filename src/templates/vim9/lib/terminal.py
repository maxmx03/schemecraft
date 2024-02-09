template="""vim9script
if (has('termguicolors') && &termguicolors) || has('gui_running')
  g:terminal_ansi_colors = [
    {% for key, value in colorscheme.palette.items() %}
        "{{ value }}",
    {% endfor %}
  ]
endif
"""
