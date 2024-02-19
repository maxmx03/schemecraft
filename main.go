package main

import (
	"log"
	"os"
	"schemecraft/nvim"
	"schemecraft/scheme"
	"schemecraft/system"
	"schemecraft/vim"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf(`
    Usage:
    schemecraft create colorscheme.json ... # to generate colorscheme
    `)
	}

	var argument = os.Args[1]

	switch argument {
	case "create":
		var arguments []string = os.Args[2:]

		for index, arg := range arguments {
			var scheme scheme.Scheme = ReadFile(arg)
			var isMainTheme bool = index == 0
			nvim.Create(scheme, isMainTheme)
			vim.Create(scheme)
		}
	case "update":
		var argument = os.Args[2]
		var arguments []string = os.Args[3:]

		for index, arg := range arguments {
			var scheme scheme.Scheme = ReadFile(arg)
			var isMainTheme bool = index == 0

			if argument == "nvim" {
				nvim.Update(scheme, isMainTheme)
			} else if argument == "vim" {
				vim.Update(scheme)
			} else {
				log.Fatalf("Invalid Argument %v\nUsage: schemecraft update nvim colorscheme.json", argument)
				break
			}
		}
	default:
		log.Fatalf("Invalid Argument %v\nUsage: schemecraft create colorscheme.json", argument)
	}
}

func isYaml(value string) bool {
	return strings.HasSuffix(value, "yml") || strings.HasSuffix(value, "yaml")
}

func ReadFile(file string) scheme.Scheme {
	var scheme scheme.Scheme

	if isYaml(file) {
		scheme = system.ReadYaml(file)
	} else {
		scheme = system.ReadJson(file)
	}

	return scheme
}
