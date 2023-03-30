package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/sx931210/app-mall/admin-api/domain/employee"
	"github.com/sx931210/app-mall/utils/api_helper"
	"net/http"
)

type Controller struct {
	employeeService *employee.Service
}

func NewEmployeeController(service *employee.Service) *Controller {
	return &Controller{employeeService: service}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// GetAll godoc
// @Summary 查询所有员工
// @Tags Admin
// @Accept json
// @Product json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} api_helper.ErrorResponse
// @Failure 400 {object} api_helper.ErrorResponse
// @Router /employee/list [get]
func (c *Controller) GetAll(g *gin.Context) {
	g.JSON(http.StatusBadRequest, api_helper.ErrorResponse{Message: "123"})
}
