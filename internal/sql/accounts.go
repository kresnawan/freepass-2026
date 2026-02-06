package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"strconv"
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

func InsertAccount(acc models.Account, tx *sql.Tx, uid ulid.ULID) error {
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

func DeactiveAccount(uid ulid.ULID) error {
	query := `
		UPDATE
			accounts
		SET
			is_active = 0, deactived_at = ?
		WHERE
			account_id = ?
	`

	_, err := mariadb.Db.Exec(query, time.Now().UTC(), uid)

	if err != nil {
		return err
	}

	return nil
}

func ReactivateAccount(uid ulid.ULID) error {
	query := `
		UPDATE
			accounts
		SET
			is_active = 1
		WHERE
			account_id = ?
	`

	_, err := mariadb.Db.Exec(query, uid)

	if err != nil {
		return err
	}

	return nil
}

func GetAccounts(page string) ([]models.Account, error) {
	var accounts = make([]models.Account, 0)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return accounts, err
	}

	query := `
		SELECT
			account_id,
			username,
			email,
			first_name,
			last_name,
			role,
			is_active,
			created_at,
			updated_at
		FROM
			accounts
		LIMIT
			10
		OFFSET ?
	`

	rows, err := mariadb.Db.Query(query, ((pageInt - 1) * 10))

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
			&account.Role,
			&account.IsActive,
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
