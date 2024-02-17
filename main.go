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
	var scheme scheme.Scheme = ReadYaml()

	if len(os.Args) < 2 {
		log.Fatalf(`
    Usage:
    yeahboy create # to generate colorscheme
    yeahboy update # to update colorscheme
    `)
	}

	var argument = os.Args[1]

	switch argument {
	case "create":
		nvim.Create(scheme)
		vim.Create(scheme)
	default:
		log.Printf("Invalid Argument %v\n", argument)
		log.Println("Usage: yeahboy create or yeahboy update")
	}
}

func ReadYaml() scheme.Scheme {
	var scheme scheme.Scheme
	var yamlFile, err = os.ReadFile("colorscheme.yml")

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
