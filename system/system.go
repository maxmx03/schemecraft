package system

import (
	"github.com/Masterminds/sprig/v3"
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

	var tmpl = template.Must(template.New(scheme.Name).Funcs(sprig.FuncMap()).Parse(schemeTemplate))
	err = tmpl.Execute(file, scheme)

	if err != nil {
		return err
	}

	return nil
}
