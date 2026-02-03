package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"errors"

	"github.com/oklog/ulid/v2"
)

func SelectMyCart(uid ulid.ULID) ([]models.CartItem, error) {
	var cart = make([]models.CartItem, 0)

	query := `
		SELECT
			*
		FROM
			cart
	`

	rows, err := mariadb.Db.Query(query)

	if err != nil {
		return cart, err
	}

	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(
			&item.AccountId,
			&item.CanteenId,
			&item.MenuId,
			&item.Quantity,
			&item.PricePerItem,
		)
		if err != nil {
			return cart, err
		}

		cart = append(cart, item)
	}

	return cart, nil
}

func CheckAndInsertToCart(tx *sql.Tx, item models.CartItem) error {
	var canteenId int
	err := mariadb.Db.QueryRow("SELECT canteen_id FROM cart WHERE account_id = ?", item.AccountId).Scan(&canteenId)

	if err == sql.ErrNoRows {

	}

	var menuCanteenId int
	err = mariadb.Db.QueryRow("SELECT canteen_id FROM menu WHERE menu_id = ?", item.MenuId).Scan(&menuCanteenId)

	if err != nil {
		return err
	}

	if menuCanteenId != canteenId {
		return errors.New("Every item must from the same canteen")

	}

	err = InsertToCart(tx, item)
	if err != nil {
		return err
	}

	return nil
}

func InsertToCart(tx *sql.Tx, item models.CartItem) error {
	query := `
			INSERT INTO 
				cart 
				(account_id, canteen_id, menu_id, quantity, price_per_item) 
			VALUES 
				(?,?,?,?,?)
		`
	_, err := tx.Exec(query, item.AccountId, item.CanteenId, item.MenuId, item.Quantity, item.PricePerItem)

	if err != nil {
		return err
	}

	return nil
}
