package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/utils"
	"path/filepath"
	"reflect"
	"strings"
)

const listRowsPerPage = 20
const maxLimit = 500

type BaseController struct {
	beego.Controller
}

type TokenCheckController struct {
	Uid      string
	Uname    string
	RawToken string
	BaseController
}

var CurrentLang string

var langTypes []*langType // Languages are supported.

type langType struct {
	Lang, Name string
}

// 固定返回的json数据格式
// code: 错误码
// msg: 错误信息
// data: 返回数据
func (self *BaseController) Resp(code int, msg string, data ...interface{}) {
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

func (self *BaseController) Fail(errs *components.ControllerError, moreErrInfo ...string) {
	errs.Message = i18n.Tr(CurrentLang, errs.Langkey)
	self.Data["json"] = errs
	errs.Moreinfo = ""
	for _, v := range moreErrInfo {
		errs.Moreinfo += v + " // "
	}
	self.ServeJSON()
}

func (c *TokenCheckController) Prepare() {
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
	jwt_public_key, err := filepath.Abs(components.Args.ConfigFile + "/keys/jwt_public.pem")
	if err != nil {
		c.Fail(components.ErrChkJwt)
		return
	}
	jh := components.NewJwtHandler()
	jh.SetPublicKey(utils.LoadRSAPublicKeyFromDisk(jwt_public_key))
	defer jh.Release()
	claims, err := jh.Validate(tokenString)
	if err != nil {
		cerr := components.ErrChkJwt
		cerr.Moreinfo = err.Error()
		c.Fail(cerr)
		return
	}
	c.Uid = claims.Uid
	c.Uname = claims.Uname
	c.RawToken = tokenString

	c.setLangVer() //设置语言
}

func (b *BaseController) ParseAndValidate(obj interface{}) {
	if err := b.ParseForm(obj); err != nil {
		b.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	valid := &validation.Validation{}
	if v, _ := valid.Valid(obj); !v {
		errs := ""
		for _, err := range valid.Errors {
			errs += err.Key + ":" + err.Message + " "
		}
		b.Fail(components.ErrInvalidParams, errs)
		return
	}
}

// 解析并验证表单，返回第一个错误信息
func (b *BaseController) ParseAndValidateFirstErr(obj interface{}) error {
	if err := b.ParseForm(obj); err != nil {
		return err
	}
	valid := &validation.Validation{}
	if v, _ := valid.Valid(obj); !v {
		// stuctTag
		tags := make(map[string]interface{})
		s := reflect.TypeOf(obj).Elem()
		for i := 0; i < s.NumField(); i++ {
			tags[s.Field(i).Name] = s.Field(i).Tag.Get("form")
		}
		for _, err := range valid.Errors {
			return errors.New(tags[err.Field].(string) + ":" + err.Message)
		}
	}

	return nil
}

// 获取分页参数
func (b *BaseController) GetPaginationParams() (start, limit int) {
	start, err := b.GetInt("start", 1)
	if err != nil || start <= 0 {
		start = 1
	}

	limit, err = b.GetInt("limit", listRowsPerPage)
	if err != nil || limit <= 0 {
		limit = listRowsPerPage
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return start, limit
}

/**
设置语言
*/
func (b *BaseController) setLangVer() bool {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes := make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := b.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = b.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := b.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "zh-CN"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		b.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	CurrentLang = curLang.Lang
	return isNeedRedir
}
