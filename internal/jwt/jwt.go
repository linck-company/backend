package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	USER_ID                       = "uid"
	TOKEN_ID                      = "tid"
	USER_NAME                     = "uname"
	USER_ROLE                     = "rl"
	EXPIRATION                    = "exp"
	JWT_VALUE_FETCH_ERROR_MESSAGE = "Error getting value"
)

type JWT struct {
	Secret         []byte
	RememberMeDays time.Duration
	ExpirationHour time.Duration
}

func (j *JWT) GenerateJWT(userId, userName, userRole, tokenId string, rememberMe bool) string {
	var expirationTime time.Time
	if rememberMe {
		expirationTime = time.Now().Add(j.RememberMeDays * 24 * time.Hour)
	} else {
		expirationTime = time.Now().Add(24 * time.Hour)
	}
	expirationTime = expirationTime.Truncate(24 * time.Hour).Add(j.ExpirationHour * time.Hour)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[USER_ID] = userId
	claims[USER_NAME] = userName
	claims[USER_ROLE] = userRole
	claims[EXPIRATION] = expirationTime.Unix()
	claims[TOKEN_ID] = tokenId

	tokenString, err := token.SignedString(j.Secret)
	if err != nil {
		fmt.Println("Error signing the token:", err)
		return ""
	}

	return tokenString
}

func (j *JWT) ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.Secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims[EXPIRATION].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token has expired")
			}
		} else {
			return nil, fmt.Errorf("expiration claim not found")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (j *JWT) IsJWTValid(c jwt.MapClaims) bool {
	if exp, ok := c[EXPIRATION].(float64); ok {
		return time.Now().Unix() < int64(exp)
	}
	return false
}

func (j *JWT) getJwtValue(c jwt.MapClaims, key string) interface{} {
	if val, ok := c[key]; ok {
		return val
	}
	return JWT_VALUE_FETCH_ERROR_MESSAGE
}

func (j *JWT) GetUserName(c jwt.MapClaims) string {
	val := j.getJwtValue(c, USER_NAME)
	if val == JWT_VALUE_FETCH_ERROR_MESSAGE {
		return ""
	}
	return val.(string)
}

func (j *JWT) GetUserId(c jwt.MapClaims) string {
	val := j.getJwtValue(c, USER_ID)
	if val == JWT_VALUE_FETCH_ERROR_MESSAGE {
		return ""
	}
	return val.(string)
}

func (j *JWT) GetExpirationEpoch(token string) int64 {
	c, err := j.ParseJWT(token)
	if err != nil {
		return 0
	}
	if exp, ok := c["exp"].(int64); ok {
		return exp
	}
	return 0
}
