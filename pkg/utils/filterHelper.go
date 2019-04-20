package utils

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

func TransformFieldsCdt(cdt []string, fields map[string]interface{}) map[string]interface{} {
	var finals = map[string]interface{}{}
	for _, v := range cdt {
		cdv := strings.Split(v, "=")
		if f, ok := fields[cdv[0]]; ok {
			finals[f.(string)] = cdv[1]
		}
	}
	return finals
}

//TODO:多种搜索条件支持，比如like,between..and,大于,小于等..
func TransformQset(qs orm.QuerySeter, k string, v string) orm.QuerySeter {
	if strings.Index(v, "~") == 0 {
		qs = qs.Filter(k+"__startswith", v[1:])
	} else {
		qs = qs.Filter(k, v)
	}
	return qs
}
