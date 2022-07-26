package serve

import "github.com/urfave/cli"

func Command() cli.Command {
	return cli.Command{
		Name:        "serve",
		Description: "Start the Payslip web service",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "bind-addr",
				Usage:  "Address on which the server listen on",
				EnvVar: "BIND_ADDR",
				Value:  ":80",
			},
		},
		Action: serve,
	}
}
