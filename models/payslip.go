package models

type Payslip struct {
	Month       string  `json:"month"`
	Year        int     `json:"year"`
	Salary      float32 `json:"salary"`
	PayslipImg  string  `json:"payslip_img"`
	Description string  `json:"description"`
}

func NewPayslip() *Payslip {
	return &Payslip{}
}
