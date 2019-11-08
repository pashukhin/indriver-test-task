package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pashukhin/indriver-test-task/httputil"
)

// CarModel is controller for car models
type CarModel struct {
}

// List for GET /car_model
// @Summary List car models
// @Description get car modles
// @Tags car_model
// @Accept  json
// @Produce  json
// @Success 200 {array} model.CarModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /car_model [get]
func (c *CarModel) List(context *gin.Context) {
	httputil.NewError(context, http.StatusInternalServerError, errors.New("Not implemented"))
	return
}

// Get for GET /car_model/:id
func (c *CarModel) Get(context *gin.Context) {

}

// Create for POST
func (c *CarModel) Create(context *gin.Context) {

}

// Update for PUT
func (c *CarModel) Update(context *gin.Context) {

}

// Delete for DELETE
func (c *CarModel) Delete(context *gin.Context) {

}
