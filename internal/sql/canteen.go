package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"

	"github.com/oklog/ulid/v2"
)

func AddCanteen(name string) (int64, error) {
	res, err := mariadb.Db.Exec(`
	INSERT INTO 
		canteen (name) 
	VALUES (?)`, name)

	if err != nil {
		return 0, err
	}

	inserted_id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return inserted_id, nil
}

func SelectCanteen() ([]models.Canteen, error) {
	var canteen_array = make([]models.Canteen, 0)
	rows, err := mariadb.Db.Query(`
	SELECT 
		canteen_id, 
		name 
	FROM 
		canteen`)

	if err != nil {
		return canteen_array, err
	}

	for rows.Next() {
		var canteen models.Canteen
		if err := rows.Scan(&canteen.Canteen_id, &canteen.Name); err != nil {
			return canteen_array, err
		}

		canteen_array = append(canteen_array, canteen)
	}

	if err := rows.Err(); err != nil {
		return canteen_array, err
	}

	return canteen_array, nil
}

func DeleteCanteen(id string) (int64, error) {
	res, err := mariadb.Db.Exec(`
	DELETE FROM 
		canteen 
	WHERE 
		canteen_id = ?`, id)

	if err != nil {
		return 0, err
	}

	rows_affected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows_affected, nil
}

func SelectOwnedCanteen(owid ulid.ULID) ([]models.Canteen, error) {
	var canteen_array = make([]models.Canteen, 0)

	rows, err := mariadb.Db.Query(`
	SELECT 
		c.*
	FROM 
		canteen c
	INNER JOIN
		canteen_ownership o
	ON
		c.canteen_id = o.canteen_id
	WHERE
		o.owner_id = ?
		`, owid)

	if err != nil {
		return canteen_array, err
	}

	for rows.Next() {
		var canteen models.Canteen
		if err := rows.Scan(&canteen.Canteen_id, &canteen.Name); err != nil {
			return canteen_array, err
		}

		canteen_array = append(canteen_array, canteen)
	}

	if err := rows.Err(); err != nil {
		return canteen_array, err
	}

	return canteen_array, nil
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
