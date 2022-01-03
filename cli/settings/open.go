package settings

import (
	"log"
	"io/ioutil"

	"github.com/abdfnx/resto/tools"
	"github.com/abdfnx/resto/core/editor"
	"github.com/abdfnx/resto/core/editor/runtime"

	"github.com/spf13/cobra"
	"github.com/rivo/tview"
	"github.com/tidwall/gjson"
	"github.com/gdamore/tcell/v2"
)

func SettingsOpen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Aliases: []string{"o", "."},
		Short: "Set new or update settings",
		Run: func(cmd *cobra.Command, args []string) {
			app := tview.NewApplication()
			settingsContent, err := ioutil.ReadFile(tools.SettingsFile())
			bufferSettings := editor.NewBufferFromString(string(settingsContent), tools.SettingsFile())
			if err != nil {
				log.Fatalf("could not read %v: %v", tools.SettingsFile(), err)
			}

			var colorscheme editor.Colorscheme

			vs := gjson.Get(tools.SettingsContent(), "rs_settings.request_body.theme")
			tm := ""

			if vs.Exists() {
				tm = vs.String()
			} else {
				tm = "railscast"
			}

			if theme := runtime.Files.FindFile(editor.RTColorscheme, tm); theme != nil {
				if data, err := theme.Data(); err == nil {
					colorscheme = editor.ParseColorscheme(string(data))
				}
			}

			settingsEditor := editor.NewView(bufferSettings)
			settingsEditor.SetRuntimeFiles(runtime.Files)
			settingsEditor.SetColorscheme(colorscheme)
			settingsEditor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
					case tcell.KeyCtrlS:
						tools.SaveBuffer(bufferSettings, tools.SettingsFile())
						app.Stop()
						return nil
				}

				return event
			})

			app.SetRoot(settingsEditor, true)

			if err := app.Run(); err != nil {
				log.Fatalf("%v", err)
			}
		},
	}

	return cmd
}
