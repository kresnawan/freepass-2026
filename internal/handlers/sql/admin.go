package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"

	"github.com/oklog/ulid/v2"
)

func InsertAdminProfile(acc models.Account) error {
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
			admin_profile (account_id)
		VALUES
			(?)
	`

	_, err = tx.Exec(query, uid)

	if err != nil {
		return err
	}

	return nil
}
