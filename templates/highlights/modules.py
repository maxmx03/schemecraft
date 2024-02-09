modules_template = """
local True = true
local False = false

local M = {}

M.load = function(opts)
 local hl = opts.hl

{% for key, value in colorscheme.highlights.modules.items() %}

    if opts.modules["{{ key }}"] then
        {% for subkey, subvalue in value.items() %}
            hl("{{ subkey }}",
                {
            {% for subkey2, subvalue2 in subvalue.items() %}
                {% if subvalue2 is string and colorscheme.palette[subvalue2] %}
                    {{ subkey2 }} = "{{ colorscheme.palette[subvalue2] }}",
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
