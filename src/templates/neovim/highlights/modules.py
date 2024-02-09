template = """
local True = true
local False = false

local M = {}

M.load = function(opts)
 local hl = opts.hl
 local colors = opts.colors
 local modules = opts.config.modules

{% for key, value in colorscheme.highlights.modules.items() %}

    if modules["{{ key }}"] then
        {% for subkey, subvalue in value.items() %}
            hl("{{ subkey }}",
                {
            {% for subkey2, subvalue2 in subvalue.items() %}
                {% if subvalue2 is string and colorscheme.palette[subvalue2] %}
                    {{ subkey2 }} = colors.{{ subvalue2 }},
                {% elif subvalue is string %}
                    {{ subkey2 }} = "{{ subvalue2 }}",
                {% else %}
                    {{ subkey2 }} = {{ subvalue2 }},
                {% endif %}
            {% endfor %}})
        {% endfor %}
    end

{% endfor %}
end

return M
"""
