package repositories


import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type HistoryRepo database.DBConnection

func (rp HistoryRepo) GetAllLimited(limit int, offset int) ([]entites.History, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.History)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.History{}

	for rows.Next() {
		ret := entites.History{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, ret)
	}

	return rets, nil
}

func (rp HistoryRepo) GetByUuid(uuid string) ([]entites.History, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.History)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.History{}

	for rows.Next() {
		ret := entites.History{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, ret)
	}

	return rets, nil
} 
