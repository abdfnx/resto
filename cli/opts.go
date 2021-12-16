package cli

import (
	"github.com/abdfnx/resto/core/options"
)

var basicOpts = options.CLIOptions{
	Method: &options.Method{
		AuthType: &options.Auth{
			Type: "",
			TokenAuth: "",
			BasicAuthUsername: "",
			BasicAuthPassword: "",
		},
		JustShowBody: false,
		JustShowHeaders: false,
		SaveFile: "",
	},
}

var withBodyOpts = options.CLIOptions{
	Method: &options.Method{
		AuthType: &options.Auth{
			Type: "",
			TokenAuth: "",
			BasicAuthUsername: "",
			BasicAuthPassword: "",
		},
		JustShowBody: false,
		JustShowHeaders: false,
		SaveFile: "",
		ContentType: "",
		OpenEditor: false,
		Body: "",
		IsBodyStdin: false,
	},
}
