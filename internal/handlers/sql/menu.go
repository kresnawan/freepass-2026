package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
)

func InsertMenu(menu models.Menu) (int64, error) {
	res, err := mariadb.Db.Exec(`
		INSERT INTO
			menu (canteen_id, menu_name, price)
		VALUES
			(?, ?, ?)`,
		menu.CanteenId,
		menu.MenuName,
		menu.Price)

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

func GetAllMenu() ([]models.Menu, error) {
	var menus = make([]models.Menu, 0)

	query := `
		SELECT
			menu_id,
			canteen_id,
			menu_name,
			price
		FROM
			menu
	`

	rows, err := mariadb.Db.Query(query)

	if err != nil {
		return menus, err
	}

	for rows.Next() {
		var menu models.Menu
		err := rows.Scan(
			&menu.MenuId,
			&menu.CanteenId,
			&menu.MenuName,
			&menu.Price,
		)
		if err != nil {
			return menus, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func GetAllMenuByCanteen(cid string) ([]models.Menu, error) {
	var menus = make([]models.Menu, 0)

	query := `
		SELECT
			*
		FROM
			menu
		WHERE
			canteen_id = ?
	`

	rows, err := mariadb.Db.Query(query, cid)

	if err != nil {
		return menus, err
	}

	for rows.Next() {
		var menu models.Menu
		err := rows.Scan(
			&menu.CanteenId,
			&menu.MenuId,
			&menu.MenuName,
			&menu.Price,
		)
		if err != nil {
			return menus, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}
