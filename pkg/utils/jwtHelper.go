package utils

import (
	"crypto/rsa"
	"github.com/astaxie/beego"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

var keyMap sync.Map

func GenerateJwtWithUserInfo(uid string, uname string) (string, error) {
	jwtToken := components.NewJwtHandler()
	jwt_private_key, err := filepath.Abs(components.Args.ConfigFile+ "/keys/jwt_private.pem")
	if err != nil {
		return "",err
	}
	jwtToken.SetPrivateKey(LoadRSAPrivateKeyFromDisk(jwt_private_key))
	defer jwtToken.Release()
	claims := components.JwtClaims{Uid: uid, Uname: uname}
	//token with 1 day expired
	claims.ExpiresAt = time.Now().Unix() + 3600*24*1
	return jwtToken.Generate(claims)
}
func GenerateRefreshJwtWithToken(token string) (string, error) {
	jwtToken := components.NewJwtHandler()
	jwt_private_key, err := filepath.Abs(components.Args.ConfigFile+ "/keys/jwt_private.pem")
	if err != nil {
		return "",err
	}
	jwtToken.SetPrivateKey(LoadRSAPrivateKeyFromDisk(jwt_private_key))
	defer jwtToken.Release()
	claims := components.JwtRefreshClaims{Token: token}
	//token with 1 day expired
	claims.ExpiresAt = time.Now().Unix() + 3600*24*3
	return jwtToken.GenerateRefreshToken(claims)
}
func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	if key, set := keyMap.Load(location); set {
		return key.(*rsa.PrivateKey)
	} else {
		keyData, e := ioutil.ReadFile(location)
		if e != nil {
			beego.Error(e.Error())
		}
		key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
		if e != nil {
			beego.Error(e.Error())
		}
		keyMap.Store(location, key)
		return key
	}
	return nil
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	if key, set := keyMap.Load(location); set {
		return key.(*rsa.PublicKey)
	} else {
		keyData, e := ioutil.ReadFile(location)
		if e != nil {
			beego.Error(e.Error())
		}
		key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
		if e != nil {
			beego.Error(e.Error())
		}
		keyMap.Store(location, key)
		return key
	}
	return nil
}
