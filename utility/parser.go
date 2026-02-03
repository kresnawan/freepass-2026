package utility

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

func AnyToUlid(v any) (ulid.ULID, error) {
	var parsedId ulid.ULID
	s, ok := v.(string)

	if !ok {
		return parsedId, errors.New("Type of 'any' failed to be parsed to string")
	}

	parsedId, err := ulid.Parse(s)
	if err != nil {
		return parsedId, err
	}

	return parsedId, nil
}
