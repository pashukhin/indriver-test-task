package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pashukhin/indriver-test-task/httputil"
)

// Controller interface for all rest controllers
type Controller interface {
	List(context *gin.Context)
	Get(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

func idFromContext(c *gin.Context) *int {
	idStr := c.Param("id")
	result, err := strconv.Atoi(idStr)
	if err != nil {
		httputil.NewError(c, 400, err)
		return nil
	}
	return &result
}
