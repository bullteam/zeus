package controllers

import "github.com/dchest/captcha"

type CaptchaController struct {
	BaseController
}

func (c *CaptchaController) Get() {
	Captcha := struct{ Id string }{captcha.NewLen(6)} //验证码长度为6
	c.Resp(0, "success", map[string]interface{}{
		"captcha": &Captcha,
	})
}
