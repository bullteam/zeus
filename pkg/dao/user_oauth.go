package dao

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/models"
)

type UserOAuthDao struct {
}

func (dao *UserOAuthDao) GetUserByOpenId(openid string, from int) (*models.UserOAuth, error) {
	o := GetOrmer()
	v := &models.UserOAuth{}
	err := o.QueryTable("user_oauth").Filter("openid", openid).Filter("from", from).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, nil
}
func (dao *UserOAuthDao) Create(userOAuth models.UserOAuth) (sql.Result, error) {
	o := GetOrmer()
	qs, _ := o.Raw("insert into user_oauth (from,user_id,name,openid,unionid,avatar,extra) values (?,?,?,?,?,?,?)").Prepare()
	result, err := qs.Exec(userOAuth.From, userOAuth.User_id, userOAuth.Name, userOAuth.Openid, userOAuth.Unionid, userOAuth.Avatar, userOAuth.Extra)
	return result, err
}

func (dao *UserOAuthDao) Delete(userOAuth models.UserOAuth) error {
	o := orm.NewOrm()
	qs, _ := o.Raw("delete from user_oauth where id=? limit 1").Prepare()
	_, err := qs.Exec(userOAuth.Id)
	return err
}
