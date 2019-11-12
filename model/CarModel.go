package model

import (
	"github.com/jinzhu/gorm"
)

// CarModel is car model
type CarModel struct {
	gorm.Model
	Mark string `form:"mark" json:"mark" binding:"required" example:"Lada"`
	Name string `form:"name" json:"name" binding:"required" example:"Kalina"`
}
