package install_cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"runtime"
	"time"

	"github.com/abdfnx/resto/core/options"
	"github.com/abdfnx/resto/tools"

	"github.com/spf13/cobra"
	"github.com/MakeNowJust/heredoc"
	"github.com/briandowns/spinner"
)

func InstallCMD() *cobra.Command {
	opts := options.InstallCommandOptions{
		Shell: "",
		IsHidden: false,
		URL: "",
	}

	cmd := &cobra.Command{
		Use:   "install <url>",
		Short: "Install app from script URL",
		Long:  `Install binary app from script URL and run it.`,
		Aliases: []string{"i"},
		SilenceErrors: true,
		Example: heredoc.Doc(`
			# Install binary app from script URL
			resto i https://deno.land/x/install/install.sh

			# and now you can run it
			deno -V
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.URL = args[0]
			}

			return runInstall(&opts)
		},
	}

	p := "bash"

	if runtime.GOOS == "windows" {
		p = "powershell"
	}

	cmd.Flags().StringVarP(&opts.Shell, "shell", "s", "", "shell to use (Default: " + p + ")")
	cmd.Flags().BoolVarP(&opts.IsHidden, "hidden", "H", false, "hide the output")

	return cmd
}

func runInstall(opts *options.InstallCommandOptions) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🔗 Installing..."
	s.Start()

	term := ""

	if runtime.GOOS == "windows" {
		term = "powershell.exe"
	} else {
		term = "bash"
	}

	if opts.Shell == "" {
		opts.Shell = term
	}

	resp, err := http.Get(opts.URL)
	
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()

	body, berr := ioutil.ReadAll(resp.Body)

	if berr != nil {
		return berr
	}

	err, out, errout := tools.Exec(opts.Shell, string(body))

	if err != nil {
		fmt.Println(err)
		fmt.Println(errout)
	}

	s.Stop()

	if !opts.IsHidden {
		fmt.Println(out)
	}

	return nil
}
