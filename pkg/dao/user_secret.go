package dao

import "zeus/pkg/models"

type UserSecretDao struct {
}


func (dao *UserSecretDao) GetSecretByUserid(userid int) (*models.UserSecret, error) {
	o := GetOrmer()
	v := &models.UserSecret{}
	err := o.QueryTable("user_secret").Filter("user_id", userid).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, nil
}

func (dao *UserSecretDao) Create(userSecret models.UserSecret) (int64, error) {
	o := GetOrmer()
	return o.Insert(&userSecret)
}