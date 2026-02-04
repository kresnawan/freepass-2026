package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
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
			&own.OwnerId,
			&own.CanteenId,
			&own.OwnerName,
			&own.CanteenName,
			&own.OwnedAt,
		); err != nil {
			return owns, err
		}

		owns = append(owns, own)
	}

	return owns, nil
}

func SelectCanteenOwners() ([]models.OwnerProfile, error) {
	var owners = make([]models.OwnerProfile, 0)

	query := `
		SELECT
			account_id,
			username,
			email,
			first_name,
			last_name,
			role
		FROM
			accounts
		WHERE
			role = 'owner'
	`

	res, err := mariadb.Db.Query(query)

	if err != nil {
		return owners, err
	}

	for res.Next() {
		var owner models.OwnerProfile
		if err := res.Scan(
			&owner.AccountId,
			&owner.Email,
			&owner.Username,
			&owner.FirstName,
			&owner.LastName,
			&owner.Role,
		); err != nil {
			return owners, err
		}

		owners = append(owners, owner)
	}

	return owners, nil
}

func SelectCanteenOwner(cid string) ([]models.OwnerProfile, error) {
	var owners = make([]models.OwnerProfile, 0)

	query := `
		SELECT
			acc.account_id,
			acc.username,
			acc.email,
			acc.first_name,
			acc.last_name,
			acc.role
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
			&owner.AccountId,
			&owner.Username,
			&owner.Email,
			&owner.FirstName,
			&owner.LastName,
			&owner.Role,
		); err != nil {
			return owners, err
		}

		owners = append(owners, owner)
	}

	return owners, nil
}

func InsertOwnership(cid string, uid ulid.ULID) error {
	query := `
		INSERT INTO
			canteen_ownership
			(owner_id, canteen_id)
		VALUES
			(?, ?)
	`
	_, err := mariadb.Db.Exec(query, uid, cid)

	if err != nil {
		return err
	}

	return nil
}

func CheckOwnership(uid ulid.ULID, cid int) (bool, error) {
	query := `
		SELECT
			canteen_id
		FROM
			canteen_ownership
		WHERE
			owner_id = ? AND canteen_id = ?

	`
	var canteenId int
	err := mariadb.Db.QueryRow(query, uid, cid).Scan(&canteenId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
