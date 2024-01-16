package scaffolder

import "fmt"

type mysqlBuilder struct {
	project *Project
}

func (m *mysqlBuilder) build() {
	fileGenerator([]string{"mysql-database"}, m.project)
	fileGenerator([]string{"mysql-docker-compose"}, m.project)
	goGetPackages(m.project.Path, []string{"github.com/go-sql-driver/mysql"})
}

type mongodbBuilder struct {
	project *Project
}

func (m *mongodbBuilder) build() {
	fmt.Printf("%v\n", m.project)
}

type databaseBuilderFactory func(p *Project) boilerplateBuilder

var databaseBuilderMap = map[string]databaseBuilderFactory{
	"mysql": func(p *Project) boilerplateBuilder {
		return &mysqlBuilder{p}
	},
	"mongodb": func(p *Project) boilerplateBuilder {
		return &mongodbBuilder{p}
	},
}
