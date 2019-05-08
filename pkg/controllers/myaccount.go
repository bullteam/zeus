package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/bullteam/zeus/pkg/service"
	"image/png"
	"strconv"

	// "crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/dgryski/dgoogauth"
	"github.com/skip2/go-qrcode"
	"net/url"
)

type MyAccountController struct {
	TokenCheckController
}

func (c *MyAccountController) GetInfo() {
	user_id, err := strconv.Atoi(c.Uid)
	if err != nil {
		c.Fail(components.ErrInvalidUser, err.Error())
		return
	}
	myAccountService := service.MyAccountService{}
	userSecretQuery, err := myAccountService.GetSecret(user_id)
	fmt.Println(userSecretQuery, err)
	if err != nil {
		c.Fail(components.ErrInvalidUser, err.Error())
		return
	}
	account := userSecretQuery.Account_name
	issuer := "宙斯"
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}

	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)
	params := url.Values{}
	params.Add("secret", userSecretQuery.Secret)
	params.Add("issuer", issuer)
	URL.RawQuery = params.Encode()
	p, errs := qrcode.New(URL.String(), qrcode.Medium)
	img := p.Image(256)
	if errs != nil {
		c.Fail(components.ErrInvalidParams, errs.Error())
		return
	}
	out := new(bytes.Buffer)

	errx := png.Encode(out, img)
	if errx != nil {
		c.Fail(components.ErrInvalidParams, errx.Error())
	}
	base64Img := base64.StdEncoding.EncodeToString(out.Bytes())
	data := map[string]string{
		"code ":   "data:image/png;base64," + base64Img,
		"account": account,
		"issuer":  issuer,
		"secret":  userSecretQuery.Secret,
	}
	c.Resp(0, "success", data)
}

func (c *MyAccountController) BindCode() {
	secret := []byte{'H', 'e', 'l', 'l', 'o', '!', 0xDE, 0xAD, 0xBE, 0xEF} //TODO 暂时固定
	secretBase32 := base32.StdEncoding.EncodeToString(secret)              //TODO 通过数据库/Redis存储
	bindCodeDto := &dto.BindCodeDto{}
	err := c.ParseAndValidateFirstErr(bindCodeDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	//account := service.MyAccountService{}

	otpc := &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  3,
		HotpCounter: 0,
		// UTC:         true,
	}

	val, err := otpc.Authenticate(bindCodeDto.Google2faToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !val {
		c.Fail(components.ErrGoogleBindCode)
		return
	}
	data := "Authenticated!"
	c.Resp(0, "success", data)
}
