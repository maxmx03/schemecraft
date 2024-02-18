package vim

import (
	"log"
	"path/filepath"
	"strings"
	"yeahboy/scheme"
	"yeahboy/system"
	"yeahboy/vim/template"
)

type projectStructure struct {
	colors struct {
		dir  string
		file string
	}
}

func Create(scheme scheme.Scheme) {
	var project projectStructure

	var root = "build"
	var schemeName = strings.Split(scheme.Name, "_")[0]
	root = filepath.Join(root, schemeName+".vim")
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.file = filepath.Join(project.colors.dir, scheme.Name+".vim")

	var dirs []string
	dirs = append(dirs, project.colors.dir)

	for _, dir := range dirs {
		var err error = system.CreateDir(dir)

		if err != nil {
			panic(err)
		}
	}

	createFile(project.colors.file, scheme, template.Colors())
	log.Printf("%v.vim created successfully", scheme.Name)
}

func Update() {}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)

	if err != nil {
		panic(err)
	}
}
