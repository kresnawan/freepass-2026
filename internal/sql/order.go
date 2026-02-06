package sql

import (
	"canteen/internal/models"
	"canteen/internal/sql/utility"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

func SelectOrderByCanteenId(cid string, status string, page string) ([]models.OrderClean, error) {
	var orders []models.OrderClean = make([]models.OrderClean, 0)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return orders, err
	}

	rows, err := mariadb.Db.Query(`
	SELECT 
		order_id,
		customer_id,
		canteen_id,

		CASE
			WHEN status = 6 THEN 'Waiting payment'
			WHEN status = 7 THEN 'Cooking'
			WHEN status = 8 THEN 'Ready'
			WHEN status = 9 THEN 'Completed'
			ELSE ''
		END AS status,

		is_paid,
		paid_at,
		created_at
	FROM 
		`+"`order`"+`
	WHERE
		canteen_id = ? AND status = ?
	LIMIT 
		10
	OFFSET
		?
	`, cid, status, ((pageInt - 1) * 10))

	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order models.OrderClean
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

func SelectOrderById(oid ulid.ULID) (models.OrderClean, error) {
	var order models.OrderClean

	query := `
		SELECT 
			order_id,
			customer_id,
			canteen_id,
			
		CASE
			WHEN status = 6 THEN 'Waiting payment'
			WHEN status = 7 THEN 'Cooking'
			WHEN status = 8 THEN 'Ready'
			WHEN status = 9 THEN 'Completed'
			ELSE ''
		END AS status,

			is_paid,
			paid_at,
			created_at
		FROM 
			` + "`order`" + `
		WHERE
			order_id = ?
	`

	err := mariadb.Db.QueryRow(query, oid).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.CanteenId,
		&order.Status,
		&order.IsPaid,
		&order.PaidAt,
		&order.CreatedAt,
	)

	if err != nil {
		return order, err
	}

	return order, nil
}

func SelectMyOrder(uid ulid.ULID, status string, page string) ([]models.OrderClean, error) {
	var orders []models.OrderClean = make([]models.OrderClean, 0)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return orders, err
	}

	rows, err := mariadb.Db.Query(`
		SELECT 
			order_id,
			customer_id,
			canteen_id,
		
			CASE
				WHEN status = 6 THEN 'Waiting payment'
				WHEN status = 7 THEN 'Cooking'
				WHEN status = 8 THEN 'Ready'
				WHEN status = 9 THEN 'Completed'
				ELSE ''
				END AS status,

			is_paid,
			paid_at,
			created_at
		FROM 
			`+"`order`"+`
		WHERE
			customer_id = ? AND status = ?
		LIMIT
			10
		OFFSET
			?
	`, uid, status, ((pageInt - 1) * 10))

	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order models.OrderClean
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
	oid, err := utility.GetUnusedId("order_id", "order")

	if err != nil {
		return oid, nil
	}

	query := `
		INSERT INTO 
			` + "`order`" + ` 
			(order_id, customer_id, canteen_id) 
		VALUES 
			(?, ?, ?)
	`
	_, err = tx.Exec(query, oid, customerId, canteenId)
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

	for _, item := range items {
		err := DropStock(item.MenuId, item.Quantity, tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateOrderStatus(oid ulid.ULID) (int, error) {
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
		return 0, err
	}

	if currentStatus >= 9 {
		return 0, errors.New("This order been completed")
	}

	if currentStatus == 6 {
		return 0, errors.New("This order's payment did not completed yet")
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
		return 0, err
	}

	return currentStatus, nil
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
	_, err := mariadb.Db.Exec(query, time.Now(), 7, oid)
	if err != nil {
		return err
	}

	return nil
}

func GetOrderItems(oid ulid.ULID) ([]models.OrderItem, error) {
	query := `
		SELECT
			menu_id,
			quantity,
			price_per_item
		FROM 
			order_items
		WHERE
			order_id = ?
	`
	var items []models.OrderItem

	rows, err := mariadb.Db.Query(query, oid)
	if err != nil {
		return items, err
	}

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(
			&item.MenuId,
			&item.Quantity,
			&item.PricePerItem,
		); err != nil {
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func GetOrderDetails(oid ulid.ULID) (models.OrderDetails, error) {
	var res models.OrderDetails
	var totalPrice int = 0

	order, err := SelectOrderById(oid)
	if err != nil {
		return res, err
	}

	orderItems, err := GetOrderItems(oid)
	if err != nil {
		return res, err
	}

	for _, item := range orderItems {
		totalPrice += item.PricePerItem * item.Quantity
	}

	res = models.OrderDetails{
		OrderClean: order,
		TotalPrice: totalPrice,
		Items:      orderItems,
	}

	return res, err
}

func CheckOrderOwnership(oid ulid.ULID, uid ulid.ULID) (bool, error) {
	query := `
		SELECT
			o.order_id
		FROM
			` + "`order`" + ` o
		JOIN
			canteen_ownership co ON co.canteen_id = o.canteen_id
		WHERE
			o.order_id = ? AND co.owner_id = ?
	`
	var canteenId ulid.ULID
	err := mariadb.Db.QueryRow(query, oid, uid).Scan(&canteenId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func CheckOrderCustomerOwnership(oid ulid.ULID, uid ulid.ULID) (bool, error) {
	query := `
		SELECT
			customer_id
		FROM
			` + "`order`" + `
		WHERE
			order_id = ? AND customer_id = ?
	`
	var customerId ulid.ULID
	err := mariadb.Db.QueryRow(query, oid, uid).Scan(&customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
