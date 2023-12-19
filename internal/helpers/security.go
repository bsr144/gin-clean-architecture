package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"dbo-be-task/internal/config"
	"encoding/hex"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SecurityHelper struct {
	Config *config.SecurityConfig
}

func NewSecurityHelper(config *config.SecurityConfig) *SecurityHelper {
	return &SecurityHelper{
		Config: config,
	}
}

func (s *SecurityHelper) GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func (s *SecurityHelper) HashPassword(password, salt string) string {
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(salt))
	sha256Hasher.Write([]byte(password))
	return hex.EncodeToString(sha256Hasher.Sum(nil))
}

func (s *SecurityHelper) ComparePassword(storedHash, storedSalt, providedPassword string) bool {
	hash := s.HashPassword(providedPassword, storedSalt)

	return hash == storedHash
}

func (s *SecurityHelper) GenerateToken(userID int) (string, error) {
	jwtDuration := time.Duration(s.Config.JWTExpiryTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * jwtDuration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.Config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
