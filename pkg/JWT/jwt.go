package JWT

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"gozakupki-api/domain"
	"time"
)

var (
	secretKey = []byte("TA9[ZPw=rt&4&f1v/zl%gXn't53}d!3") // Replace with your actual secret key
)

func GenerateToken(id int64, login, email string) (string, error) {
	// Define the expiration time (15 days from now)
	expirationTime := time.Now().Add(15 * 24 * time.Hour)

	// Create the claims with user information and expiration time
	claims := jwt.MapClaims{
		"userID": id,
		"login":  login,
		"email":  email,
		"exp":    expirationTime.Unix(),
	}

	// Create a new JWT token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValid(t string) error {
	parsedToken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	fmt.Println(parsedToken.Claims)
	// Check for parsing errors
	if err != nil {
		return domain.ErrInternalServerError
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		return domain.ErrUnauthorized
	}

	return nil
}
