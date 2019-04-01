package components

import (
	"sync"
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego"
	"crypto/rsa"
)

var (
	JwtHandlerPool = &sync.Pool{New: func() interface{} {
		return &JwtHandler{}
	}}
	//这样设置影响单测,应移动至具体的生成和验证的逻辑代码处
	//privateKey = LoadRSAPrivateKeyFromDisk(beego.AppPath + "/keys/jwt_private.pem")
	//publicKey  = LoadRSAPublicKeyFromDisk(beego.AppPath + "/keys/jwt_public.pem")
)
func NewJwtHandler() *JwtHandler{
	return JwtHandlerPool.Get().(*JwtHandler)
}
type JwtClaims struct{
	Uid string `json:"uid"`
	Uname string `json:"uname"`
	jwt.StandardClaims
}
type JwtRefreshClaims struct{
	Token string `json:token`
	jwt.StandardClaims
}
type JwtHandler struct{
	//for unit testing
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}
func (j *JwtHandler)Generate(claims JwtClaims) (string,error){
	m := jwt.NewWithClaims(jwt.SigningMethodRS256,claims)
	//for unit testing
	pk := j.privateKey
	//if pk == nil{
	//	pk = privateKey
	//}
	k,e := m.SignedString(pk)
	if e != nil{
		beego.Error(e.Error())
	}
	return k,e

}
func (j *JwtHandler)GenerateRefreshToken(claims JwtRefreshClaims) (string,error){
	m := jwt.NewWithClaims(jwt.SigningMethodRS256,claims)
	//for unit testing
	pk := j.privateKey
	//if pk == nil{
	//	pk = privateKey
	//}
	k,e := m.SignedString(pk)
	if e != nil{
		beego.Error(e.Error())
	}
	return k,e

}
func (j *JwtHandler) SetPublicKey(key *rsa.PublicKey){
	j.publicKey = key
}
func (j *JwtHandler) GetPublicKey() *rsa.PublicKey{
	return j.publicKey
}
func (j *JwtHandler) SetPrivateKey(key *rsa.PrivateKey){
	j.privateKey = key
}
func (j *JwtHandler) GetPrivateKey() *rsa.PrivateKey{
	return j.privateKey
}
func (j *JwtHandler)Validate(token string) (*JwtClaims,error){
	parsedClaims := &JwtClaims{}
	//for unit testing
	pk := j.publicKey
	//if pk == nil{
	//	pk = publicKey
	//}
	_,err := jwt.ParseWithClaims(token,parsedClaims,func(*jwt.Token) (interface{}, error) {
		return pk, nil
	})
	if err != nil{
		beego.Error(err.Error())
	}
	return parsedClaims,err
}

func (j *JwtHandler)ValidateRefreshToken(token string) (*JwtRefreshClaims,error){
	parsedClaims := &JwtRefreshClaims{}
	//for unit testing
	pk := j.publicKey
	//if pk == nil{
	//	pk = publicKey
	//}
	_,err := jwt.ParseWithClaims(token,parsedClaims,func(*jwt.Token) (interface{}, error) {
		return pk, nil
	})
	if err != nil{
		beego.Error(err.Error())
	}
	return parsedClaims,err
}
func (j *JwtHandler) Release(){
	JwtHandlerPool.Put(j)
}
