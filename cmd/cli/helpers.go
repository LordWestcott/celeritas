package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/lordwestcott/celeritas"
)

func setup(arg1, arg2 string) {
	if arg1 != "new" && arg1 != "help" && arg1 != "version" {
		err := godotenv.Load()
		if err != nil {
			gracefulExit(err)
		}

		path, err := os.Getwd()
		if err != nil {
			gracefulExit(err)
		}

		cel.DB = &celeritas.Database{}
		cel.RootPath = path
		cel.DB.DatabaseType = os.Getenv("DATABASE_TYPE")
	}
}

func getDSN() string {
	dbType := cel.DB.DatabaseType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		}
		return dsn
	}

	return "mysql://" + cel.BuildDSN()
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	//check for an error before doing anything else.
	if err != nil {
		return err
	}

	//check if current file is directory
	if fi.IsDir() {
		return nil
	}

	//only check .go files
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	//matching go file
	if matched {
		//read the file
		read, err := os.ReadFile(path)
		if err != nil {
			//something has really gone wrong here.
			gracefulExit(err)
		}

		newContents := strings.Replace(string(read), "myapp", appURL, -1) //-1 means replace all occurrences

		//write the changed file
		err = os.WriteFile(path, []byte(newContents), 0) //0 means default permissions
		if err != nil {
			//something has really gone wrong here.
			gracefulExit(err)
		}
	}

	return nil
}

func updateSource() {
	//walk the entire project folder, including subfolders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		gracefulExit(err)
	}
}

func showHelp() {
	color.Cyan(`Available Commands:

	help                  - Shows this help menu
	version               - Shows the current version of Celeritas CLI
	migrate               - Runs all up migrations that have not been run previously
	migrate down          - Reverses the most recent migration
	migrate reset         - Runs all down migrations in reverse order, and then all up migrations
	make migration <name> - Creates a new up and down migration in the migrations folder.
	make auth             - Creates and runs migrations for authentication tables, and creates models and middleware
	make handler <name>   - Creates a new stub handler in the handlers folder
	make model <name>     - Creates a new stub model in the data folder
	make session          - Creates a table in the database as a session store
	make mail <name>	  - Creates both a html and plaintext email template in the mail folder
	`)
}
