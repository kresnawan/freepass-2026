package sql

import "canteen/internal/storage/mariadb"

func InsertMenu(cid string, menuName string, price int) (int64, error) {
	res, err := mariadb.Db.Exec(`
		INSERT INTO
			menu (canteen_id, menu_name, price)
		VALUES
			(?, ?, ?)`,
		cid,
		menuName,
		price)

	if err != nil {
		return 0, err
	}

	lastInsertedId, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

func UpdateMenu(mid string, menuName string, price int) (int64, error) {
	res, err := mariadb.Db.Exec(`
		UPDATE
			menu
		SET
			menu_name = ?,
			price = ?
		WHERE
			menu_id = ?	
		`,
		menuName,
		price,
		mid)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteMenuById(mid string) (int64, error) {
	res, err := mariadb.Db.Exec(`
		DELETE FROM
			menu
		WHERE
			menu_id = ?
		`,
		mid)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
