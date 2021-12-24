package cli

import (
	"github.com/spf13/cobra"
)

func GetLatestCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-latest <repo || id> [flags]",
		Short: "Get Latest repository tag",
		Long:  `Get The Latest Tag of a repository from a registry (github, gitlab, bitbucket).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				getLatestOpts.Repo = args[0]
			}

			return runGetLatest(&getLatestOpts)
		},
	}

	cmd.Flags().StringVarP(&getLatestOpts.Registry, "registry", "r", "", "The registry to use")
	cmd.Flags().StringVarP(&getLatestOpts.Token, "token", "t", "", "The access token to use it the registry requires authentication")

	return cmd
}
