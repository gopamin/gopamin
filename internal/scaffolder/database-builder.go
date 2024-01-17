package scaffolder

type mysqlBuilder struct {
	project *Project
}

func (m *mysqlBuilder) build() {
	fileGenerator([]string{"mysql-database"}, m.project)
	fileGenerator([]string{"mysql-docker-compose"}, m.project)
	goGetPackages(m.project.Path, []string{"github.com/go-sql-driver/mysql"})
}

type postgresBuilder struct {
	project *Project
}

func (p *postgresBuilder) build() {
	fileGenerator([]string{"postgres-database"}, p.project)
	fileGenerator([]string{"postgres-docker-compose"}, p.project)
	goGetPackages(p.project.Path, []string{"github.com/jackc/pgx/v5"})
}

type databaseBuilderFactory func(p *Project) boilerplateBuilder

var databaseBuilderMap = map[string]databaseBuilderFactory{
	"mysql": func(p *Project) boilerplateBuilder {
		return &mysqlBuilder{p}
	},
	"postgres": func(p *Project) boilerplateBuilder {
		return &postgresBuilder{p}
	},
}
