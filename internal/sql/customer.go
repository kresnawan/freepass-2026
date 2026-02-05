package sql

import (
	"canteen/internal/models"
	"canteen/internal/sql/utility"
	"canteen/internal/storage/mariadb"

	"github.com/oklog/ulid/v2"
)

func InsertCustomerProfile(acc models.Account) error {
	uid, err := utility.GetUnusedId("account_id", "accounts")

	if err != nil {
		return err
	}
	tx, err := mariadb.Db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	acc.Role = "customer"

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

func GetCustomerProfile(uid ulid.ULID) (models.CustomerProfile, error) {
	var obj models.CustomerProfile

	query := `
		SELECT
			a.username,
			a.email,
			a.first_name,
			a.last_name,
			c.phone_number,
			c.canteen_points,
			c.instagram,
			c.bio
		FROM
			accounts a
		JOIN
			customer_profile c ON a.account_id = c.account_id
		WHERE
			a.account_id = ?
	`

	err := mariadb.Db.QueryRow(query, uid).Scan(
		&obj.Username,
		&obj.Email,
		&obj.FirstName,
		&obj.LastName,
		&obj.PhoneNumber,
		&obj.CanteenPoints,
		&obj.Instagram,
		&obj.Bio,
	)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func UpdateCustomerProfile(uid ulid.ULID, profile models.CustomerProfile) error {
	updateAccountQuery := `
		UPDATE
			accounts
		SET
			username = ?,
			first_name = ?,
			last_name = ?
		WHERE
			account_id = ?
	`
	updateProfileQuery := `
		UPDATE
			customer_profile
		SET
			phone_number = ?,
			instagram = ?,
			bio = ?
		WHERE
			account_id = ?
	`

	tx, err := mariadb.Db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(updateProfileQuery, profile.PhoneNumber, profile.Instagram, profile.Bio, uid)
	if err != nil {
		return err
	}

	_, err = tx.Exec(updateAccountQuery, profile.Username, profile.FirstName, profile.LastName, uid)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
