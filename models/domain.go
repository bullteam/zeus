package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/bullteam/zeus/utils"
	"github.com/bullteam/zeus/dto"
)

type DomainaddForm struct {
	Name  string  `form:"name"`
	Callbackurl  string `form:"callbackurl"`
	Remark  string `form:"remark"`
	Code string `form:"code"`
}

type Domain struct {
	Id           int `json:"id"`
	Name         string `json:"name"`
	Callbackurl  string `json:"callbackurl"`
	Remark       string `json:"remark"`
	Code 			string `json:"code"`
	Create_time  time.Time `json:"created_time"`
	Last_update_time time.Time `json:"updated_time"`
}

func init()  {
	orm.RegisterModel(new(Domain))
}

func NewDomain(domainform *DomainaddForm) (d *Domain,err error){
	domain := Domain{
		Name: domainform.Name,
		Callbackurl: domainform.Callbackurl,
		Remark: domainform.Remark,
		Code:domainform.Code,
	}
	return &domain,nil
}
func (d *Domain) Insert()  {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into domain(name,callbackurl,remark,code) values (?,?,?,?)",d.Name,d.Callbackurl,d.Remark,d.Code).Exec()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
}
func (d Domain) List(start int, limit int, q []string)([]*Domain,int64){
	o := orm.NewOrm()
	var dm []*Domain
	qs := o.QueryTable("domain")
	//qs.RelatedSel("domain","domain_id","id")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.DOMAIN_SEARCH) {
			qs = utils.TransformQset(qs,k,v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	qs.Limit(limit, start).All(&dm)
	c, _ := qs.Count()
	return dm, c
}
func Domain_list(page int,offset int) (domainlist []*Domain,cnt int64){
	var Domains []*Domain
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("domain")
	counts, _ := qs.Count()
	start := (page - 1) * offset
	qs.Limit(offset, start).All(&Domains)
	return Domains,counts
}


//修改
func UpdateDomain(id int, name string,callbackurl string, remark string) error {
	o := orm.NewOrm()
	domain := &Domain{Id: id}
	if o.Read(domain) == nil {
		domain.Name = name
		domain.Callbackurl = callbackurl
		domain.Remark = remark
		//domain.Last_update_time = time.Now()
		_, err := o.Update(domain,"Name","Callbackurl","Remark")
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteDomain(id int) error {
	o := orm.NewOrm()
	domain := &Domain{Id: id}
	if o.Read(domain) == nil {
		_, err := o.Delete(domain)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据id取得域
func GetDomain(id int) (Domains *Domain, err error) {
	o := orm.NewOrm()
	domain := new(Domain)
	qs := o.QueryTable("domain")
	err = qs.Filter("id", id).One(domain)
	if err != nil {
		return Domains, err
	}
	return domain, err
}

func (d Domain)GetDomainByCode(code string) (Domains *Domain, err error){
	o := orm.NewOrm()
	domain := new(Domain)
	qs := o.QueryTable("domain")
	err = qs.Filter("code", code).One(domain)
	if err != nil {
		return Domains, err
	}
	return domain, err
}