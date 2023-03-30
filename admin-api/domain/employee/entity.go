package employee

import (
	"github.com/sx931210/app-mall/utils/orm"
)

type Employee struct {
	orm.BaseModel
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Password   string `json:"-"`
	Salt       string `json:"-"`
	IsNotValid bool   `json:"-""`
}

func NewEmployee(name string, password string, avatar string) *Employee {
	return &Employee{
		Name:     name,
		Avatar:   avatar,
		Password: password,
	}
}
