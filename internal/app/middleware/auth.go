package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"music-pc-server/internal/app/plus"
	"time"
)

/*
	校验Token
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			plus.RespError(context, plus.MSC_NotToken)
			context.Abort()
			return
		}
		//解析Token
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				plus.RespError(context, plus.MSC_TokenExpired)
				context.Abort()
				return
			}
			e := plus.UnWrapResponse(plus.MSC_InvalidToken)
			plus.RespError(context, plus.WrapResponse(err, e.Code, e.Message, e.StatusCode))
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
	}
}

/*
	定义错误
*/
var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

/*
	JWT签名结构
*/
type JWT struct {
	SigningKey []byte
}

/*
	JWT实例
*/
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

/*
	获取sign
*/
func GetSignKey() string {
	return "请填写加密签名"
}

/*
	创建一个token
*/
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}

/*
	解析Token
*/
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

/*
	更新Token
*/
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(12 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
