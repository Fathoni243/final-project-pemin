package middleware

import (
	"final-project-pemin/util"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(nim string) (string, error) {
	expTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_SEC"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nim": nim,
		"exp": time.Now().Add(time.Second * time.Duration(expTime)).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("token")
		if bearerToken == "" {
			util.FailOrErrorResponse(c, http.StatusBadRequest, "token not found", nil)
			c.Abort()
			return
		}

		// bearerToken = bearerToken[7:]
		tokenExtract, err := jwt.Parse(bearerToken, tokenExtract)
		if err != nil {
			util.FailOrErrorResponse(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}
		if claims, ok := tokenExtract.Claims.(jwt.MapClaims); ok && tokenExtract.Valid {
			nim := claims["nim"].(string)
			c.Set("nim", nim)
			c.Next()
			return
		}
		util.FailOrErrorResponse(c, http.StatusForbidden, "invalid token", nil)
		c.Abort()
	}
}

func tokenExtract(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
