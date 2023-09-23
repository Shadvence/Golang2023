package consultant

type Consultant struct {
	position string
	salary   float64
	address  string
}

func (c *Consultant) GetPosition() string {
	return c.position
}

func (c *Consultant) SetPosition(position string) {
	c.position = position
}

func (c *Consultant) GetSalary() float64 {
	return c.salary
}

func (c *Consultant) SetSalary(salary float64) {
	c.salary = salary
}

func (c *Consultant) GetAddress() string {
	return c.address
}

func (c *Consultant) SetAddress(address string) {
	c.address = address
}
