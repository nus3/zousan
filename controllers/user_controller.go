package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/TechLoCo/Zousan-api/services"
)

// UserController is user controller
type UserController struct{}

// Index action: GET /users
func (u UserController) Index(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (u UserController) Create(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(201, p)
	}
}
