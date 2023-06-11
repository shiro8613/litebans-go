package repositories

import (
	"fmt"
	"strings"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type bansRepo database.DBConnection

func NewBansRepo(db database.DBConnection) bansRepo {
	return bansRepo{DB: db.DB, SqlConsts: db.SqlConsts}
}

func (rp bansRepo) GetAllLimited(limit int, offset int) ([]entites.BansReturn, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Bans)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.BansReturn{}

	for rows.Next() {
		ret := entites.Bans{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}
	
		rets = append(rets, Converter[entites.Bans, entites.BansReturn](ret))
	}

	return rets, nil
}

func (rp bansRepo) GetById(id int) (entites.BansReturn, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Bans)
	row := rp.QueryRowx(sql, id)

	ret := entites.Bans{}
	err := row.StructScan(&ret)

	if err != nil {
		return entites.BansReturn{}, err
	}

	return Converter[entites.Bans, entites.BansReturn](ret), nil
} 

func (rp bansRepo) GetByUuid(uuid string) ([]entites.BansReturn, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Bans)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.BansReturn{}

	for rows.Next() {
		ret := entites.Bans{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, Converter[entites.Bans, entites.BansReturn](ret))
	}

	return rets, nil
} 

func (rp bansRepo) GetCount() (int, error) {
	sql := strings.Replace(rp.SqlConsts.Bans, "SELECT * FROM", "SELECT COUNT( id ) FROM", -1)
	row := rp.QueryRowx(sql)

	var ret int
	err := row.Scan(&ret)

	if err != nil {
		return 0, err
	}

	return ret, nil
} 