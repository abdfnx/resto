package cli

import (
	"github.com/spf13/cobra"
)

func GetCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <url> [flags]",
		Short: "Send a GET request",
		Long:  `Send a GET request to a URL.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				basicOpts.URL = args[0]
			}

			return runBasic(&basicOpts, "GET")
		},
	}

	cmd.Flags().BoolVarP(&basicOpts.Method.JustShowBody, "just-body", "j", false, "Just show the response body")
	cmd.Flags().BoolVarP(&basicOpts.Method.JustShowHeaders, "headers", "H", false, "Just show the response headers")
	cmd.Flags().StringVarP(&basicOpts.Method.SaveFile, "save", "s", "", "Save the response body to a file")
	cmd.Flags().StringVarP(&basicOpts.Method.AuthType.BasicAuthUsername, "username", "u", "", "The username to use for basic authentication")
	cmd.Flags().StringVarP(&basicOpts.Method.AuthType.BasicAuthPassword, "password", "p", "", "The password to use for basic authentication")
	cmd.Flags().StringVarP(&basicOpts.Method.AuthType.TokenAuth, "token", "t", "", "The bearer token to use for authentication")

	return cmd
}
