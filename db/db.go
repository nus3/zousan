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

// MySQLConfig is config.
type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (c *MySQLConfig) toString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.User, c.Password, c.Host, c.Port, c.Database)
}

// Init actions initialize db
func Init() *gorm.DB {
	// TODO: 環境ごとに読み込むenvファイルを変更する
	err := godotenv.Load(fmt.Sprintf("env/local.env"))
	if err != nil {
		panic(err.Error())
	}

	// HACK: MySQLConfigつってんのにdbMsを指定するのはおかしいので直したいが頭が働かないので元気な時になおす
	dbMs := os.Getenv("DB_CONNECTION")
	mysqlConfig := &MySQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	db, err = gorm.Open(dbMs, mysqlConfig.toString())
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
