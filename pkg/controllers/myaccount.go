package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/bullteam/zeus/pkg/service"
	"github.com/bullteam/zeus/pkg/utils/mailTemplate"
	"image/png"
	"strconv"
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
		"secret":  userSecretQuery.Secret,
	}
	c.Resp(0, "success", data)
}

func (c *MyAccountController) BindCode() {
	user_id, err := strconv.Atoi(c.Uid)
	if err != nil {
		c.Fail(components.ErrInvalidUser, err.Error())
		return
	}
	myAccountService := service.MyAccountService{}
	userSecretQuery, err := myAccountService.GetSecret(user_id)
	if err != nil {
		c.Fail(components.ErrInvalidUser, err.Error())
		return
	}
	secretBase32 := userSecretQuery.Secret
	bindCodeDto := &dto.BindCodeDto{}
	errs := c.ParseAndValidateFirstErr(bindCodeDto)
	if errs != nil {
		c.Fail(components.ErrInvalidParams, errs.Error())
		return
	}

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

func (c *MyAccountController) Third() {
	user_id, err := strconv.Atoi(c.Uid)
	if err != nil {
		c.Fail(components.ErrInvalidUser, err.Error())
		return
	}
	myAccountService := service.MyAccountService{}
	oauthList := myAccountService.GetThirdList(user_id)
	c.Resp(0, "success", oauthList)
}

/**
  验证邮件地址(发送邮件)
 */
func (c *MyAccountController) Verifymail() {
	verifyEmailDto := &dto.VerifyEmail{}
	errs := c.ParseAndValidateFirstErr(verifyEmailDto)
	if errs != nil {
		c.Fail(components.ErrInvalidParams, errs.Error())
		return
	}
	username := beego.AppConfig.String("email::username")
	password := beego.AppConfig.String("email::password")
	host := beego.AppConfig.String("email::host")
	port,_ := beego.AppConfig.Int("email::port")
	from := beego.AppConfig.String("email::from")
	if port == 0 {
		port = 25
	}
	config := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%d,"from":"%s"}`, username, password, host, port, from)
	temail := utils.NewEMail(config)
	temail.To = []string{verifyEmailDto.Email}//指定收件人邮箱地址
	temail.From = from//指定发件人的邮箱地址
	temail.Subject = "验证账号邮件"//指定邮件的标题
	temail.HTML = mailTemplate.MailBody()
	err := temail.Send()
	if err != nil{
		c.Fail(components.ErrSendMail, err.Error())
		return
	}
	c.Resp(0, "success", "邮件发送成功！")
}