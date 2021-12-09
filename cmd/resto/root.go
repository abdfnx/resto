package resto

import (
	"github.com/abdfnx/resto/tools"
	"github.com/abdfnx/resto/core/layout"
	"github.com/abdfnx/resto/cmd/factory"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Execute start the CLI
func Execute(f *factory.Factory) *cobra.Command {
	tools.CheckDotResto()

	const desc = `send pretty HTTP & API requests from your terminal.`

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "resto <subcommand> [flags]",
		Short:  desc,
		Long: desc,
		SilenceErrors: true,
		Example: heredoc.Doc(`
			resto
			resto get https://api.github.com
			cat schema.graphql | resto post https://api.spacex.land/graphql --body-stdin
			resto delete https://api.secman.dev/api/logins/13 -t XXXX
		`),
		Annotations: map[string]string{
			"help:tellus": heredoc.Doc(`
				Open an issue at https://github.com/abdfnx/resto/issues
			`),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
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

	return rootCmd
}

func run() error {
	layout.Layout()

	return nil
}
