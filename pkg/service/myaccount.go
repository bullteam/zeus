package service

import (
	"encoding/base32"
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zeus/pkg/dao"
	"zeus/pkg/models"
	"zeus/pkg/utils"
)

type MyAccountService struct {
	dao          *dao.UserSecretDao
	oauthdao        *dao.UserOAuthDao
}

// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func (s *MyAccountService) GetSecret(uid int) (userSecretQuery models.UserSecretQuery,err error) {
	v,err := s.dao.GetSecretByUserid(uid)
	if err != nil  {
		return userSecretQuery,err
	}
	if !utils.IsNilObject(v) {
		userSecretQuery.Account_name = v.Account_name
		userSecretQuery.Secret = v.Secret
		return userSecretQuery,nil
	}
	secret := make([]byte, 10)
	_,errs := rand.Read(secret)
	if errs != nil {
		return userSecretQuery,errs
	}
	secretBase32 := base32.StdEncoding.EncodeToString(secret)
	userSecretQuery.Account_name = fmt.Sprintf("Zeus:%d",uid)
	userSecretQuery.Secret = secretBase32
	usersecret := models.UserSecret{
		User_id:uid,
		Account_name:userSecretQuery.Account_name,
		Secret:secretBase32,
	}
	s.dao.Create(usersecret)
  return  userSecretQuery,nil
}

/**
 	获取第三方账号绑定列表
 */
func (s *MyAccountService) GetThirdList(user_id int) (oauthlist []orm.Params) {
	return s.oauthdao.List(user_id)
}

//绑定第三方应用
func (s *MyAccountService) BindByDingtalk(code string, uid int,from int) (openid string,err error) {
	Info,err := getUserInfo(code)
	if err != nil {
		return  "",err
	}
	User,errs := s.oauthdao.GetUserByOpenId(Info.Openid, from)
	if errs != nil || !utils.IsNilObject(User){
		return  "",nil
	}
	beego.Debug(Info)
	userOAuth := models.UserOAuth{
		From:    from, // 1表示钉钉
		User_id: uid,
		Name:    Info.Nick,
		Openid:  Info.Openid,
		Unionid: Info.Unionid,
	}
	s.oauthdao.Create(userOAuth)
	return Info.Openid,nil
}