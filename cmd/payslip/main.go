package main

import (
	"fmt"
	"os"

	"github.com/moutoum/go-nounou/pkg/bookkeeper"
	"github.com/urfave/cli"
)

var (
	employerContributionShares = bookkeeper.ContributionShares{
		TaxableCSG:               0,
		NonTaxableCSG:            0,
		SocialSecurity:           0.2965,
		FNAL:                     0.001,
		CSA:                      0.003,
		ProfessionalTraining:     0.0055,
		Pension:                  0.0601,
		Providence:               0.0205,
		UnemploymentInsurance:    0.0405,
		SocialDialogContribution: 0.00016,
	}

	employeeContributionShares = bookkeeper.ContributionShares{
		TaxableCSG:               0.029,
		NonTaxableCSG:            0.068,
		SocialSecurity:           0.073,
		FNAL:                     0,
		CSA:                      0,
		ProfessionalTraining:     0,
		Pension:                  0.0401,
		Providence:               0.0112,
		UnemploymentInsurance:    0,
		SocialDialogContribution: 0,
	}
)

func printExpenses(employerContributionShares, employeeContributionShares bookkeeper.ContributionShares, employerExpenses, employeeExpenses bookkeeper.Expenses) {
	fmt.Printf("| %-30s | %8s | %8s || %8s | %8s |\n", "Name", "Employee", "Amount", "Employer", "Amount")
	printLine("Taxable CSG", employerContributionShares.TaxableCSG, employeeContributionShares.TaxableCSG, employerExpenses.TaxableCSG, employeeExpenses.TaxableCSG)
	printLine("Non-taxable CSG", employerContributionShares.NonTaxableCSG, employeeContributionShares.NonTaxableCSG, employerExpenses.NonTaxableCSG, employeeExpenses.NonTaxableCSG)
	printLine("Social Security", employerContributionShares.SocialSecurity, employeeContributionShares.SocialSecurity, employerExpenses.SocialSecurity, employeeExpenses.SocialSecurity)
	printLine("FNAL", employerContributionShares.FNAL, employeeContributionShares.FNAL, employerExpenses.FNAL, employeeExpenses.FNAL)
	printLine("CSA", employerContributionShares.CSA, employeeContributionShares.CSA, employerExpenses.CSA, employeeExpenses.CSA)
	printLine("Professional Training", employerContributionShares.ProfessionalTraining, employeeContributionShares.ProfessionalTraining, employerExpenses.ProfessionalTraining, employeeExpenses.ProfessionalTraining)
	printLine("Pension", employerContributionShares.Pension, employeeContributionShares.Pension, employerExpenses.Pension, employeeExpenses.Pension)
	printLine("Providence", employerContributionShares.Providence, employeeContributionShares.Providence, employerExpenses.Providence, employeeExpenses.Providence)
	printLine("Unemployment Insurance", employerContributionShares.UnemploymentInsurance, employeeContributionShares.UnemploymentInsurance, employerExpenses.UnemploymentInsurance, employeeExpenses.UnemploymentInsurance)
	printLine("Social Dialog Contribution", employerContributionShares.SocialDialogContribution, employeeContributionShares.SocialDialogContribution, employerExpenses.SocialDialogContribution, employeeExpenses.SocialDialogContribution)
	fmt.Printf("| %-30s | %8s | %8.2f || %8s | %8.2f |\n", "Total", "----", employeeExpenses.Total(), "----", employerExpenses.Total())
}

func printLine(name string, employerShares, employeeShares, employerExpenses, employeeExpenses float64) {
	fmt.Printf("| %-30s | %6.2f %% | %8.2f || %6.2f %% | %8.2f |\n", name, employeeShares*100, employeeExpenses, employerShares*100, employerExpenses)
}

func run(c *cli.Context) error {
	bk := bookkeeper.BookKeeper{
		Employer:          employerContributionShares,
		Employee:          employeeContributionShares,
		DailyAllowances:   c.Float64("daily-allowances"),
		HourlyGrossSalary: c.Float64("hourly-gross-salary"),
		WeeklyHours:       c.Float64("hours-per-week"),
		DaysPerWeek:       c.Uint("days-per-week"),
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

func main() {
	app := &cli.App{
		Name:   "Payslip",
		Usage:  "Babysitter declaration helper tool",
		Action: run,
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "hourly-gross-salary",
				Usage: "Gross salary per hour",
				Value: 4.92,
			},
			&cli.Float64Flag{
				Name:  "daily-allowances",
				Usage: "Daily allowances at babysitter house",
				Value: 4.60,
			},
			&cli.Float64Flag{
				Name:  "hours-per-week",
				Usage: "Number of working hours per weeks (average)",
				Value: 38,
			},
			&cli.Float64Flag{
				Name:  "days-per-week",
				Usage: "Number of working days per week (average)",
				Value: 4,
			},
			&cli.UintFlag{
				Name:     "working-days",
				Usage:    "Number of days worked in the month",
				Required: true,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "go-nounou-cli: error: %v", err)
		return
	}
}
