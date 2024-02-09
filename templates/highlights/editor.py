editor_template = """
local True = true
local False = false

return {
{% for key, value in colorscheme.highlights.editor.items() %}
        {{ key }} =
        {% if value is mapping %}
            {
            {% for subkey, subvalue in value.items() %}
                {% if subkey is defined and subvalue is defined %}
                    {% if subvalue is string and colorscheme.palette[subvalue] %}
                        {{ subkey }} = "{{ colorscheme.palette[subvalue] }}",
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
        {% endif %},
    {% endfor %}
}
"""
