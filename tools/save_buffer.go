package tools

import (
	"io/ioutil"

	"github.com/abdfnx/resto/core/editor"
)

func SaveBuffer(b *editor.Buffer, path string) error {
	return ioutil.WriteFile(path, []byte(b.String()), 0600)
}
