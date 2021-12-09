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

func CheckDotResto() {
	if _, err := os.Stat(requestFile); os.IsNotExist(err) {
		os.MkdirAll(dotResto, 0755)
		os.Create(requestFile)
	}
}

func RequestFile() string {
	return requestFile
}
