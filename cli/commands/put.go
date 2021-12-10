package commands

import (
	"github.com/spf13/cobra"
)

func PutCMD() *cobra.Command {
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
			ContentType: "",
			OpenEditor: false,
			Body: "",
			IsBodyStdin: false,
		},
	}

	cmd := &cobra.Command{
		Use:   "put <url> [flags]",
		Short: "Send a PUT request",
		Long:  `Send a PUT request to a given URL with a given body`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.URL = args[0]
			}

			return runWithBody(&opts, "PUT")
		},
	}

	cmd.Flags().StringVarP(&opts.Method.AuthType.TokenAuth, "token", "t", "", "The bearer token to use for authentication")
	cmd.Flags().StringVarP(&opts.Method.AuthType.BasicAuthUsername, "username", "u", "", "The username to use for basic authentication")
	cmd.Flags().StringVarP(&opts.Method.AuthType.BasicAuthPassword, "password", "p", "", "The password to use for basic authentication")
	cmd.Flags().BoolVarP(&opts.Method.JustShowBody, "just-body", "j", false, "Just show the response body")
	cmd.Flags().BoolVarP(&opts.Method.JustShowHeaders, "headers", "H", false, "Just show the response headers")
	cmd.Flags().StringVarP(&opts.Method.SaveFile, "save", "s", "", "Save the response to a file")
	cmd.Flags().StringVarP(&opts.Method.ContentType, "content-type", "c", "", "The content type of the body")
	cmd.Flags().StringVarP(&opts.Method.Body, "body", "b", "", "The body of the request")
	cmd.Flags().BoolVarP(&opts.Method.OpenEditor, "editor", "e", false, "Open the editor to edit the body")
	cmd.Flags().BoolVarP(&opts.Method.IsBodyStdin, "body-stdin", "i", false, "Read the body from stdin")

	return cmd
}
