package repositories

import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type MutesRepo database.DBConnection

func (rp MutesRepo) GetAllLimited(limit int, offset int) ([]entites.Mutes, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Mutes)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Mutes{}

	for rows.Next() {
		ret := entites.Mutes{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, ret)
	}

	return rets, nil
}

func (rp MutesRepo) GetById(id int) (entites.Mutes, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Mutes)
	row := rp.QueryRowx(sql, id)

	ret := entites.Mutes{}
	err := row.StructScan(ret)

	if err != nil {
		return entites.Mutes{}, err
	}

	return ret, nil
} 

func (rp MutesRepo) GetByUuid(uuid string) ([]entites.Mutes, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Mutes)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Mutes{}

	for rows.Next() {
		ret := entites.Mutes{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, ret)
	}

	return rets, nil
} 