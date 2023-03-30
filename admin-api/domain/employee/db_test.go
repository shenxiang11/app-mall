package employee

import (
	"fmt"
	"github.com/sx931210/app-mall/utils/database_handler"
	"testing"
)

func TestCreateTable(t *testing.T) {
	db := database_handler.NewMySQLDB("root:123456@tcp(127.0.0.1:33061)/app-mall?parseTime=True&loc=Local")

	r := NewEmployeeRepository(db)

	r.db.AutoMigrate(&Employee{})

	for i := 1; i < 12; i++ {
		r.Create(&Employee{
			Name:     fmt.Sprintf("沈翔 %d", i),
			Avatar:   "https://twitter.com/airi_kijima_sub/photo",
			Password: "123456",
		})
	}
}
