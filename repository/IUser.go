package repository

import "github.com/nicolas-calvario/ApiRest-Crud-Postgres/models"

type UserRepositpryInterface interface {
	Insert(models.User) error
	GetAll() ([]models.User, error)
	GetOne(uint64) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(uint64) (bool, error)
}
