package qr_generator

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

type FormInfo struct {
	Organization_tin    string
	Empployee_tin       string
	Total_income        string
	Employee_nis        string
	Medical_ins         string
	Income_tax_deducted string
}

func SetValues(o_tin string, e_tin string, t_inc string, e_nis string, m_ins string, i_ded string) {

	var form FormInfo

	form.Organization_tin = o_tin
	form.Empployee_tin = e_tin
	form.Total_income = t_inc
	form.Employee_nis = e_nis
	form.Medical_ins = m_ins
	form.Income_tax_deducted = i_ded

	err := qrcode.WriteFile(form.Organization_tin+","+
		form.Empployee_tin+","+
		form.Total_income+","+
		form.Employee_nis+","+
		form.Medical_ins+","+
		form.Income_tax_deducted,
		qrcode.Medium, 100, "qrimg.png")
	if err != nil {
		fmt.Println("Error creating qr code", err)
	}

}
