package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gopamin/cli/internal/scaffolder"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		createNewProject()
	},
}

type projectType struct {
	Name string
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createNewProject() {
	projectName := promptProjectName()
	projectName = strings.Replace(strings.TrimSpace(projectName), " ", "-", -1)

	if isTaken := isProjectNameTaken(projectName); isTaken {
		fmt.Println("This project name is already taken in the current directory")
		return
	}

	projectType := promptSelectProjectType()

	scaffolder.New(projectName, projectType)
}

func isProjectNameTaken(name string) bool {
	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)
		if err != nil {
			log.Println("Could not read the directory")
		}
		if len(dirEntries) > 0 {
			return true
		}
	}

	return false
}

func promptProjectName() string {
	validate := func(input string) error {
		if len(strings.TrimSpace(input)) == 0 {
			return errors.New("Please enter a name")
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     "What's your project name? ",
		Templates: templates,
		Validate:  validate,
	}

	projectName, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return projectName
}

func promptSelectProjectType() string {
	projectTypes := []projectType{
		{Name: "Hello World"},
		{Name: "Microservice"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "▸ {{ .Name | green }}",
		Inactive: "  {{ .Name | magenta }}",
		Selected: "▸ {{ .Name | green }}",
	}

	prompt := promptui.Select{
		Label:     "Project type",
		Items:     projectTypes,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return projectTypes[index].Name
}
