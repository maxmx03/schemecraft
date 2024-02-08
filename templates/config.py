config_template = """
return {
    {%- for key, value in colorscheme.default_config.items() -%}
        {{ key }} =
        {%- if value is mapping -%}
            {
            {%- for subkey, subvalue in value.items() -%}
                {%- if subkey is defined and subvalue is defined -%}
                    {{ subkey }} = '{{ subvalue }}',
                {%- endif -%}
            {%- endfor -%}
            }
        {%- else -%}
            '{{ value }}'
        {%- endif -%},
    {%- endfor -%}
}
"""
