package resto

import (
	"fmt"

	"github.com/abdfnx/resto/tools"
	"github.com/abdfnx/resto/core/layout"
	"github.com/abdfnx/resto/cmd/factory"
	"github.com/abdfnx/resto/cli"
	installCmd "github.com/abdfnx/resto/cli/install"
	runCmd "github.com/abdfnx/resto/cli/run"
	"github.com/abdfnx/resto/cli/settings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Execute start the CLI
func Execute(f *factory.Factory, version string, buildDate string) *cobra.Command {
	tools.CheckDotResto()

	const desc = `a CLI app can send pretty HTTP & API requests with TUI.`

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "resto <subcommand> [flags]",
		Short:  desc,
		Long: desc,
		SilenceErrors: true,
		Example: heredoc.Doc(`
			# Open Resto UI
			resto

			# Send a request to a URL
			resto get https://api.github.com

			# Send a request to a URL and use resto editor
			resto post https://api.xcode.codes --content-type json --editor

			# Read Body from stdin
			cat schema.graphql | resto post https://api.spacex.land/graphql --content-type graphql --body-stdin

			# Use Authentecation with Basic Auth or Bearer Token
			resto delete https://api.secman.dev/api/logins/13 --content-type json --token TOKEN

			# Save response to a file
			resto get http://localhost:3333/api/v1/hello --save response.json

			# Install binary app from script URL and run it.
			resto i https://get.docker.com

			# Send a request from Restofile
			# after creating a Restofile
			resto run

			# Get the latest release/tag of a repository (github, gitlab, bitbucket)
			resto get-latest microsoft/vscode

			# Update resto settings
			resto settings set theme dracula
		`),
		Annotations: map[string]string{
			"help:tellus": heredoc.Doc(`
				Open an issue at https://github.com/abdfnx/resto/issues
			`),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			layout.Layout(version)

			return nil
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Aliases: []string{"ver"},
		Short: "Print the version of your resto binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("resto version " + version + " " + buildDate)
		},
	}

	rootCmd.SetOut(f.IOStreams.Out)
	rootCmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	rootCmd.PersistentFlags().Bool("help", false, "Help for resto")
	rootCmd.SetHelpFunc(helpHelper)
	rootCmd.SetUsageFunc(rootUsageFunc)
	rootCmd.SetFlagErrorFunc(rootFlagErrorFunc)

	// Add sub-commands to root command
	rootCmd.AddCommand(
		cli.GetCMD(),
	    cli.PostCMD(),
		cli.PutCMD(),
		cli.PatchCMD(),
		cli.DeleteCMD(),
		cli.HeadCMD(),
		installCmd.InstallCMD(),
		runCmd.RunCMD(),
		cli.GetLatestCMD(),
		settings.SettingsCMD(),
		versionCmd,
	)

	return rootCmd
}
