package employee

type Service struct {
	r *Repository
}

func NewEmployeeService(r *Repository) *Service {
	return &Service{r: r}
}

func (s *Service) Create(employee *Employee) error {
	return nil
}

func (s *Service) GetAll() {

}

func (s *Service) Update(employee *Employee) error {
	return s.r.Update(employee)
}
