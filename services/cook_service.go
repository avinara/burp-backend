package services

import (
	"database/sql"

	"github.com/burp-backend/errors"
	"github.com/burp-backend/model"
)

type CookServiceAPI interface {
	GetAllCooks() ([]model.Cook, errors.ErrorInterface)
	GetCookByEmail(email string) (*model.Cook, errors.ErrorInterface)
	CreateCook(Cook *model.Cook) (bool, errors.ErrorInterface)
	UpdateCook(Cook *model.Cook) errors.ErrorInterface
	DeleteCook(email string) errors.ErrorInterface
}

type CookService struct {
	DB *sql.DB
}

func NewCookService(database *sql.DB) CookServiceAPI {
	return &CookService{
		DB: database,
	}
}

func (s *CookService) GetAllCooks() ([]model.Cook, errors.ErrorInterface) {
	rows, err := s.DB.Query("SELECT id, name, email, age FROM cooks")
	if err != nil {
		err2 := errors.DatabaseQueryError()
		return nil, err2
	}
	defer rows.Close()

	var Cooks []model.Cook
	for rows.Next() {
		var Cook model.Cook
		if err := rows.Scan(&Cook.ID, &Cook.Name, &Cook.Email, &Cook.Age); err != nil {
			err2 := errors.ScanningRowsError()
			return nil, err2
		}
		Cooks = append(Cooks, Cook)
	}
	return Cooks, nil
}

func (s *CookService) GetCookByEmail(email string) (*model.Cook, errors.ErrorInterface) {
	var Cook model.Cook
	err := s.DB.QueryRow("SELECT id, name, email, age FROM cooks WHERE email=?", email).
		Scan(&Cook.ID, &Cook.Name, &Cook.Email, &Cook.Age)
	if err != nil {
		err2 := errors.DatabaseQueryError()
		return nil, err2
	}
	return &Cook, nil
}

func (s *CookService) CreateCook(Cook *model.Cook) (bool, errors.ErrorInterface) {
	_, err := s.DB.Exec("INSERT INTO cooks(name, email, age) VALUES (?, ?, ?)", Cook.Name, Cook.Email, Cook.Age)
	if err != nil {
		err2 := errors.DatabaseInsertionError()
		return false, err2
	}
	return true, nil
}

func (s *CookService) UpdateCook(Cook *model.Cook) errors.ErrorInterface {
	_, err := s.DB.Exec("UPDATE cooks SET name=?, email=?, age=? WHERE email=?", Cook.Name, Cook.Email, Cook.Age, Cook.Email)
	if err != nil {
		err2 := errors.DatabaseUpdationError()
		return err2
	}
	return nil
}

func (s *CookService) DeleteCook(email string) errors.ErrorInterface {
	_, err := s.DB.Exec("DELETE FROM cooks WHERE email=?", email)
	if err != nil {
		err2 := errors.DatabaseDeletionError()
		return err2
	}
	return nil
}
