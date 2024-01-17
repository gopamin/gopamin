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

type mongodbBuilder struct {
	project *Project
}

func (m *mongodbBuilder) build() {
	fileGenerator([]string{"mongodb-database"}, m.project)
	fileGenerator([]string{"mongodb-docker-compose"}, m.project)
	goGetPackages(m.project.Path, []string{"go.mongodb.org/mongo-driver/mongo"})
}

type redisBuilder struct {
	project *Project
}

func (m *redisBuilder) build() {
	fileGenerator([]string{"redis-database"}, m.project)
	fileGenerator([]string{"redis-docker-compose"}, m.project)
	goGetPackages(m.project.Path, []string{"github.com/redis/go-redis/v9"})
}

type databaseBuilderFactory func(p *Project) boilerplateBuilder

var databaseBuilderMap = map[string]databaseBuilderFactory{
	"mysql": func(p *Project) boilerplateBuilder {
		return &mysqlBuilder{p}
	},
	"postgres": func(p *Project) boilerplateBuilder {
		return &postgresBuilder{p}
	},
	"mongodb": func(p *Project) boilerplateBuilder {
		return &mongodbBuilder{p}
	},
	"redis": func(p *Project) boilerplateBuilder {
		return &redisBuilder{p}
	},
}
