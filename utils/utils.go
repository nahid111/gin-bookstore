package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateAuthToken(email string, userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error generating token:", err)
		return ""
	}

	return tokenString
}

func ValidateAuthToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method (e.g., HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(err)
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid JWT token")
}

// Hashes the given string.
// Returns a pointer to string or nil
func HashPassword(password string) *string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing password.", err)
		return nil
	}
	hashPassStr := string(hashedPass)
	return &hashPassStr
}

func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
