package bookkeeper

import (
	"fmt"
	// "html/template"
	"math"
	"net/http"
)

// var tmpl = template.Must(
// 	template.
// 		New("get-payslip.html").
// 		Funcs(template.FuncMap{"PrintEuros": printEuros}).
// 		ParseFiles("templates/get-payslip.html"),
// )

type GetPayslipData struct {
	*PaySlip

	EmployeeContributionShares *ContributionShares
	EmployerContributionShares *ContributionShares
	Total                      float64
}

func (b *BookKeeper) HandleGetPayslip(rw http.ResponseWriter, req *http.Request) {
	// if err := req.ParseForm(); err != nil {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	//
	// workingDaysStr := req.Form.Get("working-days")
	// workingDays, err := strconv.ParseUint(workingDaysStr, 10, 32)
	// if err != nil {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	_, _ = fmt.Fprintf(rw, "invalid working days: %s", err)
	// 	return
	// }
	//
	// payslip := b.NewPaySlip(uint(workingDays))
	//
	// data := GetPayslipData{
	// 	PaySlip:                    payslip,
	// 	EmployeeContributionShares: &b.Employee,
	// 	EmployerContributionShares: &b.Employer,
	// 	Total:                      payslip.NetSalary + payslip.Allowances,
	// }
	//
	// if err = tmpl.Execute(rw, data); err != nil {
	// 	fmt.Printf("error while executing template: %v", err)
	// 	return
	// }
}

func printEuros(v float64) string {
	return fmt.Sprintf("%.2f", math.Round(v*100.0)/100.0)
}
