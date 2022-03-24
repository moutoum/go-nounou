package bookkeeper

import "math"

type ContributionShares struct {
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

type BookKeeper struct {
	Employer ContributionShares
	Employee ContributionShares

	DailyAllowances   float64
	HourlyGrossSalary float64
	WeeklyHours       float64
	DaysPerWeek       uint
}

const (
	weeksPerYear  = 52
	monthsPerYear = 12
)

func (b *BookKeeper) NewPaySlip(workingDays uint) *PaySlip {
	hoursPerMonth := b.WeeklyHours * weeksPerYear / monthsPerYear
	grossSalary := hoursPerMonth * b.HourlyGrossSalary
	allowances := float64(workingDays) * b.DailyAllowances
	employeeExpenses := newExpenses(grossSalary, b.Employee)
	employerExpenses := newExpenses(grossSalary, b.Employer)
	days := math.Ceil(float64(b.DaysPerWeek) * 52 / 12)

	return &PaySlip{
		Hours:       hoursPerMonth,
		Days:        days,
		GrossSalary: grossSalary,
		Employee:    employeeExpenses,
		Employer:    employerExpenses,
		Allowances:  allowances,
		NetSalary:   grossSalary - employeeExpenses.Total(),
	}
}

func newExpenses(grossSalary float64, shares ContributionShares) Expenses {
	return Expenses{
		TaxableCSG:               (0.9825 * grossSalary) * shares.TaxableCSG,
		NonTaxableCSG:            (0.9825 * grossSalary) * shares.NonTaxableCSG,
		SocialSecurity:           grossSalary * shares.SocialSecurity,
		FNAL:                     grossSalary * shares.FNAL,
		CSA:                      grossSalary * shares.CSA,
		ProfessionalTraining:     grossSalary * shares.ProfessionalTraining,
		Pension:                  grossSalary * shares.Pension,
		Providence:               grossSalary * shares.Providence,
		UnemploymentInsurance:    grossSalary * shares.UnemploymentInsurance,
		SocialDialogContribution: grossSalary * shares.SocialDialogContribution,
	}
}
