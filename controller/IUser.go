package controller

import "github.com/gin-gonic/gin"

type IUserController interface {
	Insert(*gin.Context)
	GetOne(*gin.Context)
	GetAll(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
