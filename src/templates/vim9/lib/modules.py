template = """vim9script
{% for key, groups in colorscheme.highlights.modules.items() %}

if g:{{ colorscheme.name }}.modules["{{ key }}"]
    {% for key, group in groups.items() %}
        {% if group["link"] is string -%}
            {{'  '}}hi! link {{key}} {{group["link"]}}
        {% else -%}
            {{'  '}}hi {{key}} guifg={{ colorscheme.palette[group['fg']] or "NONE" }} guibg={{ colorscheme.palette[group['bg']] or "NONE" }} gui=NONE cterm=NONE
        {% endif %}
    {% endfor %}
endif
{% endfor %}
"""
