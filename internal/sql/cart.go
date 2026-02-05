package sql

import (
	"canteen/internal/models"
	"canteen/internal/sql/utility"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
)

func SelectMyCart(uid ulid.ULID) ([]models.CartItem, error) {
	var cart = make([]models.CartItem, 0)

	query := `
		SELECT
			account_id,
			canteen_id,
			menu_id,
			quantity,
			price_per_item
		FROM
			cart
		WHERE
			account_id = ?
	`

	rows, err := mariadb.Db.Query(query, uid)

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

func CheckAndInsertToCart(tx *sql.Tx, items []models.CartItem, uid ulid.ULID) error {
	var currentUserCanteen int
	err := tx.QueryRow("SELECT canteen_id FROM cart WHERE account_id = ?", uid).Scan(&currentUserCanteen)

	if err != nil {
		if err == sql.ErrNoRows {
			err := tx.QueryRow("SELECT canteen_id FROM menu WHERE menu_id = ?", items[0].MenuId).Scan(&currentUserCanteen)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	query := fmt.Sprintf(`
		SELECT
			canteen_id,
			price
		FROM
			menu
		WHERE menu_id IN (%s)
	`, utility.GeneratePlaceholders(len(items)))

	args := make([]any, len(items))

	for i, v := range items {
		args[i] = v.MenuId
	}

	rows, err := tx.Query(query, args...)

	if err != nil {
		return err
	}

	var index = 0
	for rows.Next() {
		if err := rows.Scan(
			&items[index].CanteenId,
			&items[index].PricePerItem,
		); err != nil {
			return err
		}

		index += 1
	}

	for _, item := range items {
		if item.CanteenId != currentUserCanteen {
			return errors.New("Every item must from the same canteen")
		}

		err := CheckMenuStock(item.MenuId, item.Quantity)
		if err != nil {
			str := fmt.Sprintf("The menu %d is currently has less stock than your request", item.MenuId)
			return errors.New(str)
		}
	}

	err = InsertToCart(tx, items, uid)
	if err != nil {
		return err
	}

	return nil

}

func InsertToCart(tx *sql.Tx, items []models.CartItem, uid ulid.ULID) error {
	placeholders := make([]string, 0, len(items))
	args := make([]any, 0, len(items)*5)

	for _, item := range items {
		placeholders = append(placeholders, "(?, ?, ?, ?, ?)")

		args = append(args, uid, item.CanteenId, item.MenuId, item.Quantity, item.PricePerItem)
	}

	query := fmt.Sprintf(`
			INSERT INTO 
				cart 
				(account_id, canteen_id, menu_id, quantity, price_per_item) 
			VALUES 
				%s
		`, strings.Join(placeholders, ", "))

	_, err := tx.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMyCartItems(uid ulid.ULID) error {
	query := `
		DELETE FROM
			cart
		WHERE
			account_id = ?
	`
	_, err := mariadb.Db.Exec(query, uid)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMyCartItemsAfterOrder(uid ulid.ULID, tx *sql.Tx) error {
	query := `
		DELETE FROM
			cart
		WHERE
			account_id = ?
	`
	_, err := tx.Exec(query, uid)
	if err != nil {
		return err
	}

	return nil
}
