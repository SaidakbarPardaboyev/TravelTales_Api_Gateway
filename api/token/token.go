package token

import (
	"fmt"
	"travel/config"

	jwt "github.com/golang-jwt/jwt"
)

func ValidateToken(tokenstr string) (bool, error) {
	if _, err := ExtractClaims(tokenstr); err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenstr string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(
		t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				t.Header["alg"])
		}
		return []byte(config.Load().SINGNING_KEY_ACCESS), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error: cannot parse token: %s", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token: %s", err)
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims: %s", err)
	}
	return claims, nil
}
