package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/TechLoCo/Zousan-api/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init actions initialize db
func Init() *gorm.DB {
	// TODO: 環境変数によってenvファイルを変更する
	err := godotenv.Load(fmt.Sprintf("env/local.env"))
	if err != nil {
		panic(err.Error())
	}

	dbMs := os.Getenv("DB_CONNECTION")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	protocol := "tcp(" + host + ":" + port + ")"
	dbName := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	connect := user + ":" + pass + "@" + protocol + "/" + dbName

	db, err = gorm.Open(dbMs, connect)
	if err != nil {
		// TODO: エラー処理 panic使わない
		panic(err.Error())
		// panic(err.Error())
	}

	autoMigration()

	return db
}

// Close actions disconnect database.
func Close() {
	db.Close()
}

// GetDB returns models
func GetDB() *gorm.DB {
	return db
}

func autoMigration() {
	// TODO: マイグレーションの管理
	// 		 カラムの変更(追加や削除)をどのように対応するか
	db.AutoMigrate(&entity.User{})

	// NOTE: 必要なテーブルを追加していく
}
