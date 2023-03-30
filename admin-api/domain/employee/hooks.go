package employee

import "gorm.io/gorm"

// BeforeSave 创建用户前加盐
func (e *Employee) BeforeSave(tx *gorm.DB) error {
	return nil
}
