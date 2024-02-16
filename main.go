package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"yeahboy-colorgen/nvim"
	"yeahboy-colorgen/scheme"
)

func main() {
	var scheme scheme.Scheme = ReadYaml()
	var argument = os.Args[1]

	switch argument {
	case "create":
		if err := nvim.Create(scheme); err != nil {
			log.Fatal(err)
		}
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
