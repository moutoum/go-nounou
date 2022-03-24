package bookkeeper

type Expenses struct {
	TaxableCSG               float64
	NonTaxableCSG            float64
	SocialSecurity           float64
	FNAL                     float64
	CSA                      float64
	ProfessionalTraining     float64
	Pension                  float64
	Providence               float64
	UnemploymentInsurance    float64
	SocialDialogContribution float64
}

func (e *Expenses) Total() float64 {
	var total float64
	total += e.TaxableCSG
	total += e.NonTaxableCSG
	total += e.SocialSecurity
	total += e.FNAL
	total += e.CSA
	total += e.ProfessionalTraining
	total += e.Pension
	total += e.Providence
	total += e.UnemploymentInsurance
	total += e.SocialDialogContribution
	return total
}

type PaySlip struct {
	Hours       float64
	Days        float64
	GrossSalary float64
	Employee    Expenses
	Employer    Expenses
	Allowances  float64
	NetSalary   float64
}
