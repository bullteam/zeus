package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"zeus/pkg/utils"
	"zeus/pkg/dto"
)

type DeptaddForm struct {
	Name  string  `form:"name"`
	Parent_id  int `form:"parent_id"`
	Order_num int `formn:"order_num"`
}

type Department struct {
	Id           int `json:"id"`
	Name         string `json:"name"`
	Order_num	int `json:"order_num"`
	Parent_id  int `json:"parent_id"`
	//Create_time  time.Time `json:"create_time"`
	//Update_time time.Time `json:"update_time"`
}

func init()  {
	orm.RegisterModel(new(Department))
}

//func NewDept(deptform *DeptaddForm) (d *Department,err error){
//	dept := Department{
//		Name: deptform.Name,
//		Parent_id: deptform.Parent_id,
//	}
//	return &dept,nil
//}
func (d *Department) Insert()  (int64,error){
	o := orm.NewOrm()
	o.Using("default")
	return o.Insert(d)
	//_, err := o.Raw("insert into department(name,parent_id,order_num) values (?,?,?)",d.Name,d.Parent_id,d.Order_Num).Exec()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("insert ok")
	//}
}
func (d *Department) Update()  (int64,error){
	o := orm.NewOrm()
	o.Using("default")
	return o.Update(d)
}
func (d Department) List(start int, limit int, q []string)([]*Department,int64){
	o := orm.NewOrm()
	var ds []*Department
	qs := o.QueryTable("department").OrderBy("order_num")
	//qs.RelatedSel("domain","domain_id","id")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.DEPARTMENT_SEARCH) {
			qs = utils.TransformQset(qs,k,v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	qs.Limit(limit, start).RelatedSel().All(&ds)
	c, _ := qs.Count()
	return ds, c
}
func Dept_list(page int,offset int) (deptlist []*Department,cnt int64){
	var Depts []*Department
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("department")
	counts, _ := qs.Count()
	start := (page - 1) * offset
	qs.Limit(offset, start).All(&Depts)
	return Depts,counts
}


//修改
func UpdateDept(id int, name string,parent_id int) error {
	o := orm.NewOrm()
	dept := &Department{Id: id}
	if o.Read(dept) == nil {
		dept.Name = name
		dept.Parent_id = parent_id
		_, err := o.Update(dept)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteDept(id int) error {
	o := orm.NewOrm()
	dept := &Department{Id: id}
	if o.Read(dept) == nil {
		_, err := o.Delete(dept)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDept(id int) (Depts *Department, err error) {
	o := orm.NewOrm()
	dept := new(Department)
	qs := o.QueryTable("department")
	err = qs.Filter("id", id).One(dept)
	if err != nil {
		return Depts, err
	}
	return dept, err
}
func (d Department) GetDeptByName(name string) (Department, error) {
	o := orm.NewOrm()
	o.Using("default")
	var department Department
	err := o.Raw("select * from department where name = ?", name).QueryRow(&department)
	return department,err
}

func (d Department) GetDeptByOtherName(id int,name string) (Department, error) {
	o := orm.NewOrm()
	o.Using("default")
	var department Department
	err := o.Raw("select * from department where name = ? and id <> ?", name,id).QueryRow(&department)
	return department,err
}