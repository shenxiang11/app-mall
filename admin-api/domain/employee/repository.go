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
	var employee *Employee
	employee.ID = eid

	if err := r.db.First(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Employee, int64) {
	var count int64
	var list []Employee

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Count(&count)

	return list, count
}

func (r *Repository) Update(e *Employee) error {
	return r.db.Save(&e).Error
}
