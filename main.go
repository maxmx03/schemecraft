package main

import (
	"log"
	"os"
	"strings"

	"github.com/maxmx03/schemecraft/scheme"
	"github.com/maxmx03/schemecraft/system"
	vim "github.com/maxmx03/schemecraft/template"
)

const msg = `
    Usage:
    schemecraft colorscheme.json ... # to generate colorscheme
`

func main() {
	if len(os.Args) < 2 {
		log.Fatalf(msg)
	}

	for _, arg := range os.Args[1:] {
		var scheme scheme.Scheme = ReadFile(arg)
		vim.Create(scheme)
	}
}

func ReadFile(file string) scheme.Scheme {
	var scheme scheme.Scheme
	if strings.HasSuffix(file, "yml") || strings.HasSuffix(file, "yaml") {
		scheme = system.ReadYaml(file)
	}
	return scheme
}
