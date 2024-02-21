package template

func Colors() string {
	return `-- {{title (mustRegexFind "^[a-z]+" .Name)}} colorscheme
-- Repo: {{.Repo}}
-- Maintainer: {{title .Author}} <{{.Contact}}>
-- License: {{.License}}
require("{{mustRegexFind "^[a-z]+" .Name}}").load("{{.Name}}")`
}
