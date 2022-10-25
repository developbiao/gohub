package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// JWT define jwt objet
type JWT struct {
	// Secret key read from app.key
	SignKey []byte

	// Refresh token max expire time
	MaxRefresh time.Duration
}

type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims struct implement Claims interface
	// - iss (issuer): 发布者
	// - sub (subject): 主题
	// - iat (Issued At): 生成签名的时间
	// - exp (expiration time): 签名过期时间
	// - aud (audience): 观众，相当于接受者
	// - nbf (Not Before): 生效时间
	// - jti (JWT ID): 编号
	jwtpkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	// 1. Invoke jwt library paring user token
	token, err := jwt.parseTokenString(tokenString)

	// 2. Paring error information
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 3. Paring claims information from JWTCustomClaims
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// getTokenFromHeader using jwtpkg.ParseWithClaims paring token
// Authorization: Bearer
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// Split with space
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// RefreshToken update token refresh token api
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 1. Get token from header
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. Invoke jwt paring user token
	token, err := jwt.parseTokenString(tokenString)

	// 3. Paring error, not error is valid token
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", errr
		}
	}

	// 4. Paring JWTCustomClaims data
	claims := token.Claims.(*JWTCustomClaims)

	// 5. Check JWTCustomClaims data
	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		// Change expire time
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}
	return "", ErrTokenExpiredMaxRefresh
}
