package checker

import (
	"fmt"
	"strings"

	"github.com/abdfnx/resto/cmd/factory"
	"github.com/abdfnx/resto/core/api"

	"github.com/mgutz/ansi"
	tcexe "github.com/Timothee-Cardoso/tc-exe"
)

func Check(buildVersion string, isCmd bool) {
	cmdFactory := factory.New()
	stderr := cmdFactory.IOStreams.ErrOut

	latestVersion := api.GetLatest()
	isFromHomebrewTap := isUnderHomebrew()
	isFromGo := isUnderGo()
	isFromUsrBinDir := isUnderUsr()
	isFromGHCLI := isUnderGHCLI()

	var command = func() string {
		if isFromHomebrewTap {
			return "brew upgrade resto"
		} else if isFromGo {
			return "go get -u github.com/abdfnx/resto"
		} else if isFromUsrBinDir {
			return "curl -fsSL https://git.io/resto | bash"
		} else if isFromGHCLI {
			return "gh extention upgrade resto"
		}

		return ""
	}

	if buildVersion != latestVersion {
		if isCmd {
			fmt.Fprintf(stderr, "\n%s %s → %s\n",
			ansi.Color("There's a new version of ", "yellow") + ansi.Color("resto", "cyan") + ansi.Color(" is avalaible:", "yellow"),
			ansi.Color(buildVersion, "cyan"),
			ansi.Color(latestVersion, "cyan"))

			if command() != "" {
				fmt.Fprintf(stderr, ansi.Color("To upgrade, run: %s\n", "yellow"), ansi.Color(command(), "black:white"))
			}
		}
	}
}

var restoExe, _ = tcexe.LookPath("resto")

func isUnderHomebrew() bool {
	return strings.Contains(restoExe, "brew")
}

func isUnderGo() bool {
	return strings.Contains(restoExe, "go")
}

func isUnderUsr() bool {
	return strings.Contains(restoExe, "usr")
}

func isUnderGHCLI() bool {
	return strings.Contains(restoExe, "gh")
}
