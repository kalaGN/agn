// Package factories 存放工厂方法
package factories

import (
	"agn/app/models/dial_record"
	"agn/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeDials(times int) []dial_record.Bd_dial_Record {

	var objs []dial_record.Bd_dial_Record

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := dial_record.Bd_dial_Record{

			Enterprise_id: "2",
			Manager_id:    helpers.RandomNumber(8),
			Task_identify: faker.CCNumber(),
			Phone_id:      faker.CCNumber(),
			Phoneno:       helpers.RandomNumber(11),
			Dialresult:    "1",
			Desire:        "1",
			Dialtime:      faker.TimeString(),
			Dialendtime:   faker.TimeString(),
			Talkbegintime: faker.TimeString(),
			Talkendtime:   faker.TimeString(),
			Invokeid:      faker.CCNumber(),
			Trunkgroupid:  faker.CCNumber(),
			Dialtimes:     faker.CCNumber(),
			Ib_timestamp:  faker.Timestamp(),
			Talktime:      faker.TimeString(),
		}
		objs = append(objs, model)
	}

	return objs
}
