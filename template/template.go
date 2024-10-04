package template

import (
	"github.com/maxmx03/schemecraft/scheme"
	"github.com/maxmx03/schemecraft/system"
	"log"
	"path/filepath"
	"strings"
)

type ProjectStructure struct {
	colors struct {
		dir  string
		file string
	}
	readme string
}

func NewProject(colorscheme, filename string) ProjectStructure {
	var project ProjectStructure
	var root string = "build"
	root = filepath.Join(root, colorscheme)
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.file = filepath.Join(project.colors.dir, filename+".vim")
	project.readme = filepath.Join(root, "README.md")
	return project
}

func Create(scheme scheme.Scheme) {
	var colorscheme = strings.Split(scheme.Name, "-")[0]
	var vim8 = NewProject(colorscheme+"vim8", scheme.Name)
	createProjectDirs(vim8)
	createFile(vim8.colors.file, scheme, Vim8Colors())
	createFile(vim8.readme, scheme, Vim8Readme())

	var vim9 = NewProject(colorscheme+"vim9", scheme.Name)
	createProjectDirs(vim9)
	createFile(vim9.colors.file, scheme, Vim9Colors())
	createFile(vim9.readme, scheme, Vim9Readme())

	log.Printf("%v.vim created successfully", scheme.Name)
}

func createProjectDirs(project ProjectStructure) {
	var dirs []string
	dirs = append(dirs, project.colors.dir)
	for _, dir := range dirs {
		var err error = system.CreateDir(dir)
		if err != nil {
			panic(err)
		}
	}
}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)
	if err != nil {
		panic(err)
	}
}