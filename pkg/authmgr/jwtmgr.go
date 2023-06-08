package authmgr

import (
	"context"
	"encoding/json"
	"evildoer/conf"
	"evildoer/pkg/cryptic"
	"evildoer/pkg/jwtx"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct{}

func NewJWTManger() *JWTManager {
	jwtx.SetDefaultIssuer(jwtx.NewIssuer(jwt.SigningMethodHS256, time.Second, conf.GetConf().Secret))
	return &JWTManager{}
}

type Credentials struct {
	Username string `json:"username"`
}

func (c *Credentials) Claims(secret cryptic.Secret) (string, error) {
	values, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return cryptic.AESEncrypt(string(values), secret)
}

func (c *Credentials) Values(values string, secret cryptic.Secret) (err error) {
	values, err = cryptic.AESDecrypt(values, secret)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(values), c)
}

func (m *JWTManager) IsAllowedUser(username, password string) (string, error) {
	if !isAllowedUser(username, password) {
		panic("todo: invalid username and password")
	}
	return m.GenerateToken(username)
}

func (m *JWTManager) GenerateToken(username string) (string, error) {
	return jwtx.Signin(&Credentials{Username: username}, int(time.Hour/time.Second))
}

func (m *JWTManager) GetCredentials(tknStr string) (*Credentials, error) {
	credentials := new(Credentials)
	if err := jwtx.Welcome(tknStr, credentials); err != nil {
		return nil, err
	}
	return credentials, nil
}

type authInfoCtxKey int

type authInfo struct {
	username string
}

func (m *JWTManager) SetCtxInfo(ctx context.Context, credentials *Credentials) context.Context {
	return context.WithValue(ctx, authInfoCtxKey(0), &authInfo{username: credentials.Username})
}

func (m *JWTManager) GetUsernameFromCtx(ctx context.Context) string {
	return ctx.Value(authInfoCtxKey(0)).(*authInfo).username
}
