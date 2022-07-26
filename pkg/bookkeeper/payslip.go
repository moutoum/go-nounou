package bookkeeper

type Expenses struct {
	TaxableCSG               Expense
	NonTaxableCSG            Expense
	SocialSecurity           Expense
	FNAL                     Expense
	CSA                      Expense
	ProfessionalTraining     Expense
	Pension                  Expense
	Providence               Expense
	UnemploymentInsurance    Expense
	SocialDialogContribution Expense
}

type Expense struct {
	Base   float64
	Amount float64
}

func (e *Expenses) Total() float64 {
	var total float64
	total += e.TaxableCSG.Amount
	total += e.NonTaxableCSG.Amount
	total += e.SocialSecurity.Amount
	total += e.FNAL.Amount
	total += e.CSA.Amount
	total += e.ProfessionalTraining.Amount
	total += e.Pension.Amount
	total += e.Providence.Amount
	total += e.UnemploymentInsurance.Amount
	total += e.SocialDialogContribution.Amount
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
