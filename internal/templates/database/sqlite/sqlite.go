package templates

import (
	_ "embed"
)

//go:embed files/internal/databases/sqlite.tmpl
var sqliteDatabase []byte

//go:embed files/env.tmpl
var sqliteEnv []byte

//go:embed files/makefile.tmpl
var sqliteMakefile []byte

//go:embed files/readme.tmpl
var sqliteReadme []byte

//go:embed files/cmd/main.tmpl
var sqliteMain []byte

func SqliteMakefileTemplate() ([]byte, string) {
	return sqliteMakefile, "Makefile"
}

func SqliteMainTemplate() ([]byte, string) {
	return sqliteMain, "cmd/main.go"
}

func SqliteDatabaseTemplate() ([]byte, string) {
	return sqliteDatabase, "internal/databases/sqlite.go"
}

func SqliteEnvTemplate() ([]byte, string) {
	return sqliteEnv, ".env"
}

func SqliteReadmeTemplate() ([]byte, string) {
	return sqliteReadme, "README.md"
}
