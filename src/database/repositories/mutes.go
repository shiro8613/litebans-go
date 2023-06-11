package repositories

import (
	"fmt"
	"strings"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type mutesRepo database.DBConnection

func NewMutesRepo(db database.DBConnection) mutesRepo {
	return mutesRepo{DB: db.DB, SqlConsts: db.SqlConsts}
}

func (rp mutesRepo) GetAllLimited(limit int, offset int) ([]entites.MutesReturn, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Mutes)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.MutesReturn{}

	for rows.Next() {
		ret := entites.Mutes{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, Converter[entites.Mutes, entites.MutesReturn](ret))
	}

	return rets, nil
}

func (rp mutesRepo) GetById(id int) (entites.MutesReturn, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Mutes)
	row := rp.QueryRowx(sql, id)

	ret := entites.Mutes{}
	err := row.StructScan(&ret)

	if err != nil {
		return entites.MutesReturn{}, err
	}

	return Converter[entites.Mutes, entites.MutesReturn](ret), nil
} 

func (rp mutesRepo) GetByUuid(uuid string) ([]entites.MutesReturn, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Mutes)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.MutesReturn{}

	for rows.Next() {
		ret := entites.Mutes{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, Converter[entites.Mutes, entites.MutesReturn](ret))
	}

	return rets, nil
} 

func (rp mutesRepo) GetCount() (int, error) {
	sql := strings.Replace(rp.SqlConsts.Mutes, "SELECT * FROM", "SELECT COUNT( id ) FROM", -1)
	row := rp.QueryRowx(sql)

	var ret int
	err := row.Scan(&ret)

	if err != nil {
		return 0, err
	}

	return ret, nil
} 