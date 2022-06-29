package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken generates a new token with user permissions
func CreateToken(userId uint64) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	return token.SignedString([]byte(config.SecretKey))
}

// Validate token verify if token is valid
func ValidateToken(r *http.Request) error {
	stringToken := getToken(r)
	token, erro := jwt.Parse(stringToken, returnVerificationKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

// Returns user id from token
func GetUserId(r *http.Request) (uint64, error) {
	stringToken := getToken(r)
	token, erro := jwt.Parse(stringToken, returnVerificationKey)
	if erro != nil {
		return 0, erro
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permission["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userId, nil
	}

	return 0, errors.New("Invalid token")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected sign method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
