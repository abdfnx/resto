package checker

import (
	"fmt"
	"strings"

	"github.com/abdfnx/resto/tools"
	"github.com/abdfnx/resto/core/api"
	"github.com/abdfnx/resto/cmd/factory"

	"github.com/mgutz/ansi"
	"github.com/tidwall/gjson"
	"github.com/abdfnx/looker"
)

func Check(buildVersion string) {
	cmdFactory := factory.New()
	stderr := cmdFactory.IOStreams.ErrOut

	latestVersion := api.GetLatest()
	isFromHomebrewTap := isUnderHomebrew()
	isFromGo := isUnderGo()
	isFromUsrBinDir := isUnderUsr()
	isFromGHCLI := isUnderGHCLI()
	isFromAppData := isUnderAppData()

	var command = func() string {
		if isFromHomebrewTap {
			return "brew upgrade resto"
		} else if isFromGo {
			return "go get -u github.com/abdfnx/resto"
		} else if isFromUsrBinDir {
			return "curl -fsSL https://git.io/resto | bash"
		} else if isFromGHCLI {
			return "gh extention upgrade resto"
		} else if isFromAppData {
			return "iwr -useb https://git.io/resto-win | iex"
		}

		return ""
	}

	if buildVersion != latestVersion && gjson.Get(tools.SettingsContent(), "rs_settings.show_update").Bool() != false {
		fmt.Fprintf(stderr, "%s %s → %s\n",
		ansi.Color("There's a new version of ", "yellow") + ansi.Color("resto", "cyan") + ansi.Color(" is avalaible:", "yellow"),
		ansi.Color(buildVersion, "cyan"),
		ansi.Color(latestVersion, "cyan"))

		if command() != "" {
			fmt.Fprintf(stderr, ansi.Color("To upgrade, run: %s\n", "yellow"), ansi.Color(command(), "black:white"))
		}
	}
}

var restoExe, _ = looker.LookPath("resto")

func isUnderHomebrew() bool {
	return strings.Contains(restoExe, "brew")
}

func isUnderGo() bool {
	return strings.Contains(restoExe, "go")
}

func isUnderUsr() bool {
	return strings.Contains(restoExe, "usr")
}

func isUnderAppData() bool {
	return strings.Contains(restoExe, "AppData")
}

func isUnderGHCLI() bool {
	return strings.Contains(restoExe, "gh")
}
