package httpx

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/keington/go-templet/pkg/cfg"
	"github.com/keington/go-templet/pkg/zlog"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/14 22:40
 * @file: jwt.go
 * @description:
 */

type AuthClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (a *AuthClaims) Valid() error {
	return nil
}

var (
	accessExpired  = cfg.GetDuration("http.auth.access_expired")
	refreshExpired = cfg.GetDuration("http.auth.refresh_expired")
	secretKey      = []byte(cfg.GetString("http.auth.secret_key"))
	issUser        = "lark"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	return secretKey, nil
}

// GenToken 生成 access_token 和 refresh_token
func GenToken(userId, UserName string) (aToken, rToken string, err error) {

	// aToken
	aClaims := &AuthClaims{
		Id:   userId,
		Name: UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issUser, // 签发人
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpired)),
		},
	}
	aToken, aErr := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims).SignedString(secretKey)
	if err != nil {
		zlog.Errorf("jwt.NewWithClaims err: %v", aErr)
		return "", "", aErr
	}

	// rToken
	rClaims := jwt.RegisteredClaims{
		Issuer:    issUser,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpired)),
	}
	rToken, rErr := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims).SignedString(secretKey)
	if err != nil {
		zlog.Debugf("jwt.NewWithClaims err: %v", rErr)
		return "", "", rErr
	}

	return aToken, rToken, nil
}

// VerifyToken 校验 access_token
func VerifyToken(aToken string) (claims *AuthClaims, err error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(aToken, &AuthClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新 access_token
// 第一步 : 判断 rToken 格式对的，没有过期的
// 第二步 : 判断 aToken 格式对的，但是是过期的
// 第三步 : 生成双 token
func RefreshToken(aToken, rToken string) (newAccessToken, newRefreshToken string, err error) {
	// 第一步 : 判断 rToken 格式对的，没有过期的
	if _, err := jwt.Parse(rToken, keyFunc); err != nil {
		return "", "", err
	}

	// todo 第二步：从旧的 aToken 中解析出 claims 数据   过期了还能解析出来吗
	var claims AuthClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	var v *jwt.ValidationError
	errors.As(err, &v)

	// 当 access token 是过期错误，并且 refresh token 没有过期就创建一个新的 access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.Id, claims.Name)
	}
	return "", "", err
}
