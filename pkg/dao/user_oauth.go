package dao

import (
	"github.com/bullteam/zeus/pkg/models"
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

func (dao *UserOAuthDao) Delete(id int) error {
	o := GetOrmer()
	UserOAuth := &models.UserOAuth{Id: id}
	if o.Read(UserOAuth) == nil {
		_, err := o.Delete(UserOAuth)
		if err != nil {
			return err
		}
	}

	return nil
}