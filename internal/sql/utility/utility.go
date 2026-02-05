package utility

import (
	"canteen/internal/storage/mariadb"
	"database/sql"
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
)

func PrepareMultipleInsert(canteenID []byte, userIDs []int) (string, []any) {
	var placeholders []string
	var args []any

	for _, id := range userIDs {
		placeholders = append(placeholders, "(?, ?)")
		args = append(args, canteenID, id)
	}

	queryValues := strings.Join(placeholders, ", ")

	return queryValues, args
}

func GeneratePlaceholders(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat("?,", n-1) + "?"
}

func GetUnusedId(column string, table string) (ulid.ULID, error) {
	var uid ulid.ULID = ulid.Make()
	var tempId ulid.ULID

	query := fmt.Sprintf("SELECT %s FROM `%s` WHERE %s = ?", column, table, column)

	for {
		err := mariadb.Db.QueryRow(query, uid).Scan(&tempId)
		if err != nil {
			if err == sql.ErrNoRows {
				return uid, nil
			} else {
				return tempId, err
			}
		}
	}

}
