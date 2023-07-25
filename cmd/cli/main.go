package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/lordwestcott/celeritas"
)

const version = "1.0.0"

var cel celeritas.Celeritas

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		gracefulExit(err)
	}

	setup(arg1, arg2)

	switch arg1 {
	case "help":
		showHelp()

	case "new":
		if arg2 == "" {
			gracefulExit(errors.New("new command requires a project name"))
		}
		doNew(arg2)

	case "version":
		color.Yellow("Celeritas CLI Version: %s", version)
	case "make":
		if arg2 == "" {
			gracefulExit(errors.New("make command requires a subcommand: (migration|model|handler)"))
		}
		err = doMake(arg2, arg3)
		if err != nil {
			gracefulExit(err)
		}

	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		err := doMigrate(arg2, arg3)
		if err != nil {
			gracefulExit(err)
		}
		message = "Migrations Complete!"

	default:
		showHelp()
	}

	gracefulExit(nil, message)

}

func gracefulExit(err error, msg ...string) {
	message := ""
	if len(msg) > 0 { //this is a kinda hacky way to make msg an optional arg.
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}

	os.Exit(1)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	//first argument is the command

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) > 2 {
			arg2 = os.Args[2]

			if len(os.Args) > 3 {
				arg3 = os.Args[3]
			}
		}
	} else {
		color.Red("Error: Command Required")
		showHelp()
		return "", "", "", errors.New("Error: Command Required")
	}

	return arg1, arg2, arg3, nil
}
