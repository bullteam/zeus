package models

import (
	"time"
)

type LoginInfo struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type User struct {
	Id            int         `json:"id"`
	Username      string      `json:"username"`
	Mobile        string      `json:"mobile"`
	Sex           int         `json:"sex"`
	Realname      string      `json:"realname"`
	Password      string      `json:"-"`
	Salt          string      `json:"-"`
	Department    *Department `orm:"rel(fk)";json:"department"`
	Faceicon      string      `json:"faceicon"`
	Email         string      `json:"email"`
	Title         string      `json:"title"`
	Status        int         `json:"status"`
	CreateTime    time.Time   `orm:"auto_now_add;type(datetime)" json:"create_time"`
	LastLoginTime time.Time   `orm:"auto_now_add;type(datetime)" json:"-"`
	Roles         []*Role     `orm:"rel(m2m);rel_through(zeus/pkg/models.UserRole)"`
}
