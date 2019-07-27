package datastore

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kntaka/go-sample-api/domain/model"
	"os"
)

func NewMySqlDB() *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open("mysql", connectionString)
	if nil != err {
		panic(err)
	}

	// 応答確認
	err = db.DB().Ping()
	if nil != err {
		panic(err)
	}
	// Out put the detail of the sql log
	db.LogMode(true)

	// Set DB Engine
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	// transaction
	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.UserActive{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(model.UserDetail{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(model.UserLeave{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
