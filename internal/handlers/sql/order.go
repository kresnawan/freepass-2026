package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"

	"github.com/oklog/ulid/v2"
)

func SelectOrderByCanteenId(cid string) ([]models.Order, error) {
	var orders []models.Order = make([]models.Order, 0)

	rows, err := mariadb.Db.Query(`
	SELECT 
		*
	FROM 
		order
	WHERE
		canteen_id = ?	
	`, cid)

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
		*
	FROM 
		order
	WHERE
		customer_id = ?	
	`, uid)

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

func InsertOrder(customerId ulid.ULID, canteenId int, tx *sql.Tx) (ulid.ULID, error) {
	var oid ulid.ULID
	oid = ulid.Make()
	query := `
		INSERT INTO 
			order 
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
	query := `
		INSERT INTO
			order_items
			(menu_id, order_id, quantity, price_per_item)
		VALUES
			(?, ?, ?, ?)
	`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	for _, item := range items {
		_, err = stmt.Exec(item.MenuId)
		if err != nil {
			return err
		}
	}

	return nil
}
