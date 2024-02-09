template = """
local True = true
local False = false

local M = {}

M.load = function(opts)
 local hl = opts.hl
 local colors = opts.colors
 local transparent = opts.config.transparent

{% for key, value in colorscheme.highlights.editor.items() %}
    {% if key in ["Normal", "CursorLineNr", "LineNr", "SignColumn", "WinSeparator"] %}
        hl("{{ key }}",
        {% if value is mapping %}
            {
            {% for subkey, subvalue in value.items() %}
                {% if subkey is defined and subvalue is defined %}
                    {% if subvalue is string and colorscheme.palette[subvalue] %}
                        {{ subkey }} = colors.{{ subvalue }},
                    {% elif subvalue is string %}
                        {{ subkey }} = "{{ subvalue }}",
                    {% else %}
                        {{ subkey }} = {{ subvalue }},
                    {% endif %}
                {% endif %}
            {% endfor %}
            }
        {% else %}
            {{ value }}
        {% endif %}, { transparent = transparent })
    {% else %}
        hl("{{ key }}",
        {% if value is mapping %}
            {
            {% for subkey, subvalue in value.items() %}
                {% if subkey is defined and subvalue is defined %}
                    {% if subvalue is string and colorscheme.palette[subvalue] %}
                        {{ subkey }} = colors.{{ subvalue }},
                    {% elif subvalue is string %}
                        {{ subkey }} = "{{ subvalue }}",
                    {% else %}
                        {{ subkey }} = {{ subvalue }},
                    {% endif %}
                {% endif %}
            {% endfor %}
            }
        {% else %}
            {{ value }}
        {% endif %})
    {% endif %}
{% endfor %}
end

return M
"""
