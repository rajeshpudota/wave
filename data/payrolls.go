package data

import (
	"encoding/json"
	"io"
)

type Payroll struct {
	Date       string `json:"date"`
	Hours      string `json:"hours worked"`
	EmployeeId int    `json:"employee id"`
	Group      string `json:"job group"`
}

type Payrolls []*Payroll

func (p *Payrolls) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetPayrolls() Payrolls {
	return Payrolls
}

func AddPayroll(p *Payroll) {
	// p.ID = getNextID()
	productList = append(productList, p)
}
