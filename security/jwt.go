package security

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct{}

var jwtKey = []byte("secretadawdawdawdawdaw")

type Claims struct {
	Username string `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func (j JWT) GenerateToken(username string) (string, error) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		Roles:    []string{"adminngu"},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWT) ParseToken(tokenString string) (*Claims, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	

	return claims, nil
}

func (j JWT) AuthRoles(roles []string) error { 
	hasRole := false

	for _, role := range roles {
		if role == "admin" {
			hasRole = true
		}
	}

	if !hasRole {
		return errors.New("invalid role")
	}
	return nil
}
