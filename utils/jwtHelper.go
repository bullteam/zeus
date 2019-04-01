package utils

import (
	"time"
	"github.com/bullteam/zeus/components"
	"github.com/astaxie/beego"
	"crypto/rsa"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"sync"
)
var keyMap sync.Map
func GenerateJwtWithUserInfo(uid string,uname string) (string,error){
	jwtToken := components.NewJwtHandler()
	jwtToken.SetPrivateKey(LoadRSAPrivateKeyFromDisk(beego.AppPath + "/keys/jwt_private.pem"))
	defer jwtToken.Release()
	claims := components.JwtClaims{Uid:uid,Uname:uname}
	//token with 1 day expired
	claims.ExpiresAt = time.Now().Unix() + 3600 * 24 * 1
	return jwtToken.Generate(claims)
}
func GenerateRefreshJwtWithToken(token string) (string,error){
	jwtToken := components.NewJwtHandler()
	jwtToken.SetPrivateKey(LoadRSAPrivateKeyFromDisk(beego.AppPath + "/keys/jwt_private.pem"))
	defer jwtToken.Release()
	claims := components.JwtRefreshClaims{Token:token}
	//token with 1 day expired
	claims.ExpiresAt = time.Now().Unix() + 3600 * 24 * 3
	return jwtToken.GenerateRefreshToken(claims)
}
func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	if key,set := keyMap.Load(location);set{
		return key.(*rsa.PrivateKey)
	}else {
		keyData, e := ioutil.ReadFile(location)
		if e != nil {
			beego.Error(e.Error())
		}
		key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
		if e != nil {
			beego.Error(e.Error())
		}
		keyMap.Store(location,key)
		return key
	}
	return nil
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	if key,set := keyMap.Load(location);set{
		return key.(*rsa.PublicKey)
	}else {
		keyData, e := ioutil.ReadFile(location)
		if e != nil {
			beego.Error(e.Error())
		}
		key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
		if e != nil {
			beego.Error(e.Error())
		}
		keyMap.Store(location,key)
		return key
	}
	return nil
}