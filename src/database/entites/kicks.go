package entites

type Kicks struct {
	Id					int		`db:"id"`
	Uuid				[]uint8	`db:"uuid"`
	Ip					[]uint8	`db:"ip"`
	Reason 				[]uint8	`db:"reason"`
	Banned_By_Uuid 		[]uint8	`db:"banned_by_uuid"`
	Banned_By_Name 		[]uint8	`db:"banned_by_name"`
	Time 				int		`db:"time"`
	Until 				int		`db:"until"`
	Template 			int		`db:"template"`
	Server_Scope 		[]uint8	`db:"server_scope"`
	Server_Origin 		[]uint8	`db:"server_origin"`
	Silent 				[]byte	`db:"silent"`
	IpBan 				[]byte	`db:"ipban"`
	IpBan_Wildcard 		[]byte	`db:"ipban_wildcard"`
	Active 				[]byte	`db:"active"`
}

type KicksReturn struct {
	Id					int		
	Uuid				string	
	Ip					string	
	Reason 				string	
	Banned_By_Uuid 		string	
	Banned_By_Name 		string	
	Time 				int		
	Until 				int		
	Template 			int		
	Server_Scope 		string	
	Server_Origin 		string	
	Silent 				string	
	IpBan 				string	
	IpBan_Wildcard 		string	
	Active 				string	
}