template = """vim9script
{% for key, group in colorscheme.highlights.syntax.items() -%}
    {% if group["link"] is string -%}
        hi! link {{key}} {{group["link"]}}
    {% else -%}
        hi {{key}} guifg={{ colorscheme.palette[group['fg']] or "NONE" }} guibg={{ colorscheme.palette[group['bg']] or "NONE" }} gui=NONE cterm=NONE
    {% endif %}
{% endfor %}
"""
