package template

import "fmt"

func Readme() string {
	var readme = `# {{title .Name}}

![schemecraft](https://github.com/maxmx03/schemecraft/assets/50273941/ee682aae-00cb-4282-ba24-3d9621a430a3)

## Installation

To install {{title .Name}}, you need a plugin manager. \
In the example, bellow we are going to use vim-plug.

%v

## Plugins

Bellow are the {{title .Name}} supported plugins. \
Enable the plugins you want.

%v
`
	var installBlockCode = "```vim"
	installBlockCode += `
Plug '{{.Repo}}', { 'branch': 'vim' }

colorscheme {{.Name}}
`
	installBlockCode += "```"

	var plugins = "```vim"
	plugins += `
{{range $index, $plugins := .Highlights.Plugins -}}
{{range $pluginName, $pluginConfigs:= $plugins -}}
g:{{mustRegexFind "^[a-z]+" $.Name}}.plugins["{{$pluginName}}"] = false
{{end -}}
{{end -}}
`
	plugins += "```"

	return fmt.Sprintf(readme, installBlockCode, plugins)
}
