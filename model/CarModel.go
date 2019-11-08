package model

// CarModel is car model
type CarModel struct {
	Mark  string `form:"mark" json:"mark" binding:"required"`
	Model string `form:"model" json:"model" binding:"required"`
}
