colors_template="""
-- {{ colorscheme.name }} colorscheme
-- Repo: {{ colorscheme.repo }}
-- Maintainer: {{ colorscheme.author }} <"{{ colorscheme.contact }}">
-- License: {{ colorscheme.license }}

require('{{ colorscheme.name }}').load()
"""
