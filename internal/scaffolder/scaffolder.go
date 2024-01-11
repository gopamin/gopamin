package scaffolder

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
)

type Project struct {
	ProjectName string
}

func New(projectName, projectType string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Unable to get your current working directory")
	}

	projectPath := filepath.Join(currentDir, projectName)

	err = os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	project := Project{ProjectName: projectName}

	fileGenerator(projectPath, "readme", project)
	fileGenerator(projectPath, "gitignore", project)
	fileGenerator(projectPath, "makefile", project)
	fileGenerator(projectPath, "env", project)
	fileGenerator(projectPath, "license", project)
	fileGenerator(projectPath, "load-env", project)
	fileGenerator(projectPath, "dockerfile", project)

	boilerplateSelector(projectPath, projectType, project)
	initGit(projectPath)
	initGoMod(projectName, projectPath)
	goGetPackages(projectPath, []string{"github.com/joho/godotenv"})
}

func boilerplateSelector(path, projectType string, p Project) {
	switch projectType {
	case "Hello World":
		createHelloWorldBoilerplate(path, p)
		return
	case "Microservice":
		createMicroserviceBoilerplate(path, p)
	}
}

func createHelloWorldBoilerplate(path string, p Project) {
	fileGenerator(path, "main", p)

	fmt.Printf("%v project created successfully", p.ProjectName)
}

func createMicroserviceBoilerplate(path string, p Project) {
	fmt.Println("Not Implemented Yet!")
}

func fileGenerator(path string, fileType string, p Project) {
	templateMapper := templateMapper()
	fileTemplate, fileName := templateMapper[fileType]()

	dir := filepath.Dir(filepath.Join(path, fileName))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory: %s\n", err)

			os.Exit(1)
		}
	}

	file, err := os.Create(filepath.Join(path, fileName))

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
