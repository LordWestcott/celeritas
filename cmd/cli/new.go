package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var appURL string

func doNew(appName string) {
	appName = strings.ToLower(appName)
	appURL = appName

	//sanitize appName (convert url to single word)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}

	color.Green("\tCreating new application: %s", appName)

	//git clone the skeleton application (has to be public) - Private is possible, but you have to jump through hoops.
	color.Green("\tCloning skeleton application...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/LordWestcott/celeritas-skeleton.git",
		Progress: os.Stdout,
		Depth:    1, //Only get the latest commit
	})
	if err != nil {
		gracefulExit(err)
	}

	//remove .git directory
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		gracefulExit(err)
	}

	//create a ready to go .env file
	color.Green("\tCreating .env file...")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		gracefulExit(err)
	}

	env := string(data)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${KEY}", cel.RandomString(32))

	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		gracefulExit(err)
	}

	//create a make file (appropriate for the OS)
	color.Green("\tCreating Makefile...")
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.windows", appName))
		if err != nil {
			gracefulExit(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			gracefulExit(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			gracefulExit(err)
		}
	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.mac", appName))
		if err != nil {
			gracefulExit(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			gracefulExit(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			gracefulExit(err)
		}
	}

	_ = os.Remove("./" + appName + "/Makefile.windows")
	_ = os.Remove("./" + appName + "/Makefile.mac")

	//update the go.mod file
	color.Green("\tUpdating go.mod file...")
	_ = os.Remove("./" + appName + "/go.mod")

	data, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		gracefulExit(err)
	}

	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)

	err = copyDataToFile([]byte(mod), fmt.Sprintf("./%s/go.mod", appName))
	if err != nil {
		gracefulExit(err)
	}

	//update existing .go files with correct name/imports
	color.Green("\tUpdating source files...")
	os.Chdir("./" + appName)
	updateSource()

	// run go mod tidy in the project
	color.Green("\tRunning go mod tidy...")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Start()
	if err != nil {
		gracefulExit(err)
	}

	color.Green("\tDone building " + appName + "!")
	color.Green("Go build something awesome!")
}
