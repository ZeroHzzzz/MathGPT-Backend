package mysql

import (
	"fmt"
	"log"
	"mathgpt/configs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() { // 初始化数据库

	user := config.Config.GetString("mysql.user")
	pass := config.Config.GetString("mysql.pass")
	port := config.Config.GetString("mysql.port")
	host := config.Config.GetString("mysql.host")
	name := config.Config.GetString("mysql.name")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
	})

	if err != nil {
		log.Fatal("DatabaseConnectFailed", err)
	}

	err = autoMigrate(db)
	if err != nil {
		log.Fatal("DatabaseMigrateFailed", err)
	}

	DB = db
}
