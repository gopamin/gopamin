package scaffolder

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gopamin/cli/internal/templates"
)

type Project struct {
	Database    string
	Platform    string
	IsClean     bool
	Name        string
	Path        string
	ProjectType string
}

func New(projectType, platform, name, database string, isClean bool) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Unable to get your current working directory")
	}

	projectPath := filepath.Join(currentDir, name)

	err = os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := Project{
		Database:    database,
		Platform:    platform,
		IsClean:     isClean,
		Name:        name,
		Path:        projectPath,
		ProjectType: projectType,
	}

	generateProjectAgnosticFiles(&p)

	builderFactory := buildersMap[p.ProjectType]
	builder := builderFactory(&p)
	director := &director{builder: builder}
	director.construct()
}

func generateProjectAgnosticFiles(p *Project) {
	fileGenerator("gitignore", p)
	fileGenerator("dockerignore", p)
	fileGenerator("license", p)
	fileGenerator("dockerfile", p)

	initGit(*&p.Path)
	initGoMod(*&p.Name, *&p.Path)
	goGetPackages(*&p.Path, []string{"github.com/joho/godotenv"})
}

func fileGenerator(fileType string, p *Project) {
	templateMapper := templates.Mapper()
	fileTemplate, fileName := templateMapper[fileType]()

	dir := filepath.Dir(filepath.Join(p.Path, fileName))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory: %s\n", err)

			os.Exit(1)
		}
	}

	file, err := os.Create(filepath.Join(p.Path, fileName))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	template := template.Must(
		template.New(fileName).Parse(
			string(fileTemplate),
		),
	)

	err = template.Execute(file, p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	color.Green(fileName + " created")
}

func executeCommand(name string, args []string, dir string) error {
	var out bytes.Buffer

	command := exec.Command(name, args...)
	command.Dir = dir
	command.Stdout = &out

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func initGoMod(projectName string, appDir string) error {
	if err := executeCommand("go", []string{"mod", "init", projectName}, appDir); err != nil {
		return err
	}

	return nil
}

func initGit(appDir string) {
	err := executeCommand("git", []string{"init"}, appDir)
	if err != nil {
		fmt.Printf("Error initializing git repo: %v", err)
		os.Exit(1)
	}
}

func goGetPackages(appDir string, packages []string) error {
	for _, packageName := range packages {
		if err := executeCommand("go",
			[]string{"get", "-u", packageName},
			appDir); err != nil {
			return err
		}
	}

	return nil
}
