package bookkeeper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookKeeper_NewPaySlip(t *testing.T) {
	bk := BookKeeper{
		Employer: ContributionShares{
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
		},
		Employee: ContributionShares{
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
		},
		DailyAllowances:   4.60,
		HourlyGrossSalary: 4.92,
		WeeklyHours:       38,
		DaysPerWeek:       4,
	}

	payslip := bk.NewPaySlip(16)

	assert.InDelta(t, 810.16, payslip.GrossSalary, 0.01)
	assert.Equal(t, 18.0, payslip.Days)
	assert.InDelta(t, 164.6, payslip.Hours, 0.1)
	assert.Equal(t, 73.6, payslip.Allowances)
	assert.InDelta(t, 632.24, payslip.NetSalary, 0.01)
}

func TestToto(t *testing.T) {
	ex := newExpenses(22.48, ContributionShares{
		TaxableCSG:               0.029,
		NonTaxableCSG:            0.068,
		SocialSecurity:           0.073,
		FNAL:                     0,
		CSA:                      0,
		ProfessionalTraining:     0,
		Pension:                  0.0401,
		Providence:               0.0104,
		UnemploymentInsurance:    0,
		SocialDialogContribution: 0,
	})

	fmt.Printf("Brut: 22.48, Employee Taxes: %f, Net: %f", ex.Total(), 22.48-ex.Total())
}
