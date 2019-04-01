package dto

import "github.com/astaxie/beego/validation"

var DOMAIN_SEARCH = map[string]interface{}{
	"n" 	: "name",
	"c"		: "code",
}

type DomainaddAddDto struct {
	Name  string  `form:"name"`
	Callbackurl  string `form:"callbackurl"`
	Remark  string `form:"remark"`
	Code string `form:"code"`
}
func (d *DomainaddAddDto) Valid(v *validation.Validation) {

}