package models

import (
	orm "Week04/global"
)

// Plan 等级
type Plan struct {
	ID     int64  `json:"id"        gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `json:"name"      gorm:"type:varchar(255);comment:'名称'"`
	Amount int    `json:"amount"    gorm:"type:int;comment:'价格'"`
	Token  int    `json:"token"     gorm:"type:int;comment:'积分'"`

	CreateBy  string `json:"createBy"    gorm:"type:varchar(128);"`
	UpdateBy  string `json:"updateBy"    gorm:"type:varchar(128);"`
	Remark    string `json:"remark"      gorm:"type:varchar(255);"`
	Admin     bool   `json:"admin"       gorm:"type:char(1);"`
	DataScope string `json:"dataScope"   gorm:"type:varchar(128);"`
	Params    string `json:"params"      gorm:"-"`

	BaseModel
}

func (Plan) TableName() string {
	return "sys_plan"
}

func (l *Plan) GetPage(pageSize int, pageIndex int) ([]Plan, int, error) {
	var doc []Plan

	table := orm.Eloquent.Select("*").Table(l.TableName())
	if l.ID != 0 {
		table = table.Where("id = ?", l.ID)
	}

	var count int
	if err := table.Order("id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (l *Plan) Get() (plan Plan, err error) {
	table := orm.Eloquent.Table(l.TableName())
	if l.ID != 0 {
		table = table.Where("id = ?", l.ID)
	}

	if err = table.First(&plan).Error; err != nil {
		return
	}
	return
}

func (l *Plan) GetList() (plans []Plan, err error) {
	table := orm.Eloquent.Table(l.TableName())
	if err = table.Order("id").Find(&plans).Error; err != nil {
		return
	}
	return
}

func (l *Plan) Insert() (id int64, err error) {
	l.UpdateBy = ""
	result := orm.Eloquent.Table(l.TableName()).Create(&l)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = l.ID
	return
}

// Update 修改
func (l *Plan) Update(id int64) (update Plan, err error) {
	if err = orm.Eloquent.Table(l.TableName()).First(&update, id).Error; err != nil {
		return
	}
	// 参数1: 旧数据
	// 参数2: 新数据
	if err = orm.Eloquent.Table(l.TableName()).Model(&update).Updates(&l).Error; err != nil {
		return
	}
	return
}

// Delete 删除
func (l *Plan) Delete(id int64) (success bool, err error) {
	if err = orm.Eloquent.Table(l.TableName()).Where("id = ?", id).Delete(&Plan{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
