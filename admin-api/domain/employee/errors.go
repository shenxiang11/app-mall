package employee

import "errors"

var (
	ErrEmployeeExistWithName = errors.New("员工已存在")
	ErrEmployeeNotFount      = errors.New("员工未找到到")
	ErrMisMatchedPasswords   = errors.New("密码不正确")
	ErrInvalidEmployeeName   = errors.New("无效员工名")
	ErrInvalidPassword       = errors.New("无效密码")
)
