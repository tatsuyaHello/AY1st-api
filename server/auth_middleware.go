package server

import (
	"AY1st/infra"
	"fmt"
	"net/http"
	"strings"

	"AY1st/util"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 認証したJWTのチェックを行う
func AuthMiddleware(authenticator Authenticator) gin.HandlerFunc {
	// この部分はサーバー起動時に1度だけ実行される
	logger := util.GetLogger()
	// この部分はクライアント接続毎に実行される
	return func(c *gin.Context) {
		err := AuthHandler(c, authenticator)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		}
		c.Next()
	}
}

// OptionalAuthMiddleware は認証情報がある場合だけJWTのチェックを行う
func OptionalAuthMiddleware(authenticator Authenticator) gin.HandlerFunc {
	logger := util.GetLogger()
	return func(c *gin.Context) {
		if _, ok := GetBearer(c.Request.Header["Authorization"]); ok {
			err := AuthHandler(c, authenticator)
			if err != nil {
				logger.Error(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			}
		} else {
			c.Next()
		}
	}
}

// AuthHandler はAuthMiddlewareのメイン処理です
func AuthHandler(c *gin.Context, authenticator Authenticator) error {
	util.GetLogger().Info(fmt.Sprintf("Header: %+v", c.Request.Header))

	tokenString, ok := GetBearer(c.Request.Header["Authorization"])

	if !ok {
		return fmt.Errorf("Authorization Bearer Header is missing")
	}

	authedUser, err := AuthenticateUser(tokenString, authenticator)
	if err != nil {
		return err
	}

	// 認証したtokenを渡す
	// c.Set("id", authedUser.Claims.(*jwt.StandardClaims).Id)
	c.Set("email", authedUser.Email)
	c.Set("sub", authedUser.Sub)
	c.Set("token", authedUser.Token)
	c.Set("userName", authedUser.Username)

	return nil
}

// AuthenticatedUser は認証したユーザー情報です
type AuthenticatedUser struct {
	Email    string
	Username string
	Sub      string
	Token    *jwt.Token
}

// AuthenticateUser はtokenを検証した結果を返す
func AuthenticateUser(tokenString string, authenticator Authenticator) (*AuthenticatedUser, error) {
	token, err := authenticator.ValidateToken(tokenString)
	// ローカル環境はjwtのバリデートエラーを検証しない
	if !infra.IsLocal() {
		if err != nil {
			// jwtの検証に失敗
			return nil, fmt.Errorf("token is not valid. [Reason] %v", err)
		}
	}

	// developモードでtokenの検証をスキップした場合、token == nilがあり得る
	if token == nil || token.Claims == nil {
		return nil, fmt.Errorf("wrong format token")
	}

	// よく使うユーザー情報も渡す
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("wrong format token")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("token must contain email")
	}
	sub := claims["sub"].(string)
	username := claims["cognito:username"].(string)

	authedUser := AuthenticatedUser{
		Email:    email,
		Username: username,
		Sub:      sub,
		Token:    token,
	}

	return &authedUser, nil
}

// GetBearer は AuthorizationヘッダからBearerトークンを取得します
func GetBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) == 2 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}
