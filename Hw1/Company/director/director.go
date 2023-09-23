package director

type Director struct {
	position string
	salary   float64
	address  string
}

func (d *Director) GetPosition() string {
	return d.position
}

func (d *Director) SetPosition(position string) {
	d.position = position
}

func (d *Director) GetSalary() float64 {
	return d.salary
}

func (d *Director) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Director) GetAddress() string {
	return d.address
}

func (d *Director) SetAddress(address string) {
	d.address = address
}
