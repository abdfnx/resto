package cli

import (
	"github.com/spf13/cobra"
)

func HeadCMD() *cobra.Command {
	opts := Options{
		Method: &Method{
			AuthType: &Auth{
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

	cmd := &cobra.Command{
		Use:   "head <url> [flags]",
		Short: "Send a HEAD request",
		Long:  `Send a HEAD request to a URL.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.URL = args[0]
			}

			return runBasic(&opts, "HEAD")
		},
	}

	cmd.Flags().BoolVarP(&opts.Method.JustShowBody, "just-body", "j", false, "Just show the response body")
	cmd.Flags().BoolVarP(&opts.Method.JustShowHeaders, "headers", "H", false, "Just show the response headers")
	cmd.Flags().StringVarP(&opts.Method.SaveFile, "save", "s", "", "Save the response body to a file")
	cmd.Flags().StringVarP(&opts.Method.AuthType.BasicAuthUsername, "username", "u", "", "The username to use for basic authentication")
	cmd.Flags().StringVarP(&opts.Method.AuthType.BasicAuthPassword, "password", "p", "", "The password to use for basic authentication")
	cmd.Flags().StringVarP(&opts.Method.AuthType.TokenAuth, "token", "t", "", "The bearer token to use for authentication")

	return cmd
}
