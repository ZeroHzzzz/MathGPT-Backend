package midwares

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"mathgpt/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func generateJwtSecret() string {
	// 生成一个32字节的随机密钥
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}

var jwtSecret = generateJwtSecret()

func CreateJWT(userID uint, duration time.Duration) (string, error) {
	claims := models.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "MathGPT",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("not logged in")
}

func RefreshJWT(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	// 解析Token
	claims, err := ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 检查Token是否接近过期（例如10分钟内）
	if time.Until(claims.ExpiresAt.Time) > time.Minute*10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not close to expiration"})
		return
	}

	// 生成新的Token
	newToken, err := CreateJWT(claims.UserID, time.Hour*2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// 解析Token
		claims, err := ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 将解析后的数据存储到上下文中，供后续使用
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
