package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/TechLoCo/Zousan-api/db"
	"github.com/TechLoCo/Zousan-api/entity"
)

// Service is user's behavior
type Service struct{}

// User is user struct
type User entity.User

// DB is models
var DB *gorm.DB = db.GetDB()

// GetAll returns all user.
func (s Service) GetAll() ([]User, error) {
	var u []User

	// TODO: エラー処理
	if err := DB.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create User model
func (s Service) CreateModel(c *gin.Context) (User, error) {
	var u User

	// TODO: バリデーション
	// TODO: エラー処理
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := DB.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
