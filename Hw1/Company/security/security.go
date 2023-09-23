package security

type Security struct {
	position string
	salary   float64
	address  string
}

func (s *Security) GetPosition() string {
	return s.position
}

func (s *Security) SetPosition(position string) {
	s.position = position
}

func (s *Security) GetSalary() float64 {
	return s.salary
}

func (s *Security) SetSalary(salary float64) {
	s.salary = salary
}

func (s *Security) GetAddress() string {
	return s.address
}

func (s *Security) SetAddress(address string) {
	s.address = address
}
