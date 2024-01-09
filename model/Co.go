package model

import (
	"gorm.io/plugin/soft_delete"
)

type Co struct {
	ID        int                   `json:"id" gorm:"primarykey"`
	CreatedAt int                   `json:"create_at" gorm:"column:create_at"`
	UpdatedAt int                   `json:"update_at" gorm:"column:update_at"`
	DelAt     soft_delete.DeletedAt `json:"del_at" gorm:"softDelete:milli"`
	Pid       int                   `json:"pid" gorm:"column:pid"`
	CountryId int                   `json:"country_id" gorm:"column:country_id"`
	CotypeId  uint                  `json:"cotype_id" gorm:"column:cotype_id;"`
	AdminId   int                   `json:"admin_id" gorm:"column:admin_id"`
	Name      string                `json:"name" gorm:"column:name"`
	Tn        string                `json:"tn" gorm:"column:tn"`
	Status    int                   `json:"status" gorm:"column:status"`
	Level     int                   `json:"level" gorm:"column:level"`
	Point     int                   `json:"point" gorm:"column:point"`
	Email     string                `json:"email" gorm:"column:email"`
	Mobile    string                `json:"mobile" gorm:"column:mobile"`
	Website   string                `json:"website" gorm:"column:website"`
	Cotype    *Cotypelian           `gorm:"foreignKey:cotype_id;AssociationForeignKey:id;field:'name'"`
}
type Cotypelian struct {
	Id         int    `json:"-"`
	Country_id int    `json:"-"`
	Drank      int    `json:"-"`
	create_id  int    `json:"-"`
	Name       string `json:"name"`
}

func (Cotypelian) TableName() string {
	return "nt24_cotype"
}
func (Co) TableName() string {
	return "nt24_co"
}
