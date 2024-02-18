package main

import (
	"log"
	"os"
	"yeahboy/nvim"
	"yeahboy/scheme"
	"yeahboy/vim"

	"gopkg.in/yaml.v3"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf(`
    Usage:
    yeahboy create colorscheme... # to generate colorscheme
    `)
	}
	var argument = os.Args[1]

	switch argument {
	case "create":
		for index, arg := range os.Args[2:] {
			var scheme scheme.Scheme = ReadYaml(arg + ".yml")
			var isMainTheme bool = index == 0
			nvim.Create(scheme, isMainTheme)
			vim.Create(scheme)
		}
	default:
		log.Fatalf("Invalid Argument %v\nUsage: yeahboy create colorscheme", argument)
	}
}

func ReadYaml(filename string) scheme.Scheme {
	var scheme scheme.Scheme
	var yamlFile, err = os.ReadFile(filename)

	LogFatal("Couldn't read yaml file: ", err)
	err = yaml.Unmarshal([]byte(yamlFile), &scheme)
	LogFatal("Couldn't unmarshal", err)

	return scheme
}

func LogFatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%v %v", msg, err)
	}
}
