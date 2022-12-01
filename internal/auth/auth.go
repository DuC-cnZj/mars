package auth

import (
	"crypto/rsa"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/models"
	"strings"
	"time"

	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/golang-jwt/jwt"
)

type Authn struct {
	authns   []contracts.Authenticator
	signFunc func(info contracts.UserInfo) (*contracts.SignData, error)
}

func NewAuthn(signFunc func(info contracts.UserInfo) (*contracts.SignData, error), authns ...contracts.Authenticator) *Authn {
	return &Authn{authns: authns, signFunc: signFunc}
}

func (a *Authn) VerifyToken(s string) (*contracts.JwtClaims, bool) {
	for _, authn := range a.authns {
		if token, ok := authn.VerifyToken(s); ok {
			return token, true
		}
	}

	return nil, false
}

func (a *Authn) Sign(info contracts.UserInfo) (*contracts.SignData, error) {
	return a.signFunc(info)
}

type Auth struct {
	priKey *rsa.PrivateKey
	pubKey *rsa.PublicKey
}

func NewJwtAuth(priKey *rsa.PrivateKey, pubKey *rsa.PublicKey) *Auth {
	return &Auth{priKey: priKey, pubKey: pubKey}
}

func (a *Auth) VerifyToken(t string) (*contracts.JwtClaims, bool) {
	var token string = t
	if len(t) > 6 && strings.EqualFold("bearer", t[0:6]) {
		token = strings.TrimSpace(t[6:])
	}
	if token != "" {
		parse, err := jwt.ParseWithClaims(token, &contracts.JwtClaims{}, func(token *jwt.Token) (any, error) {
			return a.pubKey, nil
		})
		if err == nil && parse.Valid {
			return parse.Claims.(*contracts.JwtClaims), true
		}
	}

	return nil, false
}

func (a *Auth) Sign(info contracts.UserInfo) (*contracts.SignData, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &contracts.JwtClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(contracts.Expired).Unix(),
			Issuer:    "mars",
			IssuedAt:  time.Now().Unix(),
			Subject:   info.Email,
		},
		UserInfo: info,
	})

	signedString, err := token.SignedString(a.priKey)
	if err != nil {
		return nil, err
	}
	return &contracts.SignData{
		Token:     signedString,
		ExpiredIn: int64(contracts.Expired.Seconds()),
	}, nil
}

type AccessTokenAuth struct {
	app contracts.ApplicationInterface
}

func NewAccessTokenAuth(app contracts.ApplicationInterface) *AccessTokenAuth {
	return &AccessTokenAuth{app: app}
}

func (a *AccessTokenAuth) VerifyToken(t string) (*contracts.JwtClaims, bool) {
	var token string = t
	if len(t) > 6 && strings.EqualFold("bearer", t[0:6]) {
		token = strings.TrimSpace(t[6:])
	}
	if token != "" {
		var at models.AccessToken
		if err := app.DB().Where("`token` = ?", token).First(&at).Error; err == nil {
			app.DB().Model(&at).Update("last_used_at", time.Now())
			return &contracts.JwtClaims{UserInfo: at.GetUserInfo()}, true
		}
	}

	return nil, false
}
