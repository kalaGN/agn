// Package user 存放用户 Model 相关逻辑
package dial_record

import (
	"agn/app/models"
	"agn/pkg/database"
)

/**
  `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `enterprise_id` varchar(64) NOT NULL DEFAULT '0' COMMENT '企业ID',
  `manager_id` varchar(64) NOT NULL,
  `task_identify` varchar(64) NOT NULL DEFAULT '' COMMENT '任务ID',
  `phone_id` int(11) DEFAULT '0' COMMENT '被叫人ID',
  `PhoneNo` varchar(25) NOT NULL DEFAULT '' COMMENT '号码',
  `DialResult` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '拨号结果，1为成功，其他为失败，NULL无结果',
  `desire` varchar(64) NOT NULL DEFAULT '0' COMMENT '意愿',
  `DialTime` datetime DEFAULT NULL COMMENT '拨号时间',
  `DialEndTime` datetime DEFAULT NULL COMMENT '拨号结束时间',
  `TalkBeginTime` datetime DEFAULT NULL COMMENT '通话开始时间',
  `TalkEndTime` datetime DEFAULT NULL COMMENT '完成时间（拨号未成功的忽略）',
  `InvokeId` varchar(64) DEFAULT '' COMMENT '呼叫的标识，与平台那边的呼叫记录结合',
  `TrunkGroupId` int(11) DEFAULT NULL,
  `DialTimes` int(11) DEFAULT NULL COMMENT '拨打次数',
  `ib_timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '时间戳',
  `TalkTime` int(11) DEFAULT NULL,
*/

// User 用户模型
type Bd_dial_Record struct {
	models.BaseModel

	Enterprise_id string `gorm:"type:int(11)"`
	Manager_id    string `gorm:"type:varchar(64)"`
	Task_identify string `gorm:"type:varchar(64)"`
	Phone_id      string `gorm:"type:int(11)"`
	Phoneno       string `gorm:"type:varchar(64)"`
	Dialresult    string `gorm:"type:int(11);unsigned`
	Desire        string `gorm:"type:varchar(64)"`
	Dialtime      string `gorm:"type:datetime"`
	Dialendtime   string `gorm:"type:datetime"`
	Talkbegintime string `gorm:"type:datetime"`
	Talkendtime   string `gorm:"type:datetime"`
	Invokeid      string `gorm:"type:varchar(64)"`
	Trunkgroupid  string `gorm:"type:int(11)"`
	Dialtimes     string `gorm:"type:datetime"`
	Ib_timestamp  string `gorm:"type:timestamp"`
	Talktime      string `gorm:"type:int(11)"`
}

// IsPhoneExist 判断手机号已被注册
func IsDataExist(enterprise_id string) (u []Bd_dial_Record) {
	database.DB.Model(Bd_dial_Record{}).Where("enterprise_id = ?", enterprise_id).Find(&u)
	return u
}
