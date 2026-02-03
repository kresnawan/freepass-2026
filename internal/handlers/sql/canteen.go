package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"

	"github.com/oklog/ulid/v2"
)

func AddCanteen(name string) (int64, error) {
	res, err := mariadb.Db.Exec(`
	INSERT INTO 
		canteen (name) 
	VALUES (?)`, name)

	if err != nil {
		return 0, err
	}

	inserted_id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return inserted_id, nil
}

func SelectCanteen() ([]models.Canteen, error) {
	var canteen_array = make([]models.Canteen, 0)
	rows, err := mariadb.Db.Query(`
	SELECT 
		canteen_id, 
		name 
	FROM 
		canteen`)

	if err != nil {
		return canteen_array, err
	}

	for rows.Next() {
		var canteen models.Canteen
		if err := rows.Scan(&canteen.Canteen_id, &canteen.Name); err != nil {
			return canteen_array, err
		}

		canteen_array = append(canteen_array, canteen)
	}

	if err := rows.Err(); err != nil {
		return canteen_array, err
	}

	return canteen_array, nil
}

func DeleteCanteen(id string) (int64, error) {
	res, err := mariadb.Db.Exec(`
	DELETE FROM 
		canteen 
	WHERE 
		canteen_id = ?`, id)

	if err != nil {
		return 0, err
	}

	rows_affected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows_affected, nil
}

func SelectOwnedCanteen(owid string) ([]models.Canteen, error) {
	var canteen_array = make([]models.Canteen, 0)
	parsedId, _ := ulid.Parse(owid)

	rows, err := mariadb.Db.Query(`
	SELECT 
		c.*
	FROM 
		canteen c
	INNER JOIN
		canteen_ownership o
	ON
		c.canteen_id = o.canteen_id
	WHERE
		o.owner_id = ?
		`, parsedId)

	if err != nil {
		return canteen_array, err
	}

	for rows.Next() {
		var canteen models.Canteen
		if err := rows.Scan(&canteen.Canteen_id, &canteen.Name); err != nil {
			return canteen_array, err
		}

		canteen_array = append(canteen_array, canteen)
	}

	if err := rows.Err(); err != nil {
		return canteen_array, err
	}

	return canteen_array, nil
}
