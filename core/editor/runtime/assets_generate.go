// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	fs := http.Dir("./files")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "runtime",
		VariableName: "files",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
