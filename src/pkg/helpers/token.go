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

/*
Generate Token
*/
func GenerateToken(user_id uuid.UUID) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_ACCESS_TOKEN")
	expiresToken := os.Getenv("JWT_ACCESS_TOKEN_EXPIRED") // example: 1 | 7 | 14 | 30 days

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

/*
Extract Token
*/
func ExtractToken(c *fiber.Ctx) string {
	// get token by query
	getTokenByQuery := c.Query("token")

	// get token by cookies
	getTokenByCookie := c.Cookies("token")

	// get token by request header
	getTokenByReqHeader := c.Get("Authorization")

	if getTokenByQuery != "" {
		logMessage := PrintLog("Auth", "Extract from Query")
		fmt.Println(logMessage, getTokenByQuery)

		return getTokenByQuery
	}

	if getTokenByCookie != "" {
		logMessage := PrintLog("Auth", "Extract from Cookie")
		fmt.Println(logMessage, getTokenByCookie)

		return getTokenByCookie
	}

	if len(strings.Split(getTokenByReqHeader, " ")) == 2 {
		logMessage := PrintLog("Auth", "Extract from Header")
		fmt.Println(logMessage, getTokenByReqHeader)

		return strings.Split(getTokenByReqHeader, " ")[1]
	}

	return ""
}

/*
Extract Token ID
*/
func ExtractTokenID(c *fiber.Ctx) (uint, error) {
	secretKey := os.Getenv("JWT_SECRET_ACCESS_TOKEN")
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

/*
Token Valid
*/
func TokenValid(c *fiber.Ctx) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_ACCESS_TOKEN")
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		PrettyJSON(claims)

		return claims, err
	}

	return nil, err
}
