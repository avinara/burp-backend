package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/burp-backend/errors"
	"github.com/burp-backend/model"
	"github.com/burp-backend/services"
	"github.com/burp-backend/utils"
)

type CookControllerAPI interface {
	GetAllCooks(w http.ResponseWriter, r *http.Request)
	GetCookByEmail(w http.ResponseWriter, r *http.Request)
	CreateCook(w http.ResponseWriter, r *http.Request)
	UpdateCook(w http.ResponseWriter, r *http.Request)
	DeleteCook(w http.ResponseWriter, r *http.Request)
}

type CookController struct {
	CookService services.CookServiceAPI
}

func NewCookController(CookService services.CookServiceAPI) *CookController {
	return &CookController{
		CookService: CookService,
	}
}

// @Summary Get all Cooks
// @Description Get all Cooks from the database
// @Tags Cooks
// @Produce json
// @Success 200 {array} model.Cook
// @Failure 500 {object} model.ErrorResponse
// @Router /cooks [get]
func (sc *CookController) GetAllCooks(w http.ResponseWriter, r *http.Request) {
	Cooks, err := sc.CookService.GetAllCooks()
	if err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
	json.NewEncoder(w).Encode(Cooks)
}

// @Summary Get a Cook by Email
// @Description Get a Cook by Email from the database
// @Tags Cooks
// @Produce json
// @Param email query string true "Cook Email"
// @Success 200 {object} model.Cook
// @Failure 400 {object} model.ErrorResponse
// @Router /cook [get]
func (sc *CookController) GetCookByEmail(w http.ResponseWriter, r *http.Request) {
	emailParam := r.URL.Query().Get("email")

	Cook, err := sc.CookService.GetCookByEmail(emailParam)
	if err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
	if Cook == nil {
		err2 := errors.CookNotFoundError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}
	json.NewEncoder(w).Encode(Cook)
}

// @Summary Create a new Cook
// @Description Create a new Cook in the database
// @Tags Cooks
// @Accept json
// @Produce json
// @Param Cook body model.Cook true "Cook object"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /cook [post]
func (sc *CookController) CreateCook(w http.ResponseWriter, r *http.Request) {
	var Cook model.Cook
	if err := json.NewDecoder(r.Body).Decode(&Cook); err != nil {
		err2 := errors.InvalidRequestError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}

	ok, err := sc.CookService.CreateCook(&Cook)
	if err != nil || !ok {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
}

// @Summary Update a Cook
// @Description Update an existing Cook in the database
// @Tags Cooks
// @Accept json
// @Produce json
// @Param Cook body model.Cook true "Cook object"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /Cook/update [post]
func (sc *CookController) UpdateCook(w http.ResponseWriter, r *http.Request) {
	var Cook model.Cook
	if err := json.NewDecoder(r.Body).Decode(&Cook); err != nil {
		err2 := errors.InvalidRequestError()
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err2))
		return
	}

	if err := sc.CookService.UpdateCook(&Cook); err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}

}

// @Summary Delete a Cook
// @Description Delete a Cook by ID from the database
// @Tags Cooks
// @Param email query string true "Cook Emails"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /Cook/delete [delete]
func (sc *CookController) DeleteCook(w http.ResponseWriter, r *http.Request) {
	emailParam := r.URL.Query().Get("email")
	if err := sc.CookService.DeleteCook(emailParam); err != nil {
		utils.WriteErrorWithMessage(w, utils.FormErrorMessage(err))
		return
	}
}
