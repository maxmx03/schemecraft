package system

import (
	"os"
	"text/template"
	"yeahboy-colorgen/scheme"
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

	var tmpl = template.Must(template.New(scheme.Name).Parse(schemeTemplate))
	err = tmpl.Execute(file, scheme)

	if err != nil {
		return err
	}

	return nil
}
