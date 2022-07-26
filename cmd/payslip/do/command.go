package do

import "github.com/urfave/cli"

func Command() cli.Command {
	return cli.Command{
		Name:        "do",
		Description: "Compute the payslip",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:     "working-days",
				Usage:    "Number of days worked in the month",
				Required: true,
			},
		},
		Action: do,
	}
}
