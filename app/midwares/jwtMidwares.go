package midwares

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"mathgpt/app/apiException"
	"mathgpt/app/models"
	"mathgpt/app/utils"
	"mathgpt/configs/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Duration = config.Config.GetDuration("jwt.duration")
var jwtSecret = generateJwtSecret()

func generateJwtSecret() string {
	// 生成一个32字节的随机密钥
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}

func CreateJWT(userID string) (string, error) {
	claims := models.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(Duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "MathGPT",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ParseJWT(tokenString string) (*models.Claims, error) {
	// Remove "Bearer " prefix if it exists
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Invalid token signing method: %v", token.Header["alg"])
			return nil, apiException.NotLogin
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, apiException.NotLogin
	}

	return claims, nil
}

func RefreshJWT(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, apiException.NotLogin)
		return
	}

	// 解析Token
	claims, err := ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apiException.NotLogin)
		return
	}

	// 检查Token是否接近过期（例如10分钟内）
	if time.Until(claims.ExpiresAt.Time) > time.Minute*10 {
		c.JSON(http.StatusUnauthorized, apiException.NotLogin)
		return
	}

	// 生成新的Token
	newToken, err := CreateJWT(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"token":     newToken,
		"expiresIn": Duration,
	})
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, apiException.NotLogin)
			c.Abort()
			return
		}

		// 解析Token
		claims, err := ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, apiException.NotLogin)
			c.Abort()
			return
		}

		// 将解析后的数据存储到上下文中，供后续使用
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
