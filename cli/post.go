package cli

import (
	"github.com/spf13/cobra"
)

func PostCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post <url> [flags]",
		Short: "Send a POST request",
		Long:  `Send a POST request to a given URL with a given body`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				withBodyOpts.URL = args[0]
			}

			return runWithBody(&withBodyOpts, "POST")
		},
	}

	cmd.Flags().StringVarP(&withBodyOpts.Method.AuthType.TokenAuth, "token", "t", "", "The bearer token to use for authentication")
	cmd.Flags().StringVarP(&withBodyOpts.Method.AuthType.BasicAuthUsername, "username", "u", "", "The username to use for basic authentication")
	cmd.Flags().StringVarP(&withBodyOpts.Method.AuthType.BasicAuthPassword, "password", "p", "", "The password to use for basic authentication")
	cmd.Flags().BoolVarP(&withBodyOpts.Method.JustShowBody, "just-body", "j", false, "Just show the response body")
	cmd.Flags().BoolVarP(&withBodyOpts.Method.JustShowHeaders, "headers", "H", false, "Just show the response headers")
	cmd.Flags().StringVarP(&withBodyOpts.Method.SaveFile, "save", "s", "", "Save the response to a file")
	cmd.Flags().StringVarP(&withBodyOpts.Method.ContentType, "content-type", "c", "", "The content type of the body")
	cmd.Flags().StringVarP(&withBodyOpts.Method.Body, "body", "b", "", "The body of the request")
	cmd.Flags().BoolVarP(&withBodyOpts.Method.OpenEditor, "editor", "e", false, "Open the editor to edit the body")
	cmd.Flags().BoolVarP(&withBodyOpts.Method.IsBodyStdin, "body-stdin", "i", false, "Read the body from stdin")

	return cmd
}
