template = """vim9script
var False = false
var True = true
g:{{ colorscheme.name }} = {
    {% for key, value in colorscheme.default_config.items() %}
        {% if key == 'modules' %}
            {{ key }}: {{ value }},
        {% endif %}
    {% endfor %}
}
"""
