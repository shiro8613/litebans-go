package entites

type History struct {
	Id	 	int		`db:"id"`
	Date	string	`db:"date"`
	Name	string	`db:"name"`
	Uuid	string	`db:"uuid"`
	Ip		string	`db:"ip"`
}