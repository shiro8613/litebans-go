package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shiro8613/litebans-go/src/config"
)

func NewDatabaseConnection(config config.Config_Database) (DBConnection, error){
	
	connect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := sqlx.Open("mysql", connect)
	
	if err != nil {
		return DBConnection{}, err
	}

	//defer db.Close()
	
	sqlText := "SELECT * FROM" 
	bans := fmt.Sprintf("%s %sbans", sqlText, config.TablePrefix)
	kicks :=  fmt.Sprintf("%s %skicks", sqlText, config.TablePrefix)
	mutes := fmt.Sprintf("%s %smutes", sqlText, config.TablePrefix)
	history := fmt.Sprintf("%s %shistory", sqlText, config.TablePrefix)
	warning := fmt.Sprintf("%s %swarnings", sqlText, config.TablePrefix)

	return DBConnection{
		DB: db,
		SqlConsts: SqlConsts{
			Bans: bans,
			History: history,
			Kicks: kicks,
			Mutes: mutes,
			Warnings: warning,
		},
	}, nil
}