package common

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaim struct {
	jwt.StandardClaims
	UserData model.UserData `json:"user"`
}

var (
	appName          = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("SIGNATURE_KEY"))
)

func GenerateTokenJwt(userData model.User, expiredAt int64) (string, error) {
	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: expiredAt,
		},
		UserData: model.UserData{
			Id:          userData.Id,
			Name:        userData.Name,
			Email:       userData.Email,
			Role:        userData.Role,
			PhoneNumber: userData.PhoneNumber,
		},
	}
	token := jwt.NewWithClaims(jwtSigningMethod, claims)
	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			SendErrorResponse(c, http.StatusForbidden, "Invalid token")
			return
		}

		// jwtSignatureKey := []byte(os.Getenv("SIGNATURE_KEY"))
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSignatureKey, nil
		})
		if err != nil {
			SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !token.Valid {
			SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized user")
			return
		}

		expiredAt := claims.ExpiresAt
		if time.Now().Unix() > expiredAt {
			SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized user")
			return
		}

		//  Validasi Role
		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.UserData.Role {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			SendErrorResponse(c, http.StatusForbidden, "Forbidden user")
			return
		}

		c.Next()
	}
}
