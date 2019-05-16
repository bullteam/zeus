package tests

import (
	"github.com/dgrijalva/jwt-go"
	"path/filepath"
	"testing"
	"time"
	"zeus/pkg/components"
	"zeus/pkg/utils"
)

var (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQC/TYKuXsgYdoICfEZOiy1L12CbyPdudhrCjrjwVcIrhGNn6Udq
/SY5rh0ixm09I2tXPWLYuA1R55kyeo5RPFX+FrD+mQwfJkV/QfhaPsNjU4nCEHFM
trsYCcLYJs9uX0tJdAtE6sg/VSulg1aMqCNWvtVtjrrVXSbu4zbyWzVkxQIDAQAB
AoGAIJrzZQjejdzU99t6mDR8eeqxmpu8IGWc1gBBYSUcvRIJZ1KJS6Dt/PLCIIU1
ZTA+QVZDHLDyBD23DLV6wDnKZgKQnTQSqfPeanT5Zomc96QUmtQBqZ5m/3P4LXTV
HlYZPYeKlOCvL6fUtrb8o9sx2jk1T+d1da4CfmOffI/ZQ8ECQQDuFP+pz3m+DnMk
yozNunRr8XcCk460A7lmHoCOg6l+jPbeXGEgQCBkzdgeEFWj1S8dAjyIETNGE7UI
gUI5o8ydAkEAzbM9k/p08fJr+Sd8+WlYfT9NeDJpcJLXQp0FekJv9bNgRaYr0Wfj
vW7XWV8VyephgAIC8S38CNyUPgqerrR8SQJAJ+9rxx8fK6se01AKeEPLXYPeU5dO
u5FYWvHI3J7nImwgyMG0JQW8qUwB8WEKDHYo9fO3FZfVAu8xUaDk6+g23QJAUJuW
2/Bf95g6O67/yHVB2gL+hsWqkBTbCh2iUeDLIwuiBGkz7qG5mzheZ4VdcnzIrHMd
WAnfJFHcPdvHh0rvEQJAVuiKX8+EPCpdymbmrK3dH3E2WHjozSksrJT0v6NdFnqY
zeJFUH2EGnbk4M5+33zM8t3f53aGdgdVTA9VOonaow==
-----END RSA PRIVATE KEY-----`
	publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/TYKuXsgYdoICfEZOiy1L12Cb
yPdudhrCjrjwVcIrhGNn6Udq/SY5rh0ixm09I2tXPWLYuA1R55kyeo5RPFX+FrD+
mQwfJkV/QfhaPsNjU4nCEHFMtrsYCcLYJs9uX0tJdAtE6sg/VSulg1aMqCNWvtVt
jrrVXSbu4zbyWzVkxQIDAQAB
-----END PUBLIC KEY-----`
)

func tokenGenerate(t *testing.T) string {
	jh := components.NewJwtHandler()
	//jh.privateKey = loadRSAPrivateKeyFromDisk("../keys/jwt_private.pem")
	//jh.SetPrivateKey("../keys/jwt_private.pem")
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Error(err.Error())
	}
	jh.SetPrivateKey(key)
	defer jh.Release()
	claims := components.JwtClaims{Uid: "123456", Uname: "tester"}
	claims.ExpiresAt = time.Now().Unix() + 1
	token, err := jh.Generate(claims)
	if err != nil {
		t.Error(err.Error())
	}
	return token
}
func TestJwtHandler_Generate(t *testing.T) {
	t.Logf("Generated token : %s", tokenGenerate(t))
}

func TestJwtHandler_Validate(t *testing.T) {
	token := tokenGenerate(t)
	jh := components.NewJwtHandler()
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Error(err.Error())
	}
	//jh.publicKey = loadRSAPublicKeyFromDisk("../keys/jwt_public.pem")
	jh.SetPublicKey(key)
	defer jh.Release()
	claims, err := jh.Validate(token)
	if err != nil {
		t.Errorf("Validate error : %s", err.Error())
	} else {
		t.Logf("Extra user : %s,%s", claims.Uid, claims.Uname)
	}
}

func TestJwtHandler_Expired(t *testing.T) {
	token := tokenGenerate(t)
	time.Sleep(time.Second * 2)
	jh := components.NewJwtHandler()
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Error(err.Error())
	}
	//jh.publicKey = loadRSAPublicKeyFromDisk("../keys/jwt_public.pem")
	jh.SetPublicKey(key)
	defer jh.Release()
	v, err := jh.Validate(token)
	if err == nil {
		t.Errorf("Expired validate error ")
	} else {
		t.Logf("Expired check successful,claims should be null like %x", v)
	}
}

func TestToken(t *testing.T) {
	tokenString := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMiIsInVuYW1lIjoi6buE56aP56WlIiwiZXhwIjoxNTU3ODM2MDQwfQ.PM1M7QiNWRyy1LpI24woiJ-1TOGQpiEGjzPfiG2NjPg8l7iXXfsZhhuzXRaKhEY6PVZ2_Td_waBb5fH3llCoKYlPZASr8jbXcK36VDO9y6KMUNAv_-57yz9H8Mq1I9CISWrPxOMnNircGIdlAeB6OimR41oo5SiV713qQtEu2wA"
	jwt_public_key, err := filepath.Abs("../../conf/keys/jwt_public.pem")
	if err != nil {
		t.Errorf("error token")
	}
	jh := components.NewJwtHandler()
	jh.SetPublicKey(utils.LoadRSAPublicKeyFromDisk(jwt_public_key))
	defer jh.Release()
	claims, _ := jh.Validate(tokenString)
	t.Log(claims)
}
