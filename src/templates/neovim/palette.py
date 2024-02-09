template = """
return {
    {% for key, value in colorscheme.palette.items() %}
        {{ key }} = "{{ value }}",
    {% endfor %}
}
"""
