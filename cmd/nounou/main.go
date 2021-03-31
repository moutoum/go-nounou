package main

import (
	"fmt"
	"github.com/urfave/cli"
	"math"
	"os"
)

const (
	WorkingWeeksPerYear = 47
	VacationWeeksPerYear = 5
)

func run(c *cli.Context) error {
	var (
		netSalary float64
		normalHours uint
		workingDays uint
	)

	hourlyNetSalary := c.Float64("hourly-net-salary")
	hoursPerMonth := c.Float64("hours-per-week") * (WorkingWeeksPerYear + VacationWeeksPerYear) / 12
	monthlyBaseSalary := hoursPerMonth * hourlyNetSalary
	roundedHoursPerMonth := math.Round(hoursPerMonth)
	monthlyWorkingDays := c.Float64("days-per-week") * (WorkingWeeksPerYear + VacationWeeksPerYear) / 12
	roundedMonthlyWorkingDays := math.Ceil(monthlyWorkingDays)
	allowances := float64(c.Uint("working-days")) * c.Float64("daily-allowances")

	absenceHours := c.Float64("absence-hours")
	if absenceHours != 0 {
		netSalary = monthlyBaseSalary - ((monthlyBaseSalary * c.Float64("absence-hours")) / c.Float64("potential-working-hours"))
		normalHours = uint(math.Round(netSalary / hourlyNetSalary))
		workingDays = uint(roundedMonthlyWorkingDays) - c.Uint("absence-days")
	} else {
		netSalary = monthlyBaseSalary
		normalHours = uint(roundedHoursPerMonth)
		workingDays = uint(roundedMonthlyWorkingDays)
	}

	fmt.Printf("Normal number of hours : %6d hours\n", normalHours)
	fmt.Printf("Number of working days : %6d days\n", workingDays)
	fmt.Printf("Net salary             : %6.2f €\n", netSalary)
	fmt.Printf("Maintenance allowances : %6.2f €\n", allowances)

	return nil
}

func main() {
	app := &cli.App{
		Name:   "Go Nounou",
		Usage:  "Babysitter declaration helper tool",
		Action: run,
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name: "hourly-net-salary",
				Usage: "Net salary per hour",
				Value: 3.76,
			},
			&cli.Float64Flag{
				Name: "daily-allowances",
				Usage: "Daily allowances at babysitter house in euros",
				Value: 4.50,
			},
			&cli.Float64Flag{
				Name: "hours-per-week",
				Usage: "Number of working hours per weeks (average)",
				Value: 42.75,
			},
			&cli.Float64Flag{
				Name: "days-per-week",
				Usage: "Number of working days per week (average)",
				Value: 4.5,
			},
			&cli.UintFlag{
				Name: "working-days",
				Usage: "Number of days worked in the month",
				Required: true,
			},
			&cli.Float64Flag{
				Name: "absence-hours",
				Usage: "Number of hours of absence",
				Value: 0,
			},
			&cli.UintFlag{
				Name: "absence-days",
				Usage: "Number of days of absence",
				Value: 0,
			},
			&cli.Float64Flag{
				Name: "potential-working-hours",
				Usage: "Potential working hours during the month (if absences)",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "go-nounou: error: %v", err)
		return
	}
}