package employee

import (
	"github.com/sx931210/app-mall/utils/hash"
	"github.com/sx931210/app-mall/utils/pagination"
)

type Service struct {
	r *Repository
}

func NewEmployeeService(r *Repository) *Service {
	return &Service{r: r}
}

func (s *Service) Create(name, password, avatar string) error {
	employee := NewEmployee(name, password, avatar)
	return s.r.Create(employee)

}

func (s *Service) GetAll(pagination *pagination.Pagination) *pagination.Pagination {
	list, count := s.r.GetAll(pagination.Page, pagination.PageSize)
	pagination.List = list
	pagination.TotalCount = count
	return pagination
}

func (s *Service) Update(employee *Employee) error {
	return s.r.Update(employee)
}

func (s *Service) GetEmployeeByName(name string, password string) (*Employee, error) {
	e, err := s.r.GetOneByName(name)
	if err != nil {
		return nil, ErrEmployeeNotFount
	}

	match := hash.CheckPassword(password+e.Salt, e.Password)
	if !match {
		return nil, ErrEmployeeNotFount
	}

	return e, nil
}

func (s *Service) GetEmployeeById(id uint) (*Employee, error) {
	e, err := s.r.GetOneById(id)
	if err != nil {
		return nil, ErrEmployeeNotFount
	}

	return e, nil
}
