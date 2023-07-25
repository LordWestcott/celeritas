package main

import (
	"fmt"
	"time"
)

func doSessionTable() error {
	dbType := cel.DB.DatabaseType
	if dbType == "mariadb" {
		dbType = "mysql"
	}

	if dbType == "postgresql" {
		dbType = "postgres"
	}

	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())

	upFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		gracefulExit(err)
	}

	err = copyDataToFile([]byte("drop table if exists sessions;"), downFile)
	if err != nil {
		gracefulExit(err)
	}

	err = doMigrate("up", "")
	if err != nil {
		gracefulExit(err)
	}
	return nil
}
