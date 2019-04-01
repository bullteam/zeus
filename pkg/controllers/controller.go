package controllers

import(
	"github.com/astaxie/beego"
	"github.com/bullteam/zeus/pkg/components"
	"strings"
	"github.com/astaxie/beego/validation"
	"github.com/bullteam/zeus/pkg/utils"
)

const LIST_ROWS_PERPAGE  =  20
type BaseController struct{
	beego.Controller
}

type TokenCheckController struct{
	Uid string
	Uname string
	RawToken string
	BaseController
}

// 固定返回的json数据格式
// code: 错误码
// msg: 错误信息
// data: 返回数据
func (self *BaseController) Resp (code int, msg string, data ... interface{}){
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	if len(data) >= 1 {
		out["data"] = data[0]
	}
	if len(data) >= 2 {
		out["total"] = data[1]
	}
	self.Data["json"] = out
	self.ServeJSON()
}

func (self *BaseController) Fail (errs *components.ControllerError,moreErrInfo ...string){
	self.Data["json"] = errs
	errs.Moreinfo = ""
	for _,v := range moreErrInfo{
		errs.Moreinfo += v+" // "
	}
	self.ServeJSON()
}

func (c *TokenCheckController) Prepare(){
	headAuth := c.Ctx.Input.Header("Authorization")
	if headAuth == "" {
		c.Fail(components.ErrChkJwt)
		return
	}
	tokenString := strings.TrimSpace(headAuth[len("Bearer "):])
	if tokenString == "" {
		c.Fail(components.ErrChkJwt)
		return
	}
	jh := components.NewJwtHandler()
	jh.SetPublicKey(utils.LoadRSAPublicKeyFromDisk(beego.AppPath + "/keys/jwt_public.pem"))
	defer jh.Release()
	claims,err := jh.Validate(tokenString)
	if err != nil {
		cerr := components.ErrChkJwt
		cerr.Moreinfo = err.Error()
		c.Fail(cerr)
		return
	}
	c.Uid = claims.Uid
	c.Uname = claims.Uname
	c.RawToken = tokenString
}

func(b *BaseController)ParseAndValidate(obj interface{}){
	if err := b.ParseForm(obj);err != nil {
		b.Fail(components.ErrInvalidParams,err.Error())
		return
	}
	valid := &validation.Validation{}
	if v,_ := valid.Valid(obj);!v{
		errs := ""
		for _, err := range valid.Errors {
			errs += err.Key + ":" + err.Message+" "
		}
		b.Fail(components.ErrInvalidParams,errs)
		return
	}
}