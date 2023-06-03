package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shiro8613/litebans-go/src/config"
)

func newDatabaseConnection(config config.Config_Database) (DBConnection, error){
	
	connect := fmt.Sprintf("%s:%s@tmp(%s:%d/%s)", config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := sqlx.Open("mysql", connect)
	
	if err != nil {
		return DBConnection{}, err
	}

	defer db.Close()

	return DBConnection{
		DB: db,
		SqlConsts: SqlConsts{
			Bans: fmt.Sprintf("%s_bans", config.TablePrefix),
			History: fmt.Sprintf("%s_history", config.TablePrefix),
			Kicks: fmt.Sprintf("%s_kicks", config.TablePrefix),
			Mutes: fmt.Sprintf("%s_mutes", config.TablePrefix),
			Warnings: fmt.Sprintf("%s_warnings", config.TablePrefix),
		},
	}, nil
}