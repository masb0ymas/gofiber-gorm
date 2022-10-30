package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var secretKey = os.Getenv("JWT_SECRET_ACCESS_TOKEN")
var expiresToken = os.Getenv("JWT_ACCESS_TOKEN_EXPIRED") // example: 1 | 7 | 14 | 30 days

// generate token
func GenerateToken(user_id uuid.UUID) (string, error) {
	expiresIn, err := strconv.Atoi(expiresToken) // expires in days

	if err != nil {
		return "", nil
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(expiresIn)).Unix() // Token expires after 7 Days

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// extract token
func ExtractToken(c *fiber.Ctx) string {
	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

// extract token id
func ExtractTokenID(c *fiber.Ctx) (uint, error) {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["uid"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}

func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}

	return nil
}