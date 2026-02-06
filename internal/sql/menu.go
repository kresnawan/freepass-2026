package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"errors"
	"strconv"

	"github.com/oklog/ulid/v2"
)

func InsertMenu(menu models.Menu) error {
	tx, err := mariadb.Db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		INSERT INTO
			menu (canteen_id, menu_name, price)
		VALUES
			(?, ?, ?)`

	res, err := tx.Exec(query, menu.CanteenId, menu.MenuName, menu.Price)

	if err != nil {
		return err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	query = `
		INSERT INTO
			menu_stock
			(menu_id, stock)
		VALUES
			(?, ?)
	`

	_, err = tx.Exec(query, lastInsertId, 0)

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateMenu(mid string, menuName string, price int) error {
	query := `
		UPDATE
			menu
		SET
			menu_name = ?,
			price = ?
		WHERE
			menu_id = ?	
	`
	_, err := mariadb.Db.Exec(query, menuName, price, mid)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMenuById(mid string) error {
	_, err := mariadb.Db.Exec(`
		UPDATE
			menu
		SET
			is_removed = 1
		WHERE
			menu_id = ?
		`,
		mid)

	if err != nil {
		return err
	}

	return err
}

func GetAllMenu(page string) ([]models.Menu, error) {
	var menus = make([]models.Menu, 0)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return menus, err
	}

	query := `
		SELECT
			m.menu_id,
			m.canteen_id,
			m.menu_name,
			m.price,
			ms.stock
		FROM
			menu m
		JOIN
			menu_stock ms ON ms.menu_id = m.menu_id
		WHERE
			m.is_removed = 0
		LIMIT
			10
		OFFSET
			?
	`

	rows, err := mariadb.Db.Query(query, ((pageInt - 1) * 10))

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
			&menu.Stock,
		)
		if err != nil {
			return menus, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func GetAllMenuByCanteen(cid string, page string) ([]models.Menu, error) {
	var menus = make([]models.Menu, 0)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return menus, err
	}

	query := `
		SELECT
			m.menu_id,
			m.canteen_id,
			m.menu_name,
			m.price,
			ms.stock
		FROM
			menu m
		JOIN
			menu_stock ms ON ms.menu_id = m.menu_id
		WHERE
			m.is_removed = 0 AND m.canteen_id = ?
		LIMIT
			10
		OFFSET
			?
	`

	rows, err := mariadb.Db.Query(query, cid, ((pageInt - 1) * 10))

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
			&menu.Stock,
		)
		if err != nil {
			return menus, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func CheckMenuStock(menuid int, q int) error {
	query := `
		SELECT
			stock
		FROM
			menu_stock
		WHERE
			menu_id = ?
	`
	var stock int

	err := mariadb.Db.QueryRow(query, menuid).Scan(&stock)
	if err != nil {
		return err
	}

	if stock < q {
		return errors.New("Menu unavailable (no stock)")
	}

	return nil
}

func AddStock(menuid int, q int) error {
	query := `
		SELECT
			stock
		FROM
			menu_stock ms
		JOIN
			menu m ON m.menu_id = ms.menu_id
		WHERE
			ms.menu_id = ? AND m.is_removed = 0
	`

	var stock int

	err := mariadb.Db.QueryRow(query, menuid).Scan(&stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("The menu is either removed or not exist")
		} else {
			return err
		}
	}

	stock += q
	query = `
		UPDATE
			menu_stock
		SET
			stock = ?
		WHERE
			menu_id = ?
	`

	_, err = mariadb.Db.Exec(query, stock, menuid)
	if err != nil {
		return err
	}

	return nil
}

func DropStock(menuid int, q int, tx *sql.Tx) error {
	query := `
		SELECT
			stock
		FROM
			menu_stock
		WHERE
			menu_id = ?
	`

	var stock int

	err := tx.QueryRow(query, menuid).Scan(&stock)
	if err != nil {
		return err
	}

	stock -= q
	query = `
		UPDATE
			menu_stock
		SET
			stock = ?
		WHERE
			menu_id = ?
	`

	_, err = tx.Exec(query, stock, menuid)
	if err != nil {
		return err
	}

	return nil
}

func CheckMenuOwnership(mid int, uid ulid.ULID) (bool, error) {
	query := `
		SELECT
			m.menu_id
		FROM
			canteen_ownership co
		JOIN
			menu m ON m.canteen_id = co.canteen_id
		WHERE
			co.owner_id = ? AND m.menu_id = ?
	`
	var canteenId int
	err := mariadb.Db.QueryRow(query, uid, mid).Scan(&canteenId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
