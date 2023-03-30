package employee

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sx931210/app-mall/admin-api/domain/employee"
	"github.com/sx931210/app-mall/utils/api_helper"
	jwt_helper "github.com/sx931210/app-mall/utils/jwt"
	"github.com/sx931210/app-mall/utils/pagination"
	"net/http"
	"os"
	"time"
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
// @Tags Employee
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Router /employee/list [get]
func (c *Controller) GetAll(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)

	page = c.employeeService.GetAll(page)

	g.JSON(http.StatusOK, page)
}

// Create godoc
// @Summary 创建员工
// @Tags Employee
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param CreateEmployeeRequest body CreateEmployeeRequest true "员工信息"
// @Router /employee [post]
func (c *Controller) Create(g *gin.Context) {
	var req CreateEmployeeRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	err := c.employeeService.Create(req.Name, req.Password, req.Avatar)
	if err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	g.JSON(http.StatusCreated, api_helper.Response{Message: "创建成功"})
}

// Login godoc
// @Summary 根据用户名和密码登录
// @Tags Employee
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "登录信息"
// @Router /employee/login [post]
func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	e, err := c.employeeService.GetEmployeeByName(req.Name, req.Password)
	if err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	// 如果登录者没登录，怎么搞呢？
	if e.IsNotValid {
		api_helper.HandlerError(g, errors.New("禁止登录"))
		return
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  e.ID,
		"iat": time.Now().Unix(),
		"iss": os.Getenv("ENV"),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	token, err := jwt_helper.GenerateToken(jwtClaims, "123456")
	if err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	g.JSON(http.StatusOK, LoginResponse{Token: token})
}

// GetSelfInfo godoc
// @Summary 获取当前登录员工信息
// @Tags Employee
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Router /employee/self [get]
func (c *Controller) GetSelfInfo(g *gin.Context) {
	eid := g.GetInt("uid")

	info, err := c.employeeService.GetEmployeeById(uint(eid))
	if err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	g.JSON(http.StatusOK, info)
}

// Ban godoc
// @Summary 禁止指定用户再次登录
// @Tags Employee
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param        id   path      int  true  "Account ID"
// @Router /employee/ban/{id} [delete]
func (c *Controller) Ban(g *gin.Context) {
	var params IdUri
	if err := g.BindUri(&params); err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	// 什么 bug 0, 会找到数据库第一条记录
	e, err := c.employeeService.GetEmployeeById(params.Id)
	if err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	e.IsNotValid = true
	err = c.employeeService.Update(e)
	if err != nil {
		api_helper.HandlerError(g, err)
		return
	}

	g.JSON(http.StatusOK, api_helper.Response{Message: "更新成功"})
}
