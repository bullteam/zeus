package dao

import (
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/utils"
)

type DomainDao struct {
}

func (dao *DomainDao) NewDomain(dto *dto.DomainAddDto) (d *models.Domain, err error) {
	domain := models.Domain{
		Name:        dto.Name,
		Callbackurl: dto.Callbackurl,
		Remark:      dto.Remark,
		Code:        dto.Code,
	}

	return &domain, nil
}

func (dao *DomainDao) Insert(dto *dto.DomainAddDto) (int64, error) {
	o := GetOrmer()
	var domain models.Domain
	domain.Name = dto.Name
	domain.Callbackurl = dto.Callbackurl
	domain.Remark = dto.Remark
	domain.Code = dto.Code
	id, err := o.Insert(&domain)

	return id, err
}

func (dao *DomainDao) List(start int, limit int, q []string) ([]*models.Domain, int64) {
	o := GetOrmer()
	var dm []*models.Domain
	qs := o.QueryTable("domain")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.DOMAIN_SEARCH) {
			qs = utils.TransformQset(qs, k, v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	_, err := qs.Limit(limit, start).All(&dm)
	if err != nil {
		return dm, 0
	}
	total, _ := qs.Count()

	return dm, total
}

//修改
func (dao *DomainDao) Update(id int, name string, callbackurl string, remark string) error {
	o := GetOrmer()
	domain := &models.Domain{Id: id}
	if o.Read(domain) == nil {
		domain.Name = name
		domain.Callbackurl = callbackurl
		domain.Remark = remark
		_, err := o.Update(domain, "Name", "Callbackurl", "Remark")
		if err != nil {
			return err
		}
	}

	return nil
}

//删除
func (dao *DomainDao) Delete(id int) error {
	o := GetOrmer()
	domain := &models.Domain{Id: id}
	if o.Read(domain) == nil {
		_, err := o.Delete(domain)
		if err != nil {
			return err
		}
	}

	return nil
}

//根据id取得域
func (dao *DomainDao) GetDomain(id int) (models.Domain, error) {
	o := GetOrmer()
	domain := models.Domain{
		Id: id,
	}
	err := o.Read(&domain)

	return domain, err
}

func (dao *DomainDao) GetDomainByCode(code string) (domain models.Domain, err error) {
	o := GetOrmer()
	err = o.QueryTable(&models.Domain{}).Filter("code", code).One(&domain)

	return domain, err
}
