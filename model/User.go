package model

import (
	"example.com/v2/dao"
	"fmt"
)

type User struct {
	Id       int64  `gorm:"column:id;primary_key"` // 用户ID
	Username string `form:"username" gorm:"column:username"`
	Nickname string `gorm:"column:nickname"`
}

func (User) TableName() string {
	return "nt24_admin"
}

// 详情
func GetUserTest(id int) (*User, error) {

	user := new(User)
	fmt.Println(user)
	err := dao.Db.Where("id = ?", id).Find(user).Error
	fmt.Println(user)
	return user, err
}

// 添加
func CreateUsetTest(username string) (int64, error) {
	user := User{Username: username}

	fmt.Println(dao.Db)
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

// 列表
func GetList() ([]User, error) {
	var user []User
	err := dao.Db.Find(&user).Error
	return user, err
}

// 编辑
func EditUserTest(id int, username string) {
	dao.Db.Model(&User{}).Where("id = ?", id).Update("nickname", username)
}

// 删除
func DelUserTest(id int) error {
	var user User
	// DELETE from emails where id=20;
	//
	//dao.Db.Delete(&user, id)

	err := dao.Db.Where("id = ?", id).Delete(&user).Error
	return err
}
