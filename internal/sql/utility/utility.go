package utility

import "strings"

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
