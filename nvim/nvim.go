package nvim

import (
	"log"
	"path/filepath"
	"schemecraft/nvim/template"
	"schemecraft/scheme"
	"schemecraft/system"
	"strings"
)

type projectStructure struct {
	githubWorkFlows struct {
		dir   string
		files [2]string
	}
	colors struct {
		dir  string
		file string
	}
	lua struct {
		dir     string
		file    string
		config  string
		palette struct {
			dir  string
			file string
		}
		highlights struct {
			dir  string
			file string
		}
		utils string
		color string
	}
	tests struct {
		dir       string
		loadSpec  string
		setupSpec string
	}
	dockerfile string
	compose    string
	shell      string
	readme     string
}

var project projectStructure
var isMainTheme = false

func setProject(scheme scheme.Scheme, root string) {
	var schemeName = strings.Split(scheme.GetName(), "-")[0]
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.file = filepath.Join(project.colors.dir, scheme.GetName()+".lua")
	if isMainTheme {
		project.githubWorkFlows.dir = filepath.Join(root, ".github", "workflows")
		project.githubWorkFlows.files = [2]string{}
		project.githubWorkFlows.files[0] = filepath.Join(project.githubWorkFlows.dir, "release-please.yml")
		project.githubWorkFlows.files[1] = filepath.Join(project.githubWorkFlows.dir, "test.yml")
		project.lua.dir = filepath.Join(root, "lua", schemeName)
		project.lua.palette.dir = filepath.Join(project.lua.dir, "palette")
		project.lua.highlights.dir = filepath.Join(project.lua.dir, "highlights")
		project.lua.palette.file = filepath.Join(project.lua.palette.dir, "init.lua")
		project.lua.highlights.file = filepath.Join(project.lua.highlights.dir, "init.lua")
		project.lua.file = filepath.Join(project.lua.dir, "init.lua")
		project.lua.utils = filepath.Join(project.lua.dir, "utils.lua")
		project.lua.config = filepath.Join(project.lua.dir, "config.lua")
		project.lua.color = filepath.Join(project.lua.dir, "color.lua")
		project.tests.dir = filepath.Join(root, "tests")
		project.tests.loadSpec = filepath.Join(project.tests.dir, "load_spec.lua")
		project.tests.setupSpec = filepath.Join(project.tests.dir, "setup_spec.lua")
		project.dockerfile = filepath.Join(root, "Dockerfile")
		project.compose = filepath.Join(root, "docker-compose.yml")
		project.shell = filepath.Join(root, "shell.nix")
		project.readme = filepath.Join(root, "README.md")
	} else {
		project.lua.palette.file = filepath.Join(project.lua.palette.dir, scheme.GetName()+".lua")
		project.lua.highlights.file = filepath.Join(project.lua.highlights.dir, scheme.GetName()+".lua")
	}

}

func createProjectDirs() {
	if !isMainTheme {
		return
	}

	var dirs []string
	dirs = append(dirs, project.githubWorkFlows.dir)
	dirs = append(dirs, project.colors.dir)
	dirs = append(dirs, project.lua.dir)
	dirs = append(dirs, project.lua.palette.dir)
	dirs = append(dirs, project.lua.highlights.dir)
	dirs = append(dirs, project.tests.dir)

	for _, dir := range dirs {
		var err error = system.CreateDir(dir)

		if err != nil {
			panic(err)
		}
	}
}

func createProjectFiles(scheme scheme.Scheme) {
	createFile(project.lua.palette.file, scheme, template.Palette())
	createFile(project.lua.highlights.file, scheme, template.Highlights())
	createFile(project.colors.file, scheme, template.Colors())

	if isMainTheme {
		createFile(project.readme, scheme, template.Readme())
		createFile(project.compose, scheme, template.DockerCompose())
		createFile(project.dockerfile, scheme, template.DockerFile())
		createFile(project.shell, scheme, template.NixShell())
		createFile(project.tests.setupSpec, scheme, template.SetupSpec())
		createFile(project.tests.loadSpec, scheme, template.LoadSpec())
		createFile(project.lua.utils, scheme, template.Utils())
		createFile(project.lua.config, scheme, template.Config())
		createFile(project.lua.color, scheme, template.Color())
		createFile(project.lua.file, scheme, template.Root())
		var release, test = template.GithubWorkFlows()
		createFile(project.githubWorkFlows.files[0], scheme, release)
		createFile(project.githubWorkFlows.files[1], scheme, test)
	}
}

func Create(scheme scheme.Scheme, mainTheme bool) {
	isMainTheme = mainTheme
	var schemeName = strings.Split(scheme.GetName(), "-")[0]
	var root string = "build"
	root = filepath.Join(root, schemeName+".nvim")
	setProject(scheme, root)
	createProjectDirs()
	createProjectFiles(scheme)
	log.Printf("%v.lua created successfully", scheme.GetName())
}

func Update(scheme scheme.Scheme, mainTheme bool) {
	isMainTheme = mainTheme
	var root string
	setProject(scheme, root)
	createProjectFiles(scheme)
	log.Printf("%v.lua updated successfully", scheme.GetName())
}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)

	if err != nil {
		panic(err)
	}
}
