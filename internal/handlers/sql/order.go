package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

func SelectOrderByCanteenId(cid string) ([]models.Order, error) {
	var orders []models.Order = make([]models.Order, 0)

	rows, err := mariadb.Db.Query(`
	SELECT 
		*
	FROM 
		`+"`order`"+`
	WHERE
		canteen_id = ?	
	`, cid)

	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.OrderId,
			&order.CustomerId,
			&order.CanteenId,
			&order.Status,
			&order.IsPaid,
			&order.PaidAt,
			&order.CreatedAt,
		); err != nil {
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func SelectOrderById(oid string) ([]models.Order, error) {
	var orders []models.Order = make([]models.Order, 0)

	parsedId, _ := ulid.Parse(oid)

	rows, err := mariadb.Db.Query(`
	SELECT 
		*
	FROM 
		order
	WHERE
		order_id = ?	
	`, parsedId)

	if err != nil {
		return orders, nil
	}

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.OrderId,
			&order.CustomerId,
			&order.CanteenId,
			&order.Status,
			&order.IsPaid,
			&order.PaidAt,
		); err != nil {
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return orders, nil
	}

	return orders, nil
}

func SelectMyOrder(uid ulid.ULID) ([]models.Order, error) {
	var orders []models.Order = make([]models.Order, 0)

	rows, err := mariadb.Db.Query(`
	SELECT 
		order_id,
		customer_id,
		canteen_id,
		status,
		is_paid,
		paid_at,
		created_at
	FROM 
		`+"`order`"+`
	WHERE
		customer_id = ?	
	`, uid)

	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.OrderId,
			&order.CustomerId,
			&order.CanteenId,
			&order.Status,
			&order.IsPaid,
			&order.PaidAt,
			&order.CreatedAt,
		); err != nil {
			return orders, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func InsertOrder(customerId ulid.ULID, canteenId int, tx *sql.Tx) (ulid.ULID, error) {
	oid := ulid.Make()
	query := `
		INSERT INTO 
			` + "`order`" + ` 
			(order_id, customer_id, canteen_id) 
		VALUES 
			(?, ?, ?)
	`
	_, err := tx.Exec(query, oid, customerId, canteenId)
	if err != nil {
		return oid, err
	}

	return oid, nil
}

func InsertOrderItems(oid ulid.ULID, items []models.CartItem, tx *sql.Tx) error {
	args := make([]any, 0, len(items)*4)
	placeholders := make([]string, 0, len(items))

	for _, item := range items {
		placeholders = append(placeholders, "(?, ?, ?, ?)")

		args = append(args, item.MenuId, oid, item.Quantity, item.PricePerItem)
	}

	query := fmt.Sprintf(`
		INSERT INTO
			order_items
			(menu_id, order_id, quantity, price_per_item)
		VALUES
			%s
	`, strings.Join(placeholders, ", "))

	_, err := tx.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func UpdateOrderStatus(oid ulid.ULID) error {
	var currentStatus int
	query := `
		SELECT
			status
		FROM
			` + "`order`" + `
		WHERE
			order_id = ?
	`
	err := mariadb.Db.QueryRow(query, oid).Scan(&currentStatus)
	if err != nil {
		return err
	}

	if currentStatus >= 9 {
		return errors.New("This order been completed")
	}

	if currentStatus == 6 {
		return errors.New("This order's payment did not completed yet")
	}

	currentStatus = currentStatus + 1

	query = `
		UPDATE
			` + "`order`" + `
		SET
			status = ?
		WHERE
			order_id = ?
	`

	_, err = mariadb.Db.Exec(query, currentStatus, oid)
	if err != nil {
		return err
	}

	return nil
}

func PayOrder(oid ulid.ULID) error {
	query := `
		UPDATE
			` + "`order`" + `
		SET
			is_paid = 1,
			paid_at = ?,
			status = ?
		WHERE
			order_id = ?
	`
	_, err := mariadb.Db.Exec(query, time.Now().UTC(), 7, oid)
	if err != nil {
		return err
	}

	return nil
}
