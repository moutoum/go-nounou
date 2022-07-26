package do

import (
	"fmt"

	"github.com/moutoum/go-nounou/internal"
	"github.com/moutoum/go-nounou/pkg/bookkeeper"
	"github.com/urfave/cli"
)

func printExpenses(employerContributionShares, employeeContributionShares bookkeeper.ContributionShares, employerExpenses, employeeExpenses bookkeeper.Expenses) {
	fmt.Printf("| %-30s | %8s | %8s || %8s | %8s |\n", "Name", "Employee", "Amount", "Employer", "Amount")
	printLine("Taxable CSG", employerContributionShares.TaxableCSG, employeeContributionShares.TaxableCSG, employerExpenses.TaxableCSG.Amount, employeeExpenses.TaxableCSG.Amount)
	printLine("Non-taxable CSG", employerContributionShares.NonTaxableCSG, employeeContributionShares.NonTaxableCSG, employerExpenses.NonTaxableCSG.Amount, employeeExpenses.NonTaxableCSG.Amount)
	printLine("Social Security", employerContributionShares.SocialSecurity, employeeContributionShares.SocialSecurity, employerExpenses.SocialSecurity.Amount, employeeExpenses.SocialSecurity.Amount)
	printLine("FNAL", employerContributionShares.FNAL, employeeContributionShares.FNAL, employerExpenses.FNAL.Amount, employeeExpenses.FNAL.Amount)
	printLine("CSA", employerContributionShares.CSA, employeeContributionShares.CSA, employerExpenses.CSA.Amount, employeeExpenses.CSA.Amount)
	printLine("Professional Training", employerContributionShares.ProfessionalTraining, employeeContributionShares.ProfessionalTraining, employerExpenses.ProfessionalTraining.Amount, employeeExpenses.ProfessionalTraining.Amount)
	printLine("Pension", employerContributionShares.Pension, employeeContributionShares.Pension, employerExpenses.Pension.Amount, employeeExpenses.Pension.Amount)
	printLine("Providence", employerContributionShares.Providence, employeeContributionShares.Providence, employerExpenses.Providence.Amount, employeeExpenses.Providence.Amount)
	printLine("Unemployment Insurance", employerContributionShares.UnemploymentInsurance, employeeContributionShares.UnemploymentInsurance, employerExpenses.UnemploymentInsurance.Amount, employeeExpenses.UnemploymentInsurance.Amount)
	printLine("Social Dialog Contribution", employerContributionShares.SocialDialogContribution, employeeContributionShares.SocialDialogContribution, employerExpenses.SocialDialogContribution.Amount, employeeExpenses.SocialDialogContribution.Amount)
	fmt.Printf("| %-30s | %8s | %8.2f || %8s | %8.2f |\n", "Total", "----", employeeExpenses.Total(), "----", employerExpenses.Total())
}

func printLine(name string, employerShares, employeeShares, employerExpenses, employeeExpenses float64) {
	fmt.Printf("| %-30s | %6.2f %% | %8.2f || %6.2f %% | %8.2f |\n", name, employeeShares*100, employeeExpenses, employerShares*100, employerExpenses)
}

func do(c *cli.Context) error {
	bk := bookkeeper.BookKeeper{
		Employer:          internal.EmployerContributionShares,
		Employee:          internal.EmployeeContributionShares,
		DailyAllowances:   c.GlobalFloat64("daily-allowances"),
		HourlyGrossSalary: c.GlobalFloat64("hourly-gross-salary"),
		WeeklyHours:       c.GlobalFloat64("hours-per-week"),
		DaysPerWeek:       c.GlobalUint("days-per-week"),
	}

	payslip := bk.NewPaySlip(c.Uint("working-days"))

	fmt.Printf("Normal number of hours : %6.2f hours\n", payslip.Hours)
	fmt.Printf("Number of working days : %6.2f days\n", payslip.Days)
	fmt.Printf("Gross salary           : %6.2f €\n", payslip.GrossSalary)
	printExpenses(bk.Employer, bk.Employee, payslip.Employer, payslip.Employee)
	fmt.Printf("Net salary             : %6.2f €\n", payslip.NetSalary)
	fmt.Printf("Maintenance allowances : %6.2f €\n", payslip.Allowances)
	fmt.Printf("Total                  : %6.2f €\n", payslip.NetSalary+payslip.Allowances)

	return nil
}
