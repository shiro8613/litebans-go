package repositories

import (
	"fmt"

	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/database/entites"
)

type warningsRepo database.DBConnection

func NewWarningsRepo(db database.DBConnection) warningsRepo {
	return warningsRepo{DB: db.DB, SqlConsts: db.SqlConsts}
}

func (rp warningsRepo) GetAllLimited(limit int, offset int) ([]entites.WarningsReturn, error) {
	sql := fmt.Sprintf("%s ORDER BY time DESC LIMIT :limit OFFSET :offset", rp.SqlConsts.Warnings)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"limit": limit, "offset": offset })

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.WarningsReturn{}

	for rows.Next() {
		ret := entites.Warnings{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}

		rets = append(rets, Converter[entites.Warnings, entites.WarningsReturn](ret))
	}

	return rets, nil
}

func (rp warningsRepo) GetById(id int) (entites.WarningsReturn, error) {
	sql := fmt.Sprintf("%s WHERE id = ? LIMIT 1", rp.SqlConsts.Warnings)
	row := rp.QueryRowx(sql, id)

	ret := entites.Warnings{}
	err := row.StructScan(&ret)

	if err != nil {
		return entites.WarningsReturn{}, err
	}

	return Converter[entites.Warnings, entites.WarningsReturn](ret), nil
} 

func (rp warningsRepo) GetByUuid(uuid string) ([]entites.WarningsReturn, error) {
	sql := fmt.Sprintf("%s WHERE uuid = :uuid", rp.SqlConsts.Warnings)
	rows, err := rp.NamedQuery(sql, map[string]interface{}{"uuid": uuid})

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rets := []entites.WarningsReturn{}

	for rows.Next() {
		ret := entites.Warnings{}
		err := rows.StructScan(&ret)
		if err != nil {
			return nil, err
		}
		
		rets = append(rets, Converter[entites.Warnings, entites.WarningsReturn](ret))
	}

	return rets, nil
} 