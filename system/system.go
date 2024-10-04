package system

import (
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/maxmx03/schemecraft/scheme"
	"gopkg.in/yaml.v3"
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

	var tmpl = template.Must(template.New(scheme.Name).Funcs(sprig.FuncMap()).Parse(schemeTemplate))
	err = tmpl.Execute(file, scheme)

	if err != nil {
		return err
	}

	return nil
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

