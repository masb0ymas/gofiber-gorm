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
	var APP_NAME = os.Getenv("APP_NAME")
	var expiresToken = os.Getenv("JWT_ACCESS_TOKEN_EXPIRED")
	var secretKey = os.Getenv("JWT_SECRET_ACCESS_TOKEN")

	var JWT_SIGNATURE_KEY = []byte(secretKey)
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

	expiresIn, err := strconv.Atoi(expiresToken) // expires in days

	var JWT_EXPIRES_TOKEN = time.Now().Add(time.Hour * 24 * time.Duration(expiresIn)).Unix()

	if err != nil {
		return "", nil
	}

	claims := jwt.MapClaims{
		"iss": APP_NAME,
		"exp": JWT_EXPIRES_TOKEN,
		"uid": user_id,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	return token.SignedString(JWT_SIGNATURE_KEY)
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
Token Valid
*/
func VerifyToken(c *fiber.Ctx) (jwt.MapClaims, error) {
	var secretKey = os.Getenv("JWT_SECRET_ACCESS_TOKEN")

	var JWT_SIGNATURE_KEY = []byte(secretKey)
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid: %v", token.Header["alg"])
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
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
