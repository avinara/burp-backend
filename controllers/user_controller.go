package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/burp-backend/errors"
	"github.com/burp-backend/model"
	"github.com/burp-backend/services"
	"github.com/burp-backend/utils"
)

type UserControllerAPI interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserByName(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	UserService services.UserServiceAPI
}

func NewUserController(UserService services.UserServiceAPI) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

// @Summary Get all Users
// @Description Get all Users from the database
// @Tags Users
// @Produce json
// @Success 200 {array} model.User
// @Failure 400 {object} model.ErrorResponse
// @Router /users [get]
func (cc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	Users, err := cc.UserService.GetAllUsers()
	if err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
	json.NewEncoder(w).Encode(Users)
}

// @Summary Get a User by Name
// @Description Get a User by ID from the database
// @Tags Users
// @Produce json
// @Param name query string true "User Name"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Router /user [get]
func (cc *UserController) GetUserByName(w http.ResponseWriter, r *http.Request) {
	nameParam := r.URL.Query().Get("name")
	User, err := cc.UserService.GetUserByName(nameParam)
	if err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
	if User == nil {
		err2 := errors.UserNotFoundError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}
	json.NewEncoder(w).Encode(User)
}

// @Summary Create a new User
// @Description Create a new User in the database
// @Tags Users
// @Accept json
// @Produce json
// @Param User body model.User true "User object"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /user [post]
func (cc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var User model.User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		err2 := errors.InvalidRequestError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}

	_, err := cc.UserService.CreateUser(&User)
	if err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
}

// @Summary Update a User
// @Description Update an existing User in the database
// @Tags Users
// @Accept json
// @Produce json
// @Param User body model.User true "User object"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /user/update [post]
func (cc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var User model.User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		err2 := errors.InvalidRequestError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}

	if err := cc.UserService.UpdateUser(&User); err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
}

// @Summary Delete a User
// @Description Delete a User by ID from the database
// @Tags Users
// @Param id query int true "User ID"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /user/delete [delete]
func (cc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		err2 := errors.QueryParamUnavailableError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}

	if err := cc.UserService.DeleteUser(id); err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
}
