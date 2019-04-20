package models

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/scrypt"
	"io"
	"strconv"
	"time"
)

// RegisterForm definiton.
type RegisterForm struct {
	Email    string `form:"email"    valid:"Required"`
	Password string `form:"password" valid:"Required"`
	Username string `form:"username"`
}

// LoginForm definiton.
type LoginForm struct {
	Username   string `form:"username"`
	Password   string `form:"password" valid:"Required"`
	CaptchaId  string `form:"captchaid" valid:"Required"`
	CaptchaVal string `form:"captchaval" valid:"Required"`
}
type LoginInfo struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
type ChangeuserroleForm struct {
	Id       string `form:"id"`
	Username string `form:"username"`
}

type User struct {
	Id              int         `json:"id"`
	Username        string      `json:"username"`
	Mobile          string      `json:"mobile"`
	Sex             int         `json:"sex"`
	Realname        string      `json:"realname"`
	Password        string      `json:"-"`
	Salt            string      `json:"-"`
	Department      *Department `orm:"rel(fk)";json:"department"`
	Faceicon        string      `json:"faceicon"`
	Email           string      `json:"email"`
	Title           string      `json:"title"`
	Status          int         `json:"status"`
	Create_time     time.Time   `orm:"auto_now_add;type(datetime)" json:"create_time"`
	Last_login_time time.Time   `orm:"auto_now_add;type(datetime)" json:"-"`
	Roles           []*Role     `orm:"rel(m2m);rel_through(github.com/bullteam/zeus/pkg/models.UserRole)"`
}

type DeptModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const pwHashBytes = 64

func init() {
	orm.RegisterModel(new(User))
}
func (u User) List(start int, limit int, q []string) ([]*User, int64) {
	o := orm.NewOrm()
	var users []*User
	qs := o.QueryTable("user").Filter("status__gt", -1)
	//qs.RelatedSel("domain","domain_id","id")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.USER_SEARCH) {
			qs = utils.TransformQset(qs, k, v.(string))
			//qs = qs.Filter(k, v)
		}
	}
	//后期加入搜索条件可利用q参数
	qs.Limit(limit, start).RelatedSel("department").All(&users)
	c, _ := qs.Count()
	for _, v := range users {
		orm.NewOrm().LoadRelated(v, "Roles")
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
func DisplayCapcha(username string) bool {
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

func SetCapcha(username string) bool {
	if components.Cache.IsExist("login:" + username) {
		if err := components.Cache.Incr("login:" + username); err != nil {
			beego.Debug("incr err")
		}
	} else {
		components.Cache.Put("login:"+username, 1, 1000*time.Second)
		beego.Debug("put 1")
	}
	beego.Debug(components.Cache.Get("login:" + username))
	return true
}

//验证帐号密码

func (u *User) CheckPass(pass string) (bool, error) {
	hash, err := generatePassHash(pass, u.Salt)
	if err != nil {
		return false, err
	}
	return u.Password == hash, nil
}

func (u *User) FindByID(name string) (result int, err error) {
	o := orm.NewOrm()
	o.Using("default")
	err = o.Raw("select * from user where username = ?", name).QueryRow(&u)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.Password)
		result = u.Id
	}
	return
}

//根据id取得用户信息
func GetUser(id int) (userinfo *User, err error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err = qs.Filter("id", id).One(user)
	if err != nil {
		return userinfo, err
	}
	return user, err
}

func NewUser(username string, password string, mobile string, sex int, realname string,
	email string, status int, faceicon string, dept_id int, title string) (id int64, err error) {
	salt, err := generateSalt()
	hash, err := generatePassHash(password, salt)
	if err != nil {
		return 0, err
	}
	var dept Department
	dept.Id = dept_id
	newOrm := orm.NewOrm()
	newOrm.Read(&dept, "id")

	users := &User{
		Username:        username,
		Password:        hash,
		Mobile:          mobile,
		Sex:             sex,
		Realname:        realname,
		Email:           email,
		Status:          status,
		Faceicon:        faceicon,
		Department:      &dept,
		Salt:            salt,
		Title:           title,
		Create_time:     time.Now(),
		Last_login_time: time.Now(),
	}
	o := orm.NewOrm()
	o.Using("default")
	id, err = o.Insert(users)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改用户信息
func UpdateUser(id int, username string, password string, mobile string, sex int, realname string,
	email string, status int, faceicon string, dept_id int, title string) error {
	var dept Department
	dept.Id = dept_id
	newOrm := orm.NewOrm()
	newOrm.Read(&dept, "id")

	o := orm.NewOrm()
	users := &User{Id: id}
	if o.Read(users) == nil {
		users.Username = username
		users.Mobile = mobile
		users.Sex = sex
		users.Realname = realname
		users.Email = email
		users.Status = status
		users.Faceicon = faceicon
		users.Department = &dept
		users.Title = title
		var err error
		if password != "" { //判断密码是否为空，如果为空不更新
			salt, err := generateSalt()
			hash, err := generatePassHash(password, salt)
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
func UpdateStatus(id int, status int) error {
	o := orm.NewOrm()
	users := &User{Id: id}
	if o.Read(users) == nil {
		users.Status = status
		_, err := o.Update(users)
		if err != nil {
			return err
		}
	}
	return nil
}
func UpdatePassword(id int, newpwd string) error {
	o := orm.NewOrm()
	user := &User{Id: id}
	err := o.Read(user)
	if err == nil {
		//if ok,_ := user.CheckPass(oldpwd);!ok{
		//	return errors.New("原密码不正确")
		//}
		salt, _ := generateSalt()
		hash, _ := generatePassHash(newpwd, salt)
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
func UpdateDepartment(uids []interface{}, did int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("id__in", uids...).Update(orm.Params{"department_id": did})
}
func User_list(page int, offset int) (user1 []*User, cnt int64) {
	var user []*User
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("user").RelatedSel("department")
	counts, _ := qs.Count()
	start := (page - 1) * offset
	qs.Limit(offset, start).All(&user)
	for _, v := range user {
		orm.NewOrm().LoadRelated(v, "Roles")
	}
	return user, counts
}

//删除
func DeleteUser(id int) error {
	o := orm.NewOrm()
	user := &User{Id: id}
	if o.Read(user) == nil {
		user.Status = -1
		//_, err := o.Delete(user)
		_, err := o.Update(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserByUid(uid int64) (user []orm.Params, err error) {
	var users []orm.Params
	o := orm.NewOrm()
	_, err = o.Raw(`SELECT u.username,u.mobile,u.sex,u.realname,u.email,u.status,u.faceicon,d.name as dept_name,u.department_id FROM user u
		   LEFT JOIN department d ON u.department_id = d.id  WHERE u.id = ?`, uid).Values(&users)
	if err != nil {
		return user, err
	}
	return users, err
}
