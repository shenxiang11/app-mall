package employee

import (
	"github.com/sx931210/app-mall/utils/hash"
	"gorm.io/gorm"
)

// BeforeSave 创建用户前加盐
func (e *Employee) BeforeSave(tx *gorm.DB) error {

	if e.Salt == "" {
		salt := hash.CreateSalt() // TODO: 注入的话，感觉不能在这
		encryptedPassword, err := hash.EncryptPassword(e.Password + salt)
		if err != nil {
			return err
		}
		e.Password = encryptedPassword
		e.Salt = salt
	}

	return nil
}
