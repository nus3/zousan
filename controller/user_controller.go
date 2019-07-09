package controller

import (
	"github.com/gin-gonic/gin"

	user "github.com/TechLoCo/Zousan-api/service"
)

// User is user controller
type User struct{}

// Index action: GET /users
func (u User) Index(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (u User) Create(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(201, p)
	}
}
