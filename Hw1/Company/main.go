package main

import (
	"Company/cashier"
	"Company/consultant"
	"Company/director"
	"Company/employee"
	"Company/manager"
	"Company/security"
	"fmt"
)

func main() {
	mng := &manager.Manager{}
	mng.SetPosition("Manager")
	mng.SetSalary(60000)
	mng.SetAddress("Almaty")

	scr := &security.Security{}
	scr.SetPosition("Security")
	scr.SetSalary(20000)
	scr.SetAddress("Astana")

	dir := &director.Director{}
	dir.SetPosition("Director")
	dir.SetSalary(80000)
	dir.SetAddress("Aktau")

	cns := &consultant.Consultant{}
	cns.SetPosition("Consultant")
	cns.SetSalary(40000)
	cns.SetAddress("Semey")

	csh := &cashier.Cashier{}
	csh.SetPosition("Cashier")
	csh.SetSalary(40000)
	csh.SetAddress("Shymkent")

	Info(mng)
	Info(csh)
	Info(scr)
	Info(cns)
	Info(dir)
}

func Info(e employee.Employee) {
	fmt.Printf("Position: %s\n", e.GetPosition())
	fmt.Printf("Salary: $%.2f\n", e.GetSalary())
	fmt.Printf("Address: %s\n", e.GetAddress())
	fmt.Println()
}
