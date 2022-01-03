package tools

import (
	"io/ioutil"
	"github.com/tidwall/sjson"
	"github.com/tidwall/pretty"
)

func SettingsContent() string {
	stgFile, err := ioutil.ReadFile(settingsFile)

	if err != nil {
		panic(err)
	}

	return string(stgFile)
}

func UpdateSettings(value bool) {
	settings, _ := sjson.Set(settingsFile, "rs_settings.show_update", value)

	prettySettings := pretty.Pretty([]byte(settings))

	err := ioutil.WriteFile(settingsFile, []byte(string(prettySettings)), 0644)

	if err != nil {
		panic(err)
	}
}

func SetDefaultSettings() {
	defaultSettings := `
		{
			"rs_settings": {
				"show_update": true
				"enable_mouse": true
				"request_body": {
					"theme": "railscast"
				}
			}
		}
	`

	prettySettings := pretty.Pretty([]byte(defaultSettings))

	err := ioutil.WriteFile(settingsFile, []byte(string(prettySettings)), 0644)

	if err != nil {
		panic(err)
	}
}
