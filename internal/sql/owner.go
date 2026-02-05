package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"time"

	"github.com/oklog/ulid/v2"
)

type CanteenOwnership struct {
	OwnerId     ulid.ULID `json:"owner_id"`
	CanteenId   int       `json:"canteen_id"`
	OwnerName   string    `json:"owner_name"`
	CanteenName string    `json:"canteen_name"`
	OwnedAt     time.Time `json:"owned_at"`
}

func InsertOwnerProfile(acc models.Account) error {
	var uid ulid.ULID = ulid.Make()
	tx, err := mariadb.Db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	acc.Role = "owner"

	err = InsertAccount(acc, tx, uid)

	if err != nil {
		return err
	}

	query := `
		INSERT INTO
			owner_profile (account_id)
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

func UpdateOwnerProfile(uid ulid.ULID, profile models.OwnerProfileForEdit) error {
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
			phone_number = ?
		WHERE
			account_id = ?
	`

	tx, err := mariadb.Db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(updateProfileQuery, profile.PhoneNumber)
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
