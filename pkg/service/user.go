package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/dao"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/utils"
	dingtalk "github.com/icepy/go-dingtalk/src"
	"strconv"
)

type UserService struct {
	dao          *dao.UserDao
	roleDao      *dao.RoleDao
	domainDao    *dao.DomainDao
	urDao        *dao.UserRoleDao
	userOAuthDao *dao.UserOAuthDao
}

type DingtalkUserInfo struct {
	Openid  string
	Unionid string
	Nick    string
	Dingid  string
}

func (us *UserService) GetList(start, limit int, q []string) ([]*models.User, int64) {
	return us.dao.List(start, limit, q)
}

func (us *UserService) DisplayCapcha(username string) bool {
	return us.dao.DisplayCapcha(username)
}

func (us *UserService) SetCapcha(username string) bool {
	return us.dao.SetCapcha(username)
}

func (us *UserService) CheckPass(pass string, u models.User) (bool, error) {
	return us.dao.CheckPass(pass, u)
}

func (us *UserService) FindByUserName(name string) (user models.User, err error) {
	return us.dao.FindByUserName(name)
}

func (us *UserService) GetUser(id int) (userInfo *models.User, err error) {
	return us.dao.GetUser(id)
}

func (us *UserService) NewUser(dto *dto.UserAddDto) (id int64, err error) {
	return us.dao.NewUser(dto)
}

func (us *UserService) UpdateUser(dto *dto.UserEditDto) error {
	return us.dao.UpdateUser(dto)
}

func (us *UserService) UpdateStatus(id int, status int) error {
	return us.dao.UpdateStatus(id, status)
}

func (us *UserService) UpdatePassword(id int, newPwd string) error {
	return us.dao.UpdatePassword(id, newPwd)
}

func (us *UserService) UpdateDepartment(uids []interface{}, did int) (int64, error) {
	return us.dao.UpdateDepartment(uids, did)
}

func (us *UserService) UserList(page int, offset int) (user1 []*models.User, cnt int64) {
	return us.dao.UserList(page, offset)
}

func (us *UserService) DeleteUser(id int) error {
	return us.dao.DeleteUser(id)
}

func (us *UserService) GetUserByUid(uid int64) (user []orm.Params, err error) {
	return us.dao.GetUserByUid(uid)
}

func (us *UserService) GetRolesByUid(uid string) []orm.Params {
	return us.roleDao.GetRolesAndDomainByUid(uid)
}

func (us *UserService) AddRoles(uid int64, roles []string) bool {
	//1.remove previously connections
	_, _ = us.urDao.DeleteByUid(uid)
	for _, role := range roles {
		rid, err := strconv.Atoi(role)
		if err == nil {
			//create new connections
			//ignore duplicated record error
			_, _ = us.urDao.Create(int(uid), rid)
		} else {
			return false
		}
	}
	return true
}

//获取用户相关项目域
func (us *UserService) GetRelatedDomains(uid string) []map[string]interface{} {
	domains := us.roleDao.GetRolesAndDomainByUid(uid)
	var d []map[string]interface{}
	s := map[interface{}]bool{}
	for _, domain := range domains {
		if _, ok := s[domain["domain_code"]]; !ok && domain["domain_code"] != "root" {
			d = append(d, map[string]interface{}{
				"name": domain["domain_name"],
				"code": domain["domain_code"],
				"url":  domain["domain_url"],
			})
			s[domain["domain_code"]] = true
		}
	}
	return d
}

func (us *UserService) GetMenusByDomain(uid string, domain string) []orm.Params {
	menuService := MenuService{}
	roles := us.roleDao.GetRolesAndDomainByUid(uid)
	var menus string
	d, err := us.domainDao.GetDomainByCode(domain)
	if err != nil {
		return nil
	}
	for _, role := range roles {
		if role["domain_id"].(string) == strconv.Itoa(d.Id) {
			if menus != "" {
				menus += ","
			}
			menus += role["menu_ids"].(string)
		}
	}
	if menus != "" {
		return menuService.GetMenusByIds(menus)
	}
	return nil
}

//修改自己的密码
func (us *UserService) ResetPassword(uid int, dto *dto.PwdResetDto) error {
	return us.dao.UpdatePassword(uid, dto.NewRePwd)
}

//修改其它用户的密码
func (us *UserService) ResetUserPassword(uid int, dto *dto.PwdUserResetDto) error {
	return us.dao.UpdatePassword(uid, dto.NewRePwd)
}

func (us *UserService) SwitchDepartment(uids []string, did int) (int64, error) {
	euid := []interface{}{}
	for _, s := range uids {
		euid = append(euid, s)
	}
	return us.dao.UpdateDepartment(euid, did)
}

//绑定钉钉
func (us *UserService) BindByDingtalk(code string, uid int) (openid string,err error) {
	Info,err := getUserInfo(code)
	if err != nil {
		return  "",err
	}
	User,errs := us.userOAuthDao.GetUserByOpenId(Info.Openid, 1)
	if errs != nil || !utils.IsNilObject(User){
		return  "",nil
	}
	beego.Debug(Info)
	userOAuth := models.UserOAuth{
		From:    1, // 1表示钉钉
		User_id: uid,
		Name:    Info.Nick,
		Openid:  Info.Openid,
		Unionid: Info.Unionid,
	}
	us.userOAuthDao.Create(userOAuth)
	return Info.Openid,nil
}

//钉钉登陆
func (us *UserService) LoginByDingtalk(code string) (user *models.UserOAuth ,err error) {
	Info,err := getUserInfo(code)
	if err != nil {
		return nil,err
	}
	User, err := us.userOAuthDao.GetUserByOpenId(Info.Openid, 1)
	if err == nil {
		return User,nil
	}
	return nil,err
}

func getUserInfo(code string) (UserInfo *DingtalkUserInfo,err error) {
	c := GetCompanyDingTalkClient()
	c.RefreshSNSAccessToken()
	perInfo, err := c.SNSGetPersistentCode(code)
	if err != nil {
		return nil,err
	}
	snstoken, err := c.SNSGetSNSToken(perInfo.OpenID, perInfo.PersistentCode)
	if err != nil {
		return nil,err
	}
	Info, _ := c.SNSGetUserInfo(snstoken.SnsToken)
	userInfo := &DingtalkUserInfo{
		Info.UserInfo.OpenID,
		Info.UserInfo.UnionID,
		Info.UserInfo.Nick,
		Info.UserInfo.DingID,
	}
	return userInfo,nil
}

func (us *UserService) UnBindUserDingtalk(oauthid int) error {
	return us.userOAuthDao.Delete(oauthid)
}

func GetCompanyDingTalkClient() *dingtalk.DingTalkClient {
	SNSAppID := beego.AppConfig.String("dingtalk::SNSAppID")
	SNSSecret := beego.AppConfig.String("dingtalk::SNSSecret")
	config := &dingtalk.DTConfig{
		SNSAppID:  SNSAppID,
		SNSSecret: SNSSecret,
	}
	return dingtalk.NewDingTalkCompanyClient(config)
}

func (us *UserService) GetBindOauthUserInfo(userid int) (UserInfo *models.UserOAuth, err error)  {
	return us.userOAuthDao.GetUserById(userid)
}