package sql

import (
	"canteen/internal/storage/mariadb"
)

func AddCanteen(name string) (int64, error) {
	res, err := mariadb.Db.Exec("INSERT INTO canteen (name) VALUES (?)", name)

	if err != nil {
		return 0, err
	}

	inserted_id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return inserted_id, nil
}
