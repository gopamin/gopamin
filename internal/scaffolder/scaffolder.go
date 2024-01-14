package scaffolder

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
)

type Project struct {
	Database    string
	Platform    string
	IsClean     bool
	Name        string
	Path        string
	ProjectType string
}

// func (p *Project) Build() {
// 	fmt.Printf("%+v", p)
// }

func New(projectType, platform, name, database string, isClean bool) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Unable to get your current working directory")
	}

	projectPath := filepath.Join(currentDir, name)

	// err = os.Mkdir(name, 0755)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	p := Project{
		Database:    database,
		Platform:    platform,
		IsClean:     isClean,
		Name:        name,
		Path:        projectPath,
		ProjectType: projectType,
	}

	// fmt.Printf("%+v", p)

	// microserviceBuilderFactory := microserviceBuilderMap[p.Platform]
	// microserviceBuilder := microserviceBuilderFactory(p)
	// microserviceDirector := &microserviceDirector{builder: microserviceBuilder}
	// microserviceDirector.construct()

	// apiBuilderFactory := apiBuilderMap[p.Platform]
	// apiBuilder := apiBuilderFactory(p)
	// apiDirector := &apiDirector{builder: apiBuilder}
	// apiDirector.construct()

	// helloWorldBuilderFactory := helloWorldBuilderMap[p.ProjectType]
	// helloWorldBuilder := helloWorldBuilderFactory(p)
	// helloWorldDirector := &helloWorldDirector{builder: helloWorldBuilder}
	// helloWorldDirector.construct()

	// builder := concreteBuilder{}
	// builder.build(p)

	tmpBuilderFactory := tmpBuilderMap[p.ProjectType]
	webAppBuilder := tmpBuilderFactory(p)
	tmpDirector := &tmpDirector{builder: webAppBuilder}
	tmpDirector.construct()

	// webAppBuilderFactory := webAppBuilderMap[p.ProjectType]
	// webAppBuilder := webAppBuilderFactory(p)
	// webAppDirector := &webAppDirector{builder: webAppBuilder}
	// webAppDirector.construct()

	// fileGenerator("readme", p)
	// fileGenerator("gitignore", p)
	// fileGenerator("dockerignore", p)
	// fileGenerator("makefile", p)
	// fileGenerator("env", p)
	// fileGenerator("license", p)
	// fileGenerator("load-env", p)
	// fileGenerator("dockerfile", p)
	// boilerplateSelector(p)
	// initGit(projectPath)
	// initGoMod(name, projectPath)
	// goGetPackages(projectPath, []string{"github.com/joho/godotenv"})
}

func boilerplateSelector(p Project) {
	switch p.ProjectType {
	case "hello-world":
		createHelloWorldBoilerplate(p)
		return
	default:
		fmt.Println("Not implemented yet")
		return
	}
}

func createHelloWorldBoilerplate(p Project) {
	fileGenerator("main", p)

	fmt.Printf("%v project created successfully", p.Name)
}

func fileGenerator(fileType string, p Project) {
	templateMapper := templateMapper()
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
