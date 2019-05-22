package components

import (
	"github.com/dgrijalva/jwt-go"
)

var Args args

type args struct {
	ConfigFile string
}

type ControllerError struct {
	Code     int    `json:"code"`
	Langkey  string `json:"langkey"`
	Message  string `json:"msg"`
	Moreinfo string `json:"moreinfo"`
}
type ConrollerResp struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data"`
}

var (
	Err404               = &ControllerError{404, "err.Err404", "", ""}
	ErrInputData         = &ControllerError{10001, "err.ErrInputData", "", ""}
	ErrDatabase          = &ControllerError{10002, "err.ErrDatabase", "", ""}
	ErrDupUser           = &ControllerError{10003, "err.ErrDupUser", "", ""}
	ErrNoUser            = &ControllerError{10004, "err.ErrNoUser", "", ""}
	ErrPass              = &ControllerError{10005, "err.ErrPass", "", ""}
	ErrNoUserPass        = &ControllerError{10006, "err.ErrNoUserPass", "", ""}
	ErrNoUserChange      = &ControllerError{10007, "err.ErrNoUserChange", "", ""}
	ErrInvalidUser       = &ControllerError{10008, "err.ErrInvalidUser", "", ""}
	ErrOpenFile          = &ControllerError{10009, "err.ErrOpenFile", "", ""}
	ErrWriteFile         = &ControllerError{10010, "err.ErrWriteFile", "", ""}
	ErrSystem            = &ControllerError{10011, "err.ErrSystem", "", ""}
	ErrExpired           = &ControllerError{10012, "err.ErrExpired", "", ""}
	ErrPermission        = &ControllerError{10013, "err.ErrPermission", "", ""}
	Actionsuccess        = &ControllerError{90000, "err.Actionsuccess", "", ""}
	ErrGenJwt            = &ControllerError{10014, "err.ErrGenJwt", "", ""}
	ErrChkJwt            = &ControllerError{10012, "err.ErrChkJwt", "", ""}
	ErrIdData            = &ControllerError{10016, "err.ErrIdData", "", ""}
	ErrAddFail           = &ControllerError{11000, "err.ErrAddFail", "", ""}
	ErrEditFail          = &ControllerError{11001, "err.ErrEditFail", "", ""}
	ErrDelFail           = &ControllerError{11002, "err.ErrDelFail", "", ""}
	ErrInvalidParams     = &ControllerError{11003, "err.ErrInvalidParams", "", ""}
	ErrRoleAssignFail    = &ControllerError{12000, "err.ErrRoleAssignFail", "", ""}
	ErrMenuData          = &ControllerError{12001, "err.ErrMenuData", "", ""}
	ErrCaptchaEmpty      = &ControllerError{13001, "err.ErrCaptchaEmpty", "", ""}
	ErrCaptcha           = &ControllerError{13002, "err.ErrCaptcha", "", ""}
	ErrDeptDel           = &ControllerError{13003, "err.ErrDeptDel", "", "部门内仍有成员,请先行转移到其它部门"}
	ErrDeptHasMember     = &ControllerError{13004, "err.ErrDeptHasMember", "", "部门内仍有成员"}
	ErrDupRecord         = &ControllerError{13005, "err.ErrDupRecord", "", ""}
	ErrWrongRefreshToken = &ControllerError{13006, "err.ErrWrongRefreshToken", "", ""}
	ErrBindDingtalk      = &ControllerError{13007, "err.ErrBindDingtalk", "", ""}
	ErrUnBindDingtalk    = &ControllerError{13008, "err.ErrUnBindDingtalk", "", ""}
	ErrGoogleBindCode    = &ControllerError{13009, "err.ErrGoogleBindCode", "", ""}
	ErrSendMail          = &ControllerError{13010, "err.ErrSendMail", "", ""}
)

type Claims struct {
	Appid string `json:"appid"`
	jwt.StandardClaims
}
