template = """
local True = true
local False = false

local M = {}

M.load = function(opts)
 local hl = opts.hl
 local colors = opts.colors

{% for key, value in colorscheme.highlights.syntax.items() %}
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
    {% endfor %}
end

return M
"""
