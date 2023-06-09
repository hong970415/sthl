package authentication

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	UserId string
}

func GenerateJwtToken(duration time.Duration, userId string) (string, error) {
	if userId == "" {
		return "", fmt.Errorf("userId is empty")
	}
	if duration <= 0 {
		return "", fmt.Errorf("duration cannot smaller than 0")
	}
	// Get jwt secret key
	// signingKey := []byte(utils.GetEnv("JWT_SECRET"))
	signingKey := []byte("secret")
	claims := &JwtCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userId,
	}
	// Sign jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func VerifyJwtToken(tokenString string) (*JwtCustomClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("tokenString is empty")
	}

	var customClaims JwtCustomClaims
	// initialize a new JWT parser
	parser := jwt.Parser{
		ValidMethods: []string{jwt.SigningMethodHS256.Alg()},
	}

	// parse the token string using the secret key
	token, err := parser.ParseWithClaims(tokenString, &customClaims, func(token *jwt.Token) (interface{}, error) {
		// check if the signing method is the expected one
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//return the secret key
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	// check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// check if the token is not expired
	if time.Now().After(customClaims.ExpiresAt.Time) {
		return nil, fmt.Errorf("token expired")
	}

	return &customClaims, nil
}
