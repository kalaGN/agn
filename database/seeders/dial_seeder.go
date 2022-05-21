package seeders

import (
	"agn/database/factories"
	"agn/pkg/console"
	"agn/pkg/logger"
	"agn/pkg/seed"
	"fmt"

	"gorm.io/gorm"
)

func init() {

	// 添加 Seeder
	seed.Add("SeedDialTable", func(db *gorm.DB) {

		// 创建 10 个用户对象
		users := factories.MakeDials(10)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("bd_dial_record").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
