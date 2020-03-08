package jwthelper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretJWT = "ASDJADHASDKFHADS"

// Create the JWT key used to create the signature
var jwtKey = []byte(secretJWT)

// CreateJWT ...
func CreateJWT(userID string, duration int) (string, bool) {

	// expation time
	expirationTime := time.Now().Add(time.Duration(duration) * time.Minute)

	//claims é o que vai ser encodado no jwt
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Id:        userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//convert to string

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", false
	}

	return tokenString, true
}

// DecodeJWT ...
func DecodeJWT(stringJWT string) (string, bool) {
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(stringJWT, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", false
	}

	if !token.Valid {
		return "", false
	}
	// fmt.Println("OOO")

	return claims.Id, true

}
