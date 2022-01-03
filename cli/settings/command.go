package settings

import (
	"github.com/spf13/cobra"
)

func SettingsCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Update Settings Or Change it",
		Long:  `Update Resto settings like enable mouse or change editor theme`,
	}

	cmd.AddCommand(SettingsOpen(), SettingsSet())

	return cmd
}
