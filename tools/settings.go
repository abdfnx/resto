package tools

import (
	"io/ioutil"
	"github.com/tidwall/sjson"
)

func SettingsContent() string {
	stgFile, err := ioutil.ReadFile(settingsFile)

	if err != nil {
		panic(err)
	}

	return string(stgFile)
}

func UpdateSettings(value string) {
	settings, _ := sjson.Set(settingsFile, "rs_settings.show_update", value)

	err := ioutil.WriteFile(settingsFile, []byte(settings), 0644)

	if err != nil {
		panic(err)
	}
}
