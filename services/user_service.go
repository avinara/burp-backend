package services

import (
	"database/sql"

	"github.com/burp-backend/errors"
	"github.com/burp-backend/model"
)

type UserServiceAPI interface {
	GetAllUsers() ([]model.User, errors.ErrorInterface)
	GetUserByName(name string) (*model.User, errors.ErrorInterface)
	CreateUser(User *model.User) (int, errors.ErrorInterface)
	UpdateUser(User *model.User) errors.ErrorInterface
	DeleteUser(id int) errors.ErrorInterface
}

type UserService struct {
	DB *sql.DB
}

func NewUserService(database *sql.DB) UserServiceAPI {
	return &UserService{
		DB: database,
	}
}

func (c *UserService) GetAllUsers() ([]model.User, errors.ErrorInterface) {
	rows, err := c.DB.Query("SELECT id, name FROM users")
	if err != nil {
		err2 := errors.DatabaseQueryError()
		return nil, err2
	}
	defer rows.Close()

	var Users []model.User
	for rows.Next() {
		var User model.User
		if err := rows.Scan(&User.ID, &User.Name); err != nil {
			err2 := errors.ScanningRowsError()
			return nil, err2
		}
		Users = append(Users, User)
	}
	return Users, nil
}

func (c *UserService) GetUserByName(name string) (*model.User, errors.ErrorInterface) {
	var User model.User
	err := c.DB.QueryRow("SELECT id, name FROM Users WHERE name=?", name).
		Scan(&User.ID, &User.Name)
	if err != nil {
		err2 := errors.DatabaseQueryError()
		return nil, err2
	}
	return &User, nil
}

func (c *UserService) CreateUser(User *model.User) (int, errors.ErrorInterface) {
	result, err := c.DB.Exec("INSERT INTO Users(name) VALUES (?)", User.Name)
	if err != nil {
		err2 := errors.DatabaseInsertionError()
		return 0, err2
	}
	id, err := result.LastInsertId()
	if err != nil {
		err2 := errors.DatabaseQueryError()
		return 0, err2
	}
	return int(id), nil
}

func (c *UserService) UpdateUser(User *model.User) errors.ErrorInterface {
	_, err := c.DB.Exec("UPDATE Users SET name=? WHERE id=?", User.Name, User.ID)
	if err != nil {
		err2 := errors.DatabaseUpdationError()
		return err2
	}
	return nil
}

func (c *UserService) DeleteUser(id int) errors.ErrorInterface {
	_, err := c.DB.Exec("DELETE FROM Users WHERE id=?", id)
	if err != nil {
		err2 := errors.DatabaseDeletionError()
		return err2
	}
	return nil
}
