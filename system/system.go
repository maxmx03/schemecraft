package system

import (
	"encoding/json"
	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"text/template"
	"yeahboy/scheme"
)

func CreateDir(directory string) error {
	return os.MkdirAll(directory, os.ModePerm)
}

func WriteTemplateFile(filePath string, scheme scheme.Scheme, schemeTemplate string) error {
	var file, err = os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	var tmpl = template.Must(template.New(scheme.GetName()).Funcs(sprig.FuncMap()).Parse(schemeTemplate))
	err = tmpl.Execute(file, scheme)

	if err != nil {
		return err
	}

	return nil
}

func ReadYaml(filename string) scheme.SchemeYaml {
	var scheme scheme.SchemeYaml
	var yamlFile, err = os.ReadFile(filename)

	LogFatal("Couldn't read yaml file: ", err)
	err = yaml.Unmarshal([]byte(yamlFile), &scheme)
	LogFatal("Couldn't unmarshal", err)

	return scheme
}

func ReadJson(filename string) scheme.SchemeJson {
	var scheme scheme.SchemeJson
	var jsonFile, err = os.ReadFile(filename)

	LogFatal("Couldn't read yaml file: ", err)
	err = json.Unmarshal([]byte(jsonFile), &scheme)
	LogFatal("Couldn't unmarshal", err)

	return scheme
}

func LogFatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%v %v", msg, err)
	}
}
