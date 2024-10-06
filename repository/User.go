package repository

import (
	"database/sql"
	"time"

	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/models"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositpryInterface {
	return &UserRepository{Db: db}
}

func (u *UserRepository) Delete(id uint64) (bool, error) {
	_, err := u.Db.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Update implements UserRepositpryInterface.
func (u *UserRepository) Update(user models.User) (models.User, error) {
	_, err := u.Db.Exec("Update Users SET name = $1, email = $2, birthday = $3  WHERE id = $4", user.Name, user.Email, user.BirthDay, user.Id)
	if err != nil {
		return models.User{}, err
	}
	return u.GetOne(user.Id)
}

// Insert implements UserRepositpryInterface.
func (u *UserRepository) Insert(post models.User) error {
	stmt, err := u.Db.Prepare("Insert Into Users(name, email, birthday)Values ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.Name, post.Email, post.BirthDay)
	if err != nil {
		return err
	}
	return nil
}

// GetOne implements UserRepositpryInterface.
func (u *UserRepository) GetOne(id uint64) (models.User, error) {
	var user models.User
	query, err := u.Db.Query("SELECT * FROM Users WHERE id = $1", id)
	if err != nil {
		return user, err
	}
	if query != nil {
		for query.Next() {
			var (
				id       uint64
				name     string
				email    string
				birthday time.Time
			)
			err := query.Scan(&id, &name, &email, &birthday)
			if err != nil {
				return user, err
			}
			user = models.User{Id: id, Name: name, Email: email, BirthDay: birthday}
		}
	}
	return user, nil
}

// GetAll implements UserRepositpryInterface.
func (u *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	query, err := u.Db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	if query != nil {
		for query.Next() {
			var (
				id       uint64
				name     string
				email    string
				birthday time.Time
			)
			err := query.Scan(&id, &name, &email, &birthday)
			if err != nil {
				return nil, err
			}
			user := models.User{Id: id, Name: name, Email: email, BirthDay: birthday}
			users = append(users, user)

		}
	}
	return users, nil
}
