package employee

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(e *Employee) error {
	return r.db.Create(e).Error
}

func (r *Repository) GetOneById(eid uint) (*Employee, error) {
	var employee Employee

	// 写字符串感觉很不好，但是不写字符串遇到 0 值，sql 会有问题
	if err := r.db.Where("id = ?", eid).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *Repository) GetOneByName(name string) (*Employee, error) {
	var employee Employee
	employee.Name = name

	if err := r.db.First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Employee, int) {
	var count int64
	list := make([]Employee, 0) // 不然空 的 slice 会返回 null

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Offset(-1).Limit(-1).Count(&count)

	return list, int(count)
}

func (r *Repository) Update(e *Employee) error {
	return r.db.Save(&e).Error
}
