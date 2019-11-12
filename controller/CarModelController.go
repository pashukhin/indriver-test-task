package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/pashukhin/indriver-test-task/httputil"
	"github.com/pashukhin/indriver-test-task/model"

	"github.com/gin-gonic/gin"
)

// CarModelController is controller for car models
type CarModelController struct {
	db *gorm.DB
}

// NewCarModelController makes new CarModelController instance and returns it as Controller
func NewCarModelController(db *gorm.DB) Controller {
	return &CarModelController{db}
}

// List for GET /car_model
// @Summary List car models
// @Description get car models
// @Tags car_model
// @Accept  json
// @Produce  json
// @Success 200 {array} model.CarModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /car_model [get]
func (c *CarModelController) List(context *gin.Context) {
	var carModels []model.CarModel
	if err := c.db.Find(&carModels).Error; err != nil {
		context.AbortWithStatus(404)
	} else {
		context.JSON(200, carModels)
	}
}

// Get for GET /car_model/:id
// @Summary Get car model
// @Description get car model
// @Tags car_model
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CarModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /car_model/{id} [get]
// @Param id path int true "ID"
func (c *CarModelController) Get(context *gin.Context) {
	if id := idFromContext(context); id != nil {
		var carModel model.CarModel
		if err := c.db.Where("id = ?", *id).First(&carModel).Error; err != nil {
			context.AbortWithStatus(404)
		} else {
			context.JSON(200, carModel)
		}
	}
}

// Create for POST
// @Summary Create car model
// @Description create car model
// @Tags car_model
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CarModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /car_model [post]
// @Param account body model.CarModel true "Car model"
func (c *CarModelController) Create(context *gin.Context) {
	var carModel model.CarModel
	if err := context.BindJSON(&carModel); err != nil {
		httputil.NewError(context, 500, err)
	} else {
		if err := c.db.Create(&carModel).Error; err != nil {
			httputil.NewError(context, 500, err)
		} else {
			context.JSON(200, carModel)

		}
	}
}

// Update for PUT
// @Summary Update car model
// @Description update car modles
// @Tags car_model
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CarModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /car_model/{id} [put]
// @Param id path int true "ID"
// @Param account body model.CarModel true "Car model"
func (c *CarModelController) Update(context *gin.Context) {
	if id := idFromContext(context); id != nil {
		var carModel model.CarModel
		if err := c.db.Where("id = ?", id).First(&carModel).Error; err != nil {
			context.AbortWithStatus(404)
		} else {
			if err := context.BindJSON(&carModel); err != nil {
				httputil.NewError(context, 500, err)
			} else {
				if err := c.db.Save(&carModel).Error; err != nil {
					httputil.NewError(context, 500, err)
				} else {
					context.JSON(200, carModel)
				}
			}
		}
	}
}

// Delete for DELETE
// @Summary Delete car model
// @Description delete car modles
// @Tags car_model
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CarModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /car_model/{id} [delete]
// @Param id path int true "ID"
func (c *CarModelController) Delete(context *gin.Context) {
	if id := idFromContext(context); id != nil {
		var carModel model.CarModel
		if err := c.db.Where("id = ?", id).First(&carModel).Error; err != nil {
			context.AbortWithStatus(404)
		} else {
			if err := c.db.Where("id = ?", id).Delete(&carModel).Error; err != nil {
				httputil.NewError(context, 500, err)
			} else {
				context.JSON(200, carModel)
			}
		}
	}
}
