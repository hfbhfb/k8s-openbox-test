package jwtwrap

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"backend/codeasset/common/ginutils"
	. "backend/codeasset/common/share"

	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")

	TokenHoursLast int64  = 30
	defaultSignKey string = "4be63eb0aa8be9d80a95820e113fa716"
)

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(defaultSignKey),
	}
}

func (j *JWT) SetSignKey(signKey string) {
	j.SigningKey = []byte(defaultSignKey)
}

func (j *JWT) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// accesskeyID := c.Request.Header.Get("X-AccessKeyID")
		token := c.Request.Header.Get("access_token")
		if token == "" {
			ginutils.WriteString(c, 501, ErrorText(501), nil)
			c.Abort()
			return
		}
		// yclog.Debug("accesskey_id:", accesskeyID, "token:", token)

		j := NewJWT()
		// 如果accesskey不为空
		// if accesskeyID != "" {
		// 	acc := models.NewAccessKey(models.BaseData)
		// 	if err := acc.GetAccessKeyByAccessKeyID(accesskeyID); err != nil {
		// 		ginutils.WriteString(c, 504, "accesskeyID不存在", nil)
		// 		c.Abort()
		// 		return
		// 	}
		// 	j.SetSignKey(acc.AccessKeySecret)
		// }
		// 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				ginutils.WriteString(c, 502, ErrorText(502), nil)
				c.Abort()
				return
			}
			ginutils.WriteString(c, 503, ErrorText(503), nil)
			c.Abort()
			return
		}
		c.Set("X-Claims", claims)
	}
}

func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(TokenHoursLast) * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func GetClient(c *gin.Context) *Claims {
	ci, b := c.Get("X-Claims")
	if !b {
		return nil
	}
	fmt.Println("%#v", ci)
	if v, b := ci.(*Claims); b {
		return v
	}
	return nil
}

func GetToken(c *gin.Context) string {
	return c.Request.Header.Get("access_token")
}
