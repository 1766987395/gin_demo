package db

import (
	"database/sql"

	"log"

	"gin/config"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
)

var DB *sql.DB

func InitDB() {
	// 连接数据库
	db, err := sql.Open(config.Db, config.DbUser+":"+config.DbPassWord+"@tcp("+config.DbHost+":"+config.DBPort+")/"+config.DbName)

	if err != nil {
		log.Fatal(err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func CloseDB() {
	// 关闭数据库连接
	if DB != nil {
		DB.Close()
	}
}
