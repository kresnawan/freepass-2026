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
	OwnerName   ulid.ULID `json:"owner_name"`
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

	return nil
}

func SelectAllCanteenOwnership() ([]CanteenOwnership, error) {
	var owns = make([]CanteenOwnership, 0)

	query := `
		SELECT
			co.owner_id,
			co.canteen_id,
			acc.username AS owner_name,
			c.name AS canteen_name,
			co.owned_at
		FROM
			canteen_ownership co
		JOIN
			owner_profile op ON co.owner_id = op.account_id
		JOIN
			accounts acc ON acc.account_id = op.account_id
		JOIN
			canteen c ON c.canteen_id = co.canteen_id

	`

	res, err := mariadb.Db.Query(query)

	if err != nil {
		return owns, err
	}

	for res.Next() {
		var own CanteenOwnership
		if err := res.Scan(
			&own.CanteenId,
			&own.CanteenName,
			&own.OwnedAt,
			&own.OwnerId,
			&own.OwnerName,
		); err != nil {
			return owns, err
		}

		owns = append(owns, own)
	}

	return owns, nil
}

func SelectCanteenOwner(cid string) ([]models.OwnerProfile, error) {
	var owners = make([]models.OwnerProfile, 0)

	query := `
		SELECT
			acc.username,
			acc.email,
			acc.first_name,
			acc.lastname
		FROM
			canteen_ownership co
		JOIN
			owner_profile op ON co.owner_id = op.account_id
		JOIN
			accounts acc ON acc.account_id = op.account_id
		WHERE
			co.canteen_id = ?
	`

	res, err := mariadb.Db.Query(query, cid)

	if err != nil {
		return owners, err
	}

	for res.Next() {
		var owner models.OwnerProfile
		if err := res.Scan(
			&owner.Email,
			&owner.Username,
			&owner.FirstName,
			&owner.LastName,
		); err != nil {
			return owners, err
		}

		owners = append(owners, owner)
	}

	return owners, nil
}
