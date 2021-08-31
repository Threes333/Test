package Model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

type T struct {
	gorm.Model
}

func InitDB() {
	dsn := "root:qazpl.123456@tcp(127.0.0.1:3306)/now?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := DB.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	//将模型与数据库连接
	if err := DB.AutoMigrate(User{}, Posts{}, Category{}); err != nil {
		panic(err)
	}
}
