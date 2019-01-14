package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("GFDHTR&$%$%^#$GGFG5545gdgdf-90345")

// JWT 参考：http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
// 自定义 Payload（负载）
type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// 创建token
func GenerateToken(nickname string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24 * 3)
	claims := Claims{
		nickname,
		jwt.StandardClaims{
			Issuer:    "com.cinema",      // 签发人
			ExpiresAt: expireTime.Unix(), // 过期时间
			Subject:   nickname,          // 使用人
			Audience:  "all",             // 接收对象
			NotBefore: nowTime.Unix(),    // 生效时间
			IssuedAt:  nowTime.Unix(),    //签发时间
			//Id:        nil,               // 唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析 token
func ParseToken(token string) (*Claims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 返回key
func keyFunc(token *jwt.Token) (interface{}, error) {
	return jwtSecret, nil
}
