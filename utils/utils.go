package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "abc123"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePasswords(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func GenerateJWT(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyJWT(tokenString string) (int64, error) {
	// parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		// check if the signing method is HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	// check if there was an error parsing the token
	if err != nil {
		return 0, errors.New("invalid token")
	}

	// check if the token is valid
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// get the id from the claims
	id := int64(claims["id"].(float64))

	// return the id and no error
	return id, nil
}
