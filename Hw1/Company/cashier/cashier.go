package cashier

type Cashier struct {
	position string
	salary   float64
	address  string
}

func (c *Cashier) GetPosition() string {
	return c.position
}

func (c *Cashier) SetPosition(position string) {
	c.position = position
}

func (c *Cashier) GetSalary() float64 {
	return c.salary
}

func (c *Cashier) SetSalary(salary float64) {
	c.salary = salary
}

func (c *Cashier) GetAddress() string {
	return c.address
}

func (c *Cashier) SetAddress(address string) {
	c.address = address
}
