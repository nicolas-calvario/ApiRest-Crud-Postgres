package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/controller"
	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/db"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.ConnectionDB()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewUserController(a.DB)
	r.POST("/user", controller.Insert)
	r.PUT("/user", controller.Update)
	r.DELETE("/:id", controller.Delete)
	r.GET("/:id", controller.GetOne)
	r.GET("/users", controller.GetAll)
	r.Run(":8080")
}
