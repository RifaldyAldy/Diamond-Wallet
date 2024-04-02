package common

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtClaim struct {
	jwt.StandardClaims
	DataClaims model.JwtClaims `json:"data"`
}

var (
	appName          = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("TOKEN_KEY"))
)

func GenerateTokenJwt(id, name, role string, expiredAt int64) (string, error) {
	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: expiredAt, // expayet waktu login
		},
		DataClaims: model.JwtClaims{
			Id:   id,
			Name: name,
			Role: role,
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
			SendErrorResponse(c, http.StatusForbidden, "Invalid Token")
			c.Abort()
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
			c.Abort()
			return
		}

		if !token.Valid {
			SendErrorResponse(c, http.StatusUnauthorized, "Unaunthorized user")
			c.Abort()
			return
		}

		expiredAt := claims.ExpiresAt
		if time.Now().Unix() > expiredAt {
			SendErrorResponse(c, http.StatusUnauthorized, "Expired Token")
			c.Abort()
			return
		}

		// validation role

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.DataClaims.Role {
					validRole = true
					break
				}
			}
		}
		if !validRole {
			SendErrorResponse(c, http.StatusForbidden, "You dont have permission")
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
