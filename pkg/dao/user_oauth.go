package dao

import (
	"github.com/astaxie/beego/orm"
	"zeus/pkg/models"
)

type UserOAuthDao struct {
}

func (dao *UserOAuthDao) GetUserById(userid int) (*models.UserOAuth, error) {
	o := GetOrmer()
	v := &models.UserOAuth{}
	err := o.QueryTable("user_oauth").Filter("user_id", userid).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, nil
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

func (dao *UserOAuthDao) Create(userOAuth models.UserOAuth) (int64, error) {
	o := GetOrmer()
	return o.Insert(&userOAuth)
}

func (dao *UserOAuthDao) DeleteByUseridAndFrom(from int,user_id int) error {
	o := GetOrmer()
	_, err := o.QueryTable("user_oauth").Filter("from", from).Filter("user_id",user_id).Delete()
	if err != nil {
		return err
	}
	return nil
}

/**
 	获取绑定列表
 */
func (dao *UserOAuthDao) List(user_id int) (oauthlists []orm.Params) {
	var oauthlist []orm.Params
	o := GetOrmer()
	sql := "select `from`,name from user_oauth where user_id = ?"
	_, err := o.Raw(sql, user_id).Values(&oauthlist)
	if err != nil {
		return oauthlists
	}
	return oauthlist
}