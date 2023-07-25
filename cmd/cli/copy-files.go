package main

import (
	"embed"
	"errors"
	"os"
)

// This awesome feature is called embed.
// It allows us to embed files into our binary.
//
//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath, targetFile string) error {
	if fileExists(targetFile) {
		return errors.New(targetFile + " already exists!")
	}

	//Get a file from the embedded filesystem.
	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		gracefulExit(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		gracefulExit(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
