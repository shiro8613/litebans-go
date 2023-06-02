package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func create(path string) error {

	config := Config {
		Web: Config_Web {
			Listen: "127.0.0.1:8080",
			ViewsDirectory: "./views",
		},
		Database: Config_Database {
			Host: "127.0.0.1",
			Port: 3306,
			DBName: "Litebans",
			Username: "",
			Password: "",
			TablePrefix: "Litebans_",
		},
		Settings: Config_Settings {
			Lang: "JP",
			ShowInactiveBans: true,
			ShowSilentBans: true,
			Timezone: "UTC",
			DateFormat: "YYYY/MM/DD hh:mm:ss",
			Console: Config_Settings_Console {
				Aliases: []string{"console", "Console"},
				Name: "console",
				Image: "static/img/console.png",
			},
			AvatarSource: "https://cravatar.eu/avatar/{uuid}/25",
		},
	}

	b, err := yaml.Marshal(config)

	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(b)

	return nil
}