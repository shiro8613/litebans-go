package repositories

import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type KicksRepo database.DBConnection

func (rp KicksRepo) GetAllLimited(limit int, offset int) ([]entites.Kicks, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Kicks)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Kicks{}

	for rows.Next() {
		ret := entites.Kicks{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, ret)
	}

	return rets, nil
}

func (rp KicksRepo) GetById(id int) (entites.Kicks, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Kicks)
	row := rp.QueryRowx(sql, id)

	ret := entites.Kicks{}
	err := row.StructScan(ret)

	if err != nil {
		return entites.Kicks{}, err
	}

	return ret, nil
} 

func (rp KicksRepo) GetByUuid(uuid string) ([]entites.Kicks, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Kicks)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.Kicks{}

	for rows.Next() {
		ret := entites.Kicks{}
		err := rows.StructScan(ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, ret)
	}

	return rets, nil
} 
