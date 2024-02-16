package template

func Colors() string {
	return `-- {{title .Name}} colorscheme
-- Repo: {{.Repo}}
-- Maintainer: {{title .Author}} <{{.Contact}}>
-- License: {{.License}}
require("{{.Name}}").load()`
}
