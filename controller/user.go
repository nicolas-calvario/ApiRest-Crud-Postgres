package controller

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/models"
	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/repository"
	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/utils"
)

type UserController struct {
	Db *sql.DB
}

func (u *UserController) GetAll(c *gin.Context) {
	repository := repository.NewUserRepository(u.Db)
	users, err := repository.GetAll()
	if err == nil {
		c.JSON(200, gin.H{"status": "success", "data": users, "msg": "get user successfully"})
	} else {
		c.JSON(500, gin.H{"status": "Failed", "data": users, "msg": err.Error()})
	}
}

func (u *UserController) Update(c *gin.Context) {
	repository := repository.NewUserRepository(u.Db)
	var update models.User
	if err := c.ShouldBind(&update); err != nil {
		c.JSON(400, gin.H{"status": "Failed", "msg": err})
		return
	}
	users, err := repository.Update(update)
	if err == nil {
		c.JSON(200, gin.H{"status": "success", "data": users, "msg": "Update user successfully"})
	} else {
		c.JSON(500, gin.H{"status": "Failed", "data": users, "msg": err.Error()})

	}
}

func (u *UserController) Delete(c *gin.Context) {
	repository := repository.NewUserRepository(u.Db)
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	status, err := repository.Delete(id)
	if status && err == nil {
		c.JSON(200, gin.H{"status": "success", "msg": "User delete ..."})
	} else {
		c.JSON(500, gin.H{"status": "Failed", "msg": err.Error()})

	}
}

func (u *UserController) GetOne(c *gin.Context) {
	repository := repository.NewUserRepository(u.Db)
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user, err := repository.GetOne(id)
	if (user != models.User{} && err == nil) {
		c.JSON(200, gin.H{"status": "success", "data": user, "msg": "get user successfully"})
	} else {
		c.JSON(500, gin.H{"status": "Failed", "data": user, "msg": err.Error()})

	}
}

func (u *UserController) Insert(c *gin.Context) {
	var post models.User
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "Failed", "msg": err})
		return
	}

	errMail := utils.ValidateEmail(post.Email)
	if errMail != nil {
		c.JSON(500, gin.H{"status": "failed", "msg": errMail.Error()})
		return
	}

	repo := repository.NewUserRepository(u.Db)
	errInsert := repo.Insert(post)

	if errInsert == nil {
		c.JSON(200, gin.H{"status": "success", "msg": "insert user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": errInsert})
		return
	}
}

func NewUserController(db *sql.DB) IUserController {
	return &UserController{Db: db}
}
