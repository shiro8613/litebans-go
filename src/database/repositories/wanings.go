package repositories

import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type WarningsRepo database.DBConnection

func (rp WarningsRepo) GetAllLimited(limit int, offset int) ([]entites.Warnings, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Warning)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Warnings{}

	for rows.Next() {
		ret := entites.Warnings{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, ret)
	}

	return rets, nil
}

func (rp WarningsRepo) GetById(id int) (entites.Warnings, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Warning)
	row := rp.QueryRowx(sql, id)

	ret := entites.Warnings{}
	err := row.StructScan(ret)

	if err != nil {
		return entites.Warnings{}, err
	}

	return ret, nil
} 

func (rp WarningsRepo) GetByUuid(uuid string) ([]entites.Warnings, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Warning)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Warnings{}

	for rows.Next() {
		ret := entites.Warnings{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, ret)
	}

	return rets, nil
} 