package settings

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abdfnx/resto/tools"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/tidwall/sjson"
)

func SettingsSet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set new or update settings",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				var value bool

				if string(args[1]) == "true" {
					value = true
				} else {
					value = false
				}

				if strings.Contains(args[0], "theme") || strings.Contains(args[0], "colorscheme") || strings.Contains(args[0], "request_body") {
					theme, err := sjson.Set(tools.SettingsContent(), "rs_settings.request_body.theme", args[1])

					if err != nil {
						panic(err)
					}

					serr := ioutil.WriteFile(tools.SettingsFile(), []byte(string(theme)), 0644)

					if serr != nil {
						panic(serr)
					}
				} else if strings.Contains(args[0], "show_update") {
					if string(args[1]) == "true" || string(args[1]) == "false" {
						update, err := sjson.Set(tools.SettingsContent(), "rs_settings.show_update", value)

						if err != nil {
							panic(err)
						}

						uerr := ioutil.WriteFile(tools.SettingsFile(), []byte(string(update)), 0644)

						if uerr != nil {
							panic(uerr)
						}
					} else {
						fmt.Println(ansi.Color("rs_settings.show_update must be `true` or `false`", "red"))
						os.Exit(1)
					}
				} else if strings.Contains(args[0], "mouse") {
					if string(args[1]) == "true" || string(args[1]) == "false" {
						mouse, err := sjson.Set(tools.SettingsContent(), "rs_settings.enable_mouse", value)

						if err != nil {
							panic(err)
						}

						merr := ioutil.WriteFile(tools.SettingsFile(), []byte(string(mouse)), 0644)

						if merr != nil {
							panic(merr)
						}
					} else {
						fmt.Println(ansi.Color("rs_settings.enable_mouse must be `true` or `false`", "red"))
						os.Exit(1)
					}
				}

				fmt.Println(ansi.Color("Settings updated", "green"))
			}
		},
	}

	return cmd
}
