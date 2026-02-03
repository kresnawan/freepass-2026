package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"

	"github.com/oklog/ulid/v2"
)

func InsertCustomerProfile(acc models.Account) error {
	var uid ulid.ULID = ulid.Make()
	tx, err := mariadb.Db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = InsertAccount(acc, tx, uid)

	if err != nil {
		return err
	}

	query := `
		INSERT INTO
			customer_profile (account_id)
		VALUES
			(?)
	`

	_, err = tx.Exec(query, uid)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetCustomerProfile(uid string) (models.CustomerProfile, error) {
	parsedId, _ := ulid.Parse(uid)
	var obj models.CustomerProfile

	query := `
		SELECT
			a.username,
			a.email,
			a.first_name,
			a.last_name,
			c.phone_number,
			c.canteen_points
		FROM
			accounts a
		INNER JOIN
			customer_profile c
		ON
			a.account_id = c.account_id
		WHERE
			account_id = ?
	`

	err := mariadb.Db.QueryRow(query, parsedId).Scan(
		&obj.Username,
		&obj.Email,
		&obj.FirstName,
		&obj.LastName,
		&obj.PhoneNumber,
		&obj.CanteenPoints,
	)

	if err != nil {
		return obj, err
	}

	return obj, nil
}
