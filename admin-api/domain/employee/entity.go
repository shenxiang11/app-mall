package employee

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	IsNotValid bool   `json:"isNotValid""`
}

func NewEmployee(name string, password string, avatar string) *Employee {
	return &Employee{
		Name:     name,
		Avatar:   avatar,
		Password: password,
	}
}
