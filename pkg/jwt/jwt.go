package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/logger"
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

// JWTCustomClaims custom payload
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

// ParserToken paring token from header
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

// parseTokenString using jwtpkg.ParseWithClaims paring token
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
			return "", err
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

// IssueToken generate token for login invoke
func (jwt *JWT) IssueToken(userID string, userName string) string {
	// 1. Construct user claims information payload
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime,
		jwtpkg.StandardClaims{
			NotBefore: app.TimenowInTimezone().Unix(), // sign time
			IssuedAt:  app.TimenowInTimezone().Unix(), // first sign time
			ExpiresAt: expireAtTime,                   // sign expire time
			Issuer:    config.GetString("app.name"),   // Issuer
		},
	}

	// 2. Generate token object by claims
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// createToken create token by custom claims
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// Using HS256 algorithm method generate token
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime get expire at time
func (jwt *JWT) expireAtTime() int64 {
	timenow := app.TimenowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}
