package entites

type Kicks struct {
	Id					int		`db:"id"`
	Uuid				string	`db:"uuid"`
	Ip					string	`db:"ip"`
	Reason 				string	`db:"reason"`
	Banned_By_Uuid 		string	`db:"banned_by_uuid"`
	Banned_By_Name 		string	`db:"banned_by_name"`
	Time 				int		`db:"time"`
	Until 				int		`db:"until"`
	Template 			int		`db:"template"`
	Server_Scope 		string	`db:"server_scope"`
	Server_Origin 		string	`db:"server_origin"`
	Silent 				bool	`db:"silent"`
	IpBan 				bool	`db:"ipban"`
	IpBan_Wildcard 		bool	`db:"ipban_wildcard"`
	Active 				bool	`db:"active"`
}