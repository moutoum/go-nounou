package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	nounou "github.com/moutoum/go-nounou"
)

func run(c *cli.Context) error {
	core := nounou.NewCore(
		c.Float64("hourly-net-salary"),
		c.Float64("hours-per-week"),
		c.Float64("days-per-week"),
		c.Float64("daily-allowances"),
		c.Uint("working-days"),
	)

	var declaration *nounou.Declaration

	if c.Float64("absence-hours") != 0 {
		declaration = core.Absence(c.Float64("absence-hours"), c.Float64("potential-working-hours"), c.Uint("absence-days"))
	} else {
		declaration = core.Normal()
	}

	fmt.Printf("Normal number of hours : %6d hours\n", declaration.Hours)
	fmt.Printf("Number of working days : %6d days\n", declaration.Days)
	fmt.Printf("Net salary             : %6.2f €\n", declaration.Salary)
	fmt.Printf("Maintenance allowances : %6.2f €\n", declaration.Allowances)

	return nil
}

func main() {
	app := &cli.App{
		Name:   "Go Nounou",
		Usage:  "Babysitter declaration helper tool",
		Action: run,
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "hourly-net-salary",
				Usage: "Net salary per hour",
				Value: 3.76,
			},
			&cli.Float64Flag{
				Name:  "daily-allowances",
				Usage: "Daily allowances at babysitter house in euros",
				Value: 4.50,
			},
			&cli.Float64Flag{
				Name:  "hours-per-week",
				Usage: "Number of working hours per weeks (average)",
				Value: 42.75,
			},
			&cli.Float64Flag{
				Name:  "days-per-week",
				Usage: "Number of working days per week (average)",
				Value: 4.5,
			},
			&cli.UintFlag{
				Name:     "working-days",
				Usage:    "Number of days worked in the month",
				Required: true,
			},
			&cli.Float64Flag{
				Name:  "absence-hours",
				Usage: "Number of hours of absence",
				Value: 0,
			},
			&cli.UintFlag{
				Name:  "absence-days",
				Usage: "Number of days of absence",
				Value: 0,
			},
			&cli.Float64Flag{
				Name:  "potential-working-hours",
				Usage: "Potential working hours during the month (if absences)",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "go-nounou-cli: error: %v", err)
		return
	}
}
