package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
)

type CustomClaims struct {
	AccountId ulid.ULID `json:"account_id"`
	Role      string    `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(uid ulid.ULID, role string) (string, error) {
	var (
		key   []byte
		token *jwt.Token
	)

	claims := CustomClaims{
		AccountId: uid,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "bcc_canteen",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	key = []byte(os.Getenv("ACCESS_TOKEN_PK"))
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func GenerateRefreshToken(uid ulid.ULID, role string) (string, error) {
	var (
		key   []byte
		token *jwt.Token
	)

	claims := CustomClaims{
		AccountId: uid,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "kresnawan.com",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	key = []byte(os.Getenv("REFRESH_TOKEN_PK"))
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func VerifyAccessToken(tokenInput string, claims jwt.Claims) (*jwt.Token, int, error) {
	var secretKey []byte = []byte(os.Getenv("ACCESS_TOKEN_PK"))

	token, err := jwt.ParseWithClaims(tokenInput, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}

		return secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return token, 242, err
		}

		return token, 243, err
	}

	if !token.Valid {
		return token, 244, err
	}

	return token, 240, nil
}

func VerifyRefreshToken(tokenInput string, claims jwt.Claims) (*jwt.Token, int, error) {
	var secretKey []byte = []byte(os.Getenv("REFRESH_TOKEN_PK"))

	token, err := jwt.ParseWithClaims(tokenInput, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}

		return secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return token, 242, err
		}

		return token, 243, err
	}

	if !token.Valid {
		return token, 244, err
	}

	return token, 240, nil
}
