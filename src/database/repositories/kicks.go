package repositories

import (
	"fmt"
	"strings"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type kicksRepo database.DBConnection

func NewKicksRepo(db database.DBConnection) kicksRepo {
	return kicksRepo{DB: db.DB, SqlConsts: db.SqlConsts}
}

func (rp kicksRepo) GetAllLimited(limit int, offset int) ([]entites.KicksReturn, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Kicks)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.KicksReturn{}

	for rows.Next() {
		ret := entites.Kicks{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, Converter[entites.Kicks, entites.KicksReturn](ret))
	}

	return rets, nil
}

func (rp kicksRepo) GetById(id int) (entites.KicksReturn, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Kicks)
	row := rp.QueryRowx(sql, id)

	ret := entites.Kicks{}
	err := row.StructScan(&ret)

	if err != nil {
		return entites.KicksReturn{}, err
	}

	return Converter[entites.Kicks, entites.KicksReturn](ret), nil
} 

func (rp kicksRepo) GetByUuid(uuid string) ([]entites.KicksReturn, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Kicks)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.KicksReturn{}

	for rows.Next() {
		ret := entites.Kicks{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, Converter[entites.Kicks, entites.KicksReturn](ret))
	}

	return rets, nil
} 

func (rp kicksRepo) GetCount() (int, error) {
	sql := strings.Replace(rp.SqlConsts.Kicks, "SELECT * FROM", "SELECT COUNT( id ) FROM", -1)
	row := rp.QueryRowx(sql)

	var ret int
	err := row.Scan(&ret)

	if err != nil {
		return 0, err
	}

	return ret, nil
} 