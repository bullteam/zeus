package dao

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zeus/pkg/components"
	"zeus/pkg/dto"
	"zeus/pkg/models"
	"zeus/pkg/utils"
	"golang.org/x/crypto/scrypt"
	"io"
	"strconv"
	"time"
)

type UserDao struct {
}

const pwHashBytes = 64

func (dao *UserDao) List(start int, limit int, q []string) ([]*models.User, int64) {
	o := GetOrmer()
	var users []*models.User
	qs := o.QueryTable("user").Filter("status__gt", -1)
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.UserSearch) {
			qs = utils.TransformQset(qs, k, v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	_, _ = qs.Limit(limit, start).RelatedSel("department").All(&users)
	c, _ := qs.Count()
	for _, v := range users {
		_, _ = orm.NewOrm().LoadRelated(v, "Roles")
	}
	return users, c
}

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h), nil
}

//10分钟内超过三次错误显示验证码
func (dao *UserDao) DisplayCapcha(username string) bool {
	if components.Cache.IsExist("login:" + username) {
		times := components.Cache.Get("login:" + username)
		cnt, err := strconv.ParseInt(string(times.([]byte)), 10, 64)
		if err != nil {
			return false
		}
		if cnt <= 3 {
			beego.Debug(cnt)
			return false
		}
		return true
	}
	return false
}

func (dao *UserDao) SetCapcha(username string) bool {
	if components.Cache.IsExist("login:" + username) {
		if err := components.Cache.Incr("login:" + username); err != nil {
			beego.Debug("incr err")
		}
	} else {
		_ = components.Cache.Put("login:"+username, 1, 1000*time.Second)
		beego.Debug("put 1")
	}
	beego.Debug(components.Cache.Get("login:" + username))
	return true
}

//验证帐号密码
func (dao *UserDao) CheckPass(pass string, u models.User) (bool, error) {
	hash, err := generatePassHash(pass, u.Salt)
	if err != nil {
		return false, err
	}
	return u.Password == hash, nil
}

func (dao *UserDao) FindByUserName(name string) (user models.User, err error) {
	o := GetOrmer()
	err = o.Raw("select * from user where username = ?", name).QueryRow(&user)

	return user, err
}

//根据id取得用户信息
func (dao *UserDao) GetUser(id int) (userInfo *models.User, err error) {
	o := GetOrmer()
	user := new(models.User)
	qs := o.QueryTable("user")
	err = qs.Filter("id", id).One(user)
	if err != nil {
		return userInfo, err
	}
	return user, err
}

func (dao *UserDao) NewUser(dto *dto.UserAddDto) (id int64, err error) {
	salt, err := generateSalt()
	hash, err := generatePassHash(dto.Password, salt)
	if err != nil {
		return 0, err
	}
	var dept models.Department
	dept.Id = dto.DepartmentId
	o := GetOrmer()
	_ = o.Read(&dept, "id")

	users := &models.User{
		Username:   dto.Username,
		Password:   hash,
		Mobile:     dto.Mobile,
		Sex:        dto.Sex,
		Realname:   dto.Realname,
		Email:      dto.Email,
		Status:     dto.Status,
		Faceicon:   dto.Faceicon,
		Department: &dept,
		Salt:       salt,
		Title:      dto.Title,
	}
	id, err = o.Insert(users)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//修改用户信息
func (dao *UserDao) UpdateUser(dto *dto.UserEditDto) error {
	var dept models.Department
	dept.Id = dto.DepartmentId
	o := GetOrmer()
	_ = o.Read(&dept, "id")

	users := &models.User{Id: dto.Id}
	if o.Read(users) == nil {
		users.Username = dto.Username
		users.Mobile = dto.Mobile
		users.Sex = dto.Sex
		users.Realname = dto.Realname
		users.Email = dto.Email
		users.Status = dto.Status
		users.Faceicon = dto.Faceicon
		users.Department = &dept
		users.Title = dto.Title
		var err error
		if dto.Password != "" { //判断密码是否为空，如果为空不更新
			salt, err := generateSalt()
			hash, err := generatePassHash(dto.Password, salt)
			if err != nil {
				return err
			}
			users.Password = hash
			users.Salt = salt
			_, err = o.Update(users, "Username", "Mobile", "Sex", "Realname", "Email", "Status", "Faceicon", "Department", "Title", "Password", "Salt")
		} else {
			_, err = o.Update(users, "Username", "Mobile", "Sex", "Realname", "Email", "Status", "Faceicon", "Department", "Title")
		}
		if err != nil {
			return err
		}
	}

	return nil
}

//修改用户状态
func (dao *UserDao) UpdateStatus(id int, status int) error {
	o := GetOrmer()
	users := &models.User{Id: id}
	if o.Read(users) == nil {
		users.Status = status
		_, err := o.Update(users)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao *UserDao) UpdatePassword(id int, newPwd string) error {
	o := GetOrmer()
	user := &models.User{Id: id}
	err := o.Read(user)
	if err == nil {
		//if ok,_ := user.CheckPass(oldpwd);!ok{
		//	return errors.New("原密码不正确")
		//}
		salt, _ := generateSalt()
		hash, _ := generatePassHash(newPwd, salt)
		user.Password = hash
		user.Salt = salt
		_, err := o.Update(user)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

//将一个用户或多个用户移到某个部门下
func (dao *UserDao) UpdateDepartment(uids []interface{}, did int) (int64, error) {
	o := GetOrmer()
	return o.QueryTable("user").Filter("id__in", uids...).Update(orm.Params{"department_id": did})
}

func (dao *UserDao) UserList(page int, offset int) (user1 []*models.User, cnt int64) {
	var user []*models.User
	o := GetOrmer()
	qs := o.QueryTable("user").RelatedSel("department")
	counts, _ := qs.Count()
	start := (page - 1) * offset
	_, _ = qs.Limit(offset, start).All(&user)
	for _, v := range user {
		_, _ = orm.NewOrm().LoadRelated(v, "Roles")
	}

	return user, counts
}

//删除
func (dao *UserDao) DeleteUser(id int) error {
	o := GetOrmer()
	user := &models.User{Id: id}
	if o.Read(user) == nil {
		user.Status = -1
		_, err := o.Update(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao *UserDao) GetUserByUid(uid int64) (user []orm.Params, err error) {
	var users []orm.Params
	o := GetOrmer()
	_, err = o.Raw(`SELECT u.username,u.mobile,u.sex,u.realname,u.email,u.status,u.faceicon,d.name as dept_name,u.department_id FROM user u
		   LEFT JOIN department d ON u.department_id = d.id  WHERE u.id = ?`, uid).Values(&users)
	if err != nil {
		return user, err
	}

	return users, err
}
