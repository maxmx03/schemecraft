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

var project projectStructure

func setProject(scheme scheme.Scheme, root string) {
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.file = filepath.Join(project.colors.dir, scheme.Name+".vim")
}

func createProjectDirs() {
	var dirs []string
	dirs = append(dirs, project.colors.dir)

	for _, dir := range dirs {
		var err error = system.CreateDir(dir)

		if err != nil {
			panic(err)
		}
	}
}

func createProjectFiles(scheme scheme.Scheme) {
	createFile(project.colors.file, scheme, template.Colors())
}

func Create(scheme scheme.Scheme) {
	var schemeName = strings.Split(scheme.Name, "_")[0]
	var root string = "build"
	root = filepath.Join(root, schemeName+".vim")
	setProject(scheme, root)
	createProjectDirs()
	createProjectFiles(scheme)
	log.Printf("%v.vim created successfully", scheme.Name)
}

func Update(scheme scheme.Scheme) {
	var root string
	setProject(scheme, root)
	createFile(project.colors.file, scheme, template.Colors())
	createProjectFiles(scheme)
	log.Printf("%v.vim updated successfully", scheme.Name)
}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)

	if err != nil {
		panic(err)
	}
}
