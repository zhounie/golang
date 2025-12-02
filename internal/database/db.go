package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(models ...interface{}) {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/ikun_coffee?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接 mysql: %v", err)
	}
	log.Println("mysql 连接成功")

	err = DB.AutoMigrate(models...)

	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	log.Println("数据库迁移成功")
}