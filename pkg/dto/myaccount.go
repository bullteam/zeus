package dto

type BindCodeDto struct {
	Google2faToken string    `form:"google_2fa_token" valid:"Required"`          // 验证码
}
