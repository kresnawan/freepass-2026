package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/oklog/ulid/v2"
)

func GetUsernameEmailExistence(username string, email string) (bool, error) {
	var foundUsername, foundEmail string

	query := `
		SELECT 
			username,
			email
		FROM 
			accounts
		WHERE
			username = ? OR email = ?
		`

	err := mariadb.Db.QueryRow(query, username, email).Scan(&foundUsername, &foundEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		} else {
			return false, err
		}
	}

	return false, nil
}

func InsertAccount(
	acc models.Account,
	tx *sql.Tx,
	uid ulid.ULID,
) error {
	exist, err := GetUsernameEmailExistence(acc.Username, acc.Email)

	if err != nil {
		return err
	}

	hashedPasswd, _ := argon2id.CreateHash(acc.Password, argon2id.DefaultParams)

	if exist {
		query := `
			INSERT INTO
				accounts (account_id,username,email,first_name,last_name,password,role)
			VALUES
				(?,?,?,?,?,?,?)
		`

		_, err = tx.Exec(query, uid, acc.Username, acc.Email, acc.FirstName, acc.LastName, hashedPasswd, acc.Role)

		if err != nil {
			return err
		}

		return nil
	} else {
		return sql.ErrNoRows
	}
}

func DisableAccount(uid string, role string) error {
	parsedId, _ := ulid.Parse(uid)

	query := `
		UPDATE
			accounts
		SET
			is_active = 0, deactived_at = ?
		WHERE
			account_id = ? AND role = ?
	`

	_, err := mariadb.Db.Exec(query, time.Now(), parsedId, role)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(uid string) error {
	parsedId, _ := ulid.Parse(uid)
	tx, err := mariadb.Db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	deleteProfile := `
		DELETE FROM
			customer_profile
		WHERE
			account_id = ?
	`

	deleteAccount := `
		DELETE FROM
			accounts
		WHERE
			account_id = ?

	`

	_, err = tx.Exec(deleteProfile, parsedId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteAccount, parsedId)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func GetAccounts() ([]models.Account, error) {
	var accounts = make([]models.Account, 0)

	query := `
		SELECT
			*
		FROM
			accounts
	`

	rows, err := mariadb.Db.Query(query)

	if err != nil {
		return accounts, err
	}

	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&account.AccountId,
			&account.Username,
			&account.Email,
			&account.FirstName,
			&account.LastName,
			&account.Password,
			&account.Role,
			&account.IsActive,
			&account.DeactivedAt,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return accounts, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
