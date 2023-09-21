package orm

import (
	// "database/sql"

	"gin/config"

	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func GormInitDB() {

	dsn := config.DbUser + ":" + config.DbPassWord + "@tcp(" + config.DbHost + ":" + config.DBPort + ")/" + config.DbName

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //打开数据库连接

	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.DB()
	if err != nil {
		// 处理错误
		log.Fatal("Failed to get database connection:", err)
		return
	}

	err = dbConn.Ping()
	if err != nil {
		// 处理连接错误
		log.Fatal("Failed to ping database:", err)
		return
	}

	DB = db
}

func GormCloseDB() {
	// 关闭数据库连接
	if DB != nil {
		dbConn, err := DB.DB()
		if err != nil {
			// 处理错误
			log.Fatal("Failed to get database connection:", err)
			return
		}
		dbConn.Close()
	}
}
