package components

import (
	"github.com/dgrijalva/jwt-go"
)

var Args args
type args struct {
	ConfigFile     string
}

type ControllerError struct {
	Code     int    `json:"code"`
	Message  string `json:"msg"`
	Langkey 	 string `json:"langkey"`
	Moreinfo string `json:"moreinfo"`
}
type ConrollerResp struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data"`
}

var (
	Err404               = &ControllerError{404, "page not found", "",""}
	ErrInputData         = &ControllerError{10001, "数据输入错误", "",""}
	ErrDatabase          = &ControllerError{10002, "服务器错误", "",""}
	ErrDupUser           = &ControllerError{10003, "用户信息已存在", "",""}
	ErrNoUser            = &ControllerError{10004, "用户信息不存在", "",""}
	ErrPass              = &ControllerError{10005, "用户信息不存在或密码不正确", "",""}
	ErrNoUserPass        = &ControllerError{10006, "用户信息不存在或密码不正确", "",""}
	ErrNoUserChange      = &ControllerError{10007, "用户信息不存在或数据未改变", "",""}
	ErrInvalidUser       = &ControllerError{10008, "用户信息不正确", "",""}
	ErrOpenFile          = &ControllerError{10009, "服务器错误", "",""}
	ErrWriteFile         = &ControllerError{10010, "写文件出错", "",""}
	ErrSystem            = &ControllerError{10011, "操作系统错误", "",""}
	ErrExpired           = &ControllerError{10012, "登录已过期", "",""}
	ErrPermission        = &ControllerError{10013, "没有权限", "",""}
	Actionsuccess        = &ControllerError{90000, "操作成功", "",""}
	ErrGenJwt            = &ControllerError{10014, "获取令牌失败", "",""}
	ErrChkJwt            = &ControllerError{10012, "无效的令牌", "",""}
	ErrIdData            = &ControllerError{10016, "此ID无数据记录", "err.ErrIdData",""}
	ErrAddFail           = &ControllerError{11000, "创建失败", "",""}
	ErrEditFail          = &ControllerError{11001, "更新失败", "",""}
	ErrDelFail           = &ControllerError{11002, "删除失败", "",""}
	ErrInvalidParams     = &ControllerError{11003, "验证失败", "",""}
	ErrRoleAssignFail    = &ControllerError{12000, "权限分配失败", "",""}
	ErrMenuData          = &ControllerError{12001, "请传递菜单ids", "",""}
	ErrCaptchaEmpty      = &ControllerError{13001, "验证码不能为空", "",""}
	ErrCaptcha           = &ControllerError{13002, "验证码错误", "",""}
	ErrDeptDel           = &ControllerError{13003, "部门无法删除", "","部门内仍有成员,请先行转移到其它部门"}
	ErrDeptHasMember     = &ControllerError{13004, "部门不可删除", "","部门内仍有成员"}
	ErrDupRecord         = &ControllerError{13005, "记录已存在", "",""}
	ErrWrongRefreshToken = &ControllerError{13006, "无效的refresh令牌", "",""}
	//Err
)

type Claims struct {
	Appid string `json:"appid"`
	jwt.StandardClaims
}