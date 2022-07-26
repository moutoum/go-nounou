package serve

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moutoum/go-nounou/internal"
	"github.com/moutoum/go-nounou/pkg/bookkeeper"
	"github.com/urfave/cli"
)

func serve(c *cli.Context) error {
	bk := bookkeeper.BookKeeper{
		Employer:          internal.EmployerContributionShares,
		Employee:          internal.EmployeeContributionShares,
		DailyAllowances:   c.GlobalFloat64("daily-allowances"),
		HourlyGrossSalary: c.GlobalFloat64("hourly-gross-salary"),
		WeeklyHours:       c.GlobalFloat64("hours-per-week"),
		DaysPerWeek:       c.GlobalUint("days-per-week"),
	}

	router := chi.NewRouter()
	router.Get("/do", bk.HandleGetPayslip)

	server := http.Server{
		Addr:    c.String("bind-addr"),
		Handler: router,
	}

	fmt.Println("Starting server on", c.String("bind-addr"))
	return server.ListenAndServe()
}
