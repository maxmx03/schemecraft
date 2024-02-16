package nvim

import (
	"log"
	"path/filepath"
	"yeahboy/nvim/template"
	"yeahboy/scheme"
	"yeahboy/system"
)

type projectStructure struct {
	colors struct {
		dir         string
		colorscheme string
	}
	lua struct {
		dir     string
		init    string
		config  string
		palette string
		utils   string
		color   string
	}
	tests struct {
		dir       string
		loadSpec  string
		setupSpec string
	}
	docs struct {
		dir            string
		colorschemeTxt string
	}
	dockerfile string
	compose    string
	shell      string
}

func Create(scheme scheme.Scheme) {
	var project projectStructure

	var root = "build"
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.colorscheme = filepath.Join(project.colors.dir, scheme.Name+".lua")
	project.lua.dir = filepath.Join(root, "lua", scheme.Name)
	project.lua.init = filepath.Join(project.lua.dir, "init.lua")
	project.lua.utils = filepath.Join(project.lua.dir, "utils.lua")
	project.lua.config = filepath.Join(project.lua.dir, "config.lua")
	project.lua.palette = filepath.Join(project.lua.dir, "palette.lua")
	project.lua.color = filepath.Join(project.lua.dir, "color.lua")
	project.tests.dir = filepath.Join(root, "tests")
	project.tests.loadSpec = filepath.Join(project.tests.dir, "load_spec.lua")
	project.tests.setupSpec = filepath.Join(project.tests.dir, "setup_spec.lua")
	project.docs.dir = filepath.Join(root, "docs")
	project.docs.colorschemeTxt = filepath.Join(project.docs.dir, scheme.Name+".txt")
	project.dockerfile = filepath.Join(root, "Dockerfile")
	project.compose = filepath.Join(root, "docker-compose.yml")
	project.shell = filepath.Join(root, "shell.nix")

	var dirs []string
	dirs = append(dirs, project.colors.dir)
	dirs = append(dirs, project.lua.dir)
	dirs = append(dirs, project.tests.dir)
	dirs = append(dirs, project.docs.dir)

	for _, dir := range dirs {
		var err error = system.CreateDir(dir)

		if err != nil {
			panic(err)
		}
	}

	createFile(project.compose, scheme, template.DockerCompose())
	createFile(project.dockerfile, scheme, template.DockerFile())
	createFile(project.shell, scheme, template.NixShell())
	createFile(project.docs.colorschemeTxt, scheme, template.Docs())
	createFile(project.tests.setupSpec, scheme, template.SetupSpec())
	createFile(project.tests.loadSpec, scheme, template.LoadSpec())
	createFile(project.colors.colorscheme, scheme, template.Colors())
	createFile(project.lua.utils, scheme, template.Utils())
	createFile(project.lua.config, scheme, template.Config())
	createFile(project.lua.palette, scheme, template.Palette())
	createFile(project.lua.color, scheme, template.Color())
	createFile(project.lua.init, scheme, template.Root())
	log.Printf("%v.lua created successfully", scheme.Name)
}

func Update() {}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)

	if err != nil {
		panic(err)
	}
}
