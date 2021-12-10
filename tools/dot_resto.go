package tools

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

var homeDir, _ = homedir.Dir()

// check if ~/.resto/request.json exists
var dotResto = path.Join(homeDir, "./.resto")
var requestFile = path.Join(dotResto, "requestBody")
var cliDir = path.Join(dotResto, "/cli")

func CheckDotResto() {
	if _, err := os.Stat(requestFile); os.IsNotExist(err) {
		os.MkdirAll(dotResto, 0755)
		os.Create(requestFile)
	}

	if _, err := os.Stat(cliDir); os.IsNotExist(err) {
		os.MkdirAll(cliDir, 0755)
		os.Create(path.Join(cliDir, "requestBody.json"))
		os.Create(path.Join(cliDir, "requestBody.graphql"))
		os.Create(path.Join(cliDir, "requestBody.xml"))
		os.Create(path.Join(cliDir, "requestBody.html"))
		os.Create(path.Join(cliDir, "requestBody.txt"))
	}
}

func RequestFile() string {
	return requestFile
}

func CLIRequestFile(format string) string {
	return path.Join(cliDir, "requestBody." + format)
}
