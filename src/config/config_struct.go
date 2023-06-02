package config

type Config_Web struct {
	Listen				string 					`yaml:"listen"`
	ViewsDirectory		string 					`yaml:"views_directory"`
	Name				string					`yaml:"name"`
}

type Config_Database struct {
	Host				string 					`yaml:"host"`
	Port				int 					`yaml:"port"`
	DBName				string					`yaml:"db_name"`
	Username			string 					`yaml:"username"`
	Password			string 					`yaml:"password"`
	TablePrefix			string 					`yaml:"table_prefix"`
}


type Config_Settings_Console struct {
	Aliases				[]string				`yaml:"aliases"`
	Name				string					`yaml:"name"`
	Image				string					`yaml:"image"`
}

type Config_Settings struct {
	Lang			 	string					`yaml:"lang"`
	ShowInactiveBans 	bool					`yaml:"show_inactive_bans"`
	ShowSilentBans		bool 					`yaml:"show_silent_bans"`
	Timezone 			string 					`yaml:"timezone"`
	DateFormat 			string					`yaml:"date_format"`
	Console 			Config_Settings_Console `yaml:"console"`
	AvatarSource		string					`yaml:"avatar_source"`
}

type Config struct {
	Web					Config_Web				`yaml:"web"`
	Database			Config_Database			`yaml:"database"`
	Settings			Config_Settings			`yaml:"settings"`
}
