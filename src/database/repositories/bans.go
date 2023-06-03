package repositories

import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type BansRepo database.DBConnection

func (rp BansRepo) GetAllLimited(limit int, offset int) ([]entites.Bans, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Bans)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Bans{}

	for rows.Next() {
		ret := entites.Bans{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, ret)
	}

	return rets, nil
}

func (rp BansRepo) GetById(id int) (entites.Bans, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Bans)
	row := rp.QueryRowx(sql, id)

	ret := entites.Bans{}
	err := row.StructScan(ret)

	if err != nil {
		return entites.Bans{}, err
	}

	return ret, nil
} 

func (rp BansRepo) GetByUuid(uuid string) ([]entites.Bans, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Bans)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Bans{}

	for rows.Next() {
		ret := entites.Bans{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, ret)
	}

	return rets, nil
} 
