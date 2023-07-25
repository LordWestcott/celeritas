package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func doAuth() error {
	//migrations
	dbType := cel.DB.DatabaseType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := cel.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		gracefulExit(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;"), downFile)
	if err != nil {
		gracefulExit(err)
	}

	//run migrations
	err = doMigrate("up", "") //Not great if you have migrations you have yet to run.
	if err != nil {
		gracefulExit(err)
	}

	//copy files over
	err = copyFileFromTemplate("templates/data/user.go.txt", cel.RootPath+"/data/user.go")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", cel.RootPath+"/data/token.go")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/data/remember_token.go.txt", cel.RootPath+"/data/remember_token.go")
	if err != nil {
		gracefulExit(err)
	}

	//copy over middleware
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", cel.RootPath+"/middleware/auth.go")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", cel.RootPath+"/middleware/auth-token.go")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/middleware/remember.go.txt", cel.RootPath+"/middleware/remember.go")
	if err != nil {
		gracefulExit(err)
	}

	//copy over handlers
	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", cel.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		gracefulExit(err)
	}

	//copy over mail templates
	err = copyFileFromTemplate("templates/mailer/password-reset.html.tmpl", cel.RootPath+"/mail/password-reset.html.tmpl")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/mailer/password-reset.plain.tmpl", cel.RootPath+"/mail/password-reset.plain.tmpl")
	if err != nil {
		gracefulExit(err)
	}

	//copy over view templates
	err = copyFileFromTemplate("templates/views/login.jet", cel.RootPath+"/views/login.jet")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/views/forgot.jet", cel.RootPath+"/views/forgot.jet")
	if err != nil {
		gracefulExit(err)
	}

	err = copyFileFromTemplate("templates/views/reset-password.jet", cel.RootPath+"/views/reset-password.jet")
	if err != nil {
		gracefulExit(err)
	}

	color.Yellow(" - users, tokens, and remember_tokens created and executed")
	color.Yellow(" - user and token models created")
	color.Yellow(" - auth middleware created")
	color.Yellow("")
	color.Yellow(" - Don't forget to add user and token models in data/models.go and to add appropriate middleware to your routes!")

	return nil
}
