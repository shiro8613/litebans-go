package entites

type History struct {
	Id	 	int		`db:"id"`
	Date	[]uint8	`db:"date"`
	Name	[]uint8	`db:"name"`
	Uuid	[]uint8	`db:"uuid"`
	Ip		[]uint8	`db:"ip"`
}

type HistoryRetrun struct {
	Id	 	int		`db:"id"`
	Date	string	`db:"date"`
	Name	string	`db:"name"`
	Uuid	string	`db:"uuid"`
	Ip		string	`db:"ip"`
}