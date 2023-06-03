package database

import "github.com/jmoiron/sqlx"

type SqlConsts struct {
	Bans 		string
	History 	string
	Kicks 		string
	Mutes 		string
	Warnings 	string
}

type DBConnection struct {
	*sqlx.DB
	SqlConsts SqlConsts
}