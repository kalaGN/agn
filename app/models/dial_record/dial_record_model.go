// Package dial_record 存放拨打记录 Model 相关逻辑
package dial_record

import (
	"agn/app/models"
	"agn/pkg/app"
	"agn/pkg/database"
	"agn/pkg/paginator"

	"github.com/gin-gonic/gin"
)

// dial result 模型
type Bd_dial_Record struct {
	models.BaseModel

	Enterprise_id string `gorm:"type:int(11)" json:"enterprise_id"`
	Manager_id    string `gorm:"type:varchar(64)" json:"manager_id"`
	Task_identify string `gorm:"type:varchar(64)" json:"task_identify"`
	Phone_id      string `gorm:"type:int(11);index:calledid_index" json:"phone_id"`
	Phoneno       string `gorm:"type:varchar(64)" json:"PhoneNo"`
	Dialresult    string `gorm:"type:int(11);unsigned"  json:"DialResult"`
	Desire        string `gorm:"type:varchar(64)" json:"desire"`
	Dialtime      string `gorm:"type:datetime"  json:"DialTime"`
	Dialendtime   string `gorm:"type:datetime"  json:"DialEndTime"`
	Talkbegintime string `gorm:"type:datetime" json:"TalkBeginTime"`
	Talkendtime   string `gorm:"type:datetime" json:"TalkEndTime"`
	Invokeid      string `gorm:"type:varchar(64)"  json:"InvokeId"`
	Trunkgroupid  string `gorm:"type:int(11)"  json:"TrunkGroupId"`
	Dialtimes     string `gorm:"type:datetime"  json:"DialTimes"`
	Ib_timestamp  string `gorm:"type:timestamp" json:"ib_timestamp"`
	Talktime      string `gorm:"type:int(11)"  json:"TalkTime"`
}

// IsPhoneExist 判断手机号已被注册
func IsDataExist(enterprise_id string) (u []Bd_dial_Record) {
	database.DB.Model(Bd_dial_Record{}).Where("enterprise_id = ?", enterprise_id).Find(&u)
	return u
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []Bd_dial_Record, paging paginator.Paging) {

	comp_id := c.DefaultQuery("comp_id", "0")
	task_identify := c.Query("task_identify")
	paging = paginator.Paginate(
		c,
		database.DB.Model(Bd_dial_Record{}).Select([]string{"task_identify", "DialResult", "PhoneNo", "DialTime", "InvokeId"}).Where("enterprise_id=?", comp_id).Where("task_identify=?", task_identify),
		&users,
		app.V1URL(""),
		perPage,
	)
	return
}
