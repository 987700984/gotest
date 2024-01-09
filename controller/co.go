package controller

import (
	"example.com/v2/dao"
	"example.com/v2/model"
	"github.com/gin-gonic/gin"
	"gorm.io/plugin/soft_delete"
)

type CoController struct{}
type Co struct {
	ID          int                   `json:"id" gorm:"primarykey"`
	CreatedAt   int                   `json:"create_at" `
	UpdatedAt   int                   `json:"update_at" gorm:"column:update_at"`
	DelAt       soft_delete.DeletedAt `json:"del_at" gorm:"softDelete:milli"`
	Pid         int                   `json:"pid" gorm:"column:pid"`
	CountryId   int                   `json:"country_id" gorm:"column:country_id"`
	CotypeId    uint                  `json:"cotype_id" gorm:"column:cotype_id;"`
	AdminId     int                   `json:"admin_id" gorm:"column:admin_id"`
	Name        string                `json:"name" gorm:"column:name"`
	Tn          string                `json:"tn" gorm:"column:tn"`
	Status      int                   `json:"status" gorm:"column:status"`
	Level       int                   `json:"level" gorm:"column:level"`
	Point       int                   `json:"point" gorm:"column:point"`
	Email       string                `json:"email" gorm:"column:email"`
	Mobile      string                `json:"mobile" gorm:"column:mobile"`
	Website     string                `json:"website" gorm:"column:website"`
	cotype_name string
}
type Groupco struct {
	Country_id int
	Count      int
}

func (CoController) Test(c *gin.Context) {

	//一对一关联查询
	var co []model.Co
	//Preload方式
	var codemo Co
	dao.Db.Debug().Preload("Cotype").First(&co, "id =?", 5).Scan(&codemo)
	//var cog []Groupco
	//dao.Db.Debug().Select("country_id, count(*) as count").Group("country_id").Find(&co).Having("country_id >1").Scan(&cog)
	//fmt.Println(co)
	//Association  费劲
	//dao.Db.First(&co, "id =?", 5)
	//dao.Db.Model(&co).Association("Cotype").Find(&co.Cotype)
	//dao.Db.Debug().Take(&co, 1)
	Common{}.ReturnSuccess(c, 0, "获取成功", co, 1)
}
