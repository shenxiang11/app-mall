package api

import (
	"github.com/gin-gonic/gin"
	employeeApi "github.com/sx931210/app-mall/admin-api/api/employee"
	"github.com/sx931210/app-mall/admin-api/domain/employee"
	"github.com/sx931210/app-mall/utils/database_handler"
)

type Databases struct {
	employeeRepository *employee.Repository
}

func createDBs() *Databases {
	db := database_handler.NewMySQLDB("root:123456@tcp(127.0.0.1:33061)/go_database?parseTime=True&loc=Local")

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

	group.GET("/list", controller.GetAll)
}
