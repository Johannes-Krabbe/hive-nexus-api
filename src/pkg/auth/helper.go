package auth

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/common/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/scrypt"
)

func hash(p string, s string) (hp string, err error) {
	dk, err := scrypt.Key([]byte(p), []byte(s), 32768, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	hp = hex.EncodeToString(dk)
	return
}

func getSalt() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// Generate a JWT token for a user with the given ID and return the token string
func generateJWT(userID uuid.UUID) (string, error) {
	var secret = []byte(config.GetValueFromEnv("JWT_SECRET"))
	// Create the claims object with the user ID and expiration time
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create the token with the claims and sign with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify a JWT token and return the user ID if valid
func verifyJWT(tokenString string) (uuid.UUID, error) {
	var secret = []byte(config.GetValueFromEnv("JWT_SECRET"))
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	// Check if the token is valid and get the user ID from the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDstring, ok := claims["user_id"].(string)
		// TODO dont discard the error
		userID, _ := uuid.Parse(userIDstring)
		fmt.Println(claims)
		if !ok {
			return uuid.Nil, fmt.Errorf("invalid user id")
		}
		return userID, nil
	} else {
		return uuid.Nil, fmt.Errorf("invalid token")
	}
}
