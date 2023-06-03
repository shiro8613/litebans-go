package repositories

import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type historyRepo database.DBConnection

func NewHistoryRepo(db database.DBConnection) historyRepo {
	return historyRepo{DB: db.DB, SqlConsts: db.SqlConsts}
}


func (rp historyRepo) GetAllLimited(limit int, offset int) ([]entites.HistoryRetrun, error) {
	sql := fmt.Sprintf("%s ORDER BY date DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.History)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.HistoryRetrun{}

	for rows.Next() {
		ret := entites.History{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, Converter[entites.History, entites.HistoryRetrun](ret))
	}

	return rets, nil
}

func (rp historyRepo) GetByUuid(uuid string) ([]entites.HistoryRetrun, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.History)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.HistoryRetrun{}

	for rows.Next() {
		ret := entites.History{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, Converter[entites.History, entites.HistoryRetrun](ret))
	}

	return rets, nil
} 
