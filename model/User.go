package model

import (
	"gorm.io/plugin/soft_delete"
)

type UserModel struct {
	Id           int                   `json:"id" gorm:"column:id;primary_key"` // 用户ID
	Username     string                `json:"username" gorm:"column:username"`
	Nickname     string                `json:"nickname" gorm:"column:nickname"`
	Salt         string                `json:"-" gorm:"column:salt"`
	Email        string                `json:"email" gorm:"column:email"`
	Avatar       string                `json:"avatar" gorm:"column:avatar"`
	Mobile       string                `json:"mobile" gorm:"column:mobile"`
	CompanyId    int                   `json:"company_id" gorm:"column:company_id"`
	Level        int                   `json:"level" gorm:"column:level"`
	Loginfailure int                   `json:"-" gorm:"column:loginfailure"`
	Login_at     int                   `json:"login_at" gorm:"column:login_at"`
	IsBoss       int                   `json:"is_boss" gorm:"column:is_boss"`
	Loginip      string                `json:"-" gorm:"column:loginip"`
	Status       int                   `json:"status" gorm:"column:status"`
	CreateId     int                   `json:"-" gorm:"column:create_id"`
	UpdateId     int                   `json:"-" gorm:"column:update_id"`
	CreatedAt    int                   `json:"create_at" gorm:"column:create_at"`
	UpdatedAt    int                   `json:"update_at" gorm:"column:update_at"`
	DelAt        soft_delete.DeletedAt `json:"del_at" gorm:"softDelete:milli"`
	DelId        int                   `json:"-" gorm:"column:del_id"`
	Password     string                ` json:"-" form:"password" gorm:"column:password"`
	//Country  *CountryModel `json:"country" gorm:"foreignKey:country_id;AssociationForeignKey:id;field:'name'"`
}

func (UserModel) TableName() string {
	return "ca_user"
}
