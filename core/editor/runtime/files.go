//go:generate go run assets_generate.go

package runtime

import "github.com/abdfnx/resto/core/editor"

var Files = editor.NewRuntimeFiles(files)
