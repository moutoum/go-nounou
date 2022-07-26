package main

import (
	"fmt"
	"os"

	"github.com/moutoum/go-nounou/cmd/payslip/do"
	"github.com/moutoum/go-nounou/cmd/payslip/serve"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "Payslip",
		Usage: "Babysitter declaration helper tool",
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
		},
		Commands: []cli.Command{
			do.Command(),
			serve.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "go-nounou-cli: error: %v", err)
		return
	}
}
