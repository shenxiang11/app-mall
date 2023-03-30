package api

import (
	"github.com/gin-gonic/gin"
	employeeApi "github.com/sx931210/app-mall/admin-api/api/employee"
	"github.com/sx931210/app-mall/admin-api/domain/employee"
	"github.com/sx931210/app-mall/admin-api/middleware"
	"github.com/sx931210/app-mall/utils/database_handler"
)

type Databases struct {
	employeeRepository *employee.Repository
}

var authMiddleware = middleware.AuthAdminMiddleware("123456")

func createDBs() *Databases {
	db := database_handler.NewMySQLDB("root:123456@tcp(127.0.0.1:33061)/app-mall?parseTime=True&loc=Local")

	return &Databases{employeeRepository: employee.NewEmployeeRepository(db)}
}

func RegisterHandlers(r *gin.Engine) {
	dbs := createDBs()
	RegisterEmployeeHandlers(r, dbs)
}

func RegisterEmployeeHandlers(r *gin.Engine, dbs *Databases) {
	service := employee.NewEmployeeService(dbs.employeeRepository)
	controller := employeeApi.NewEmployeeController(service)

	group := r.Group("/employee")

	group.GET("/list", authMiddleware, controller.GetAll)
	group.POST("", authMiddleware, controller.Create)
	group.POST("/login", controller.Login)
	group.GET("/self", authMiddleware, controller.GetSelfInfo)
	group.DELETE("/ban/:id", authMiddleware, controller.Ban)
}
