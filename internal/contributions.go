package internal

import "github.com/moutoum/go-nounou/pkg/bookkeeper"

var (
	EmployerContributionShares = bookkeeper.ContributionShares{
		TaxableCSG:               0,
		NonTaxableCSG:            0,
		SocialSecurity:           0.2965,
		FNAL:                     0.001,
		CSA:                      0.003,
		ProfessionalTraining:     0.0055,
		Pension:                  0.0601,
		Providence:               0.0215,
		UnemploymentInsurance:    0.0405,
		SocialDialogContribution: 0.00016,
	}

	EmployeeContributionShares = bookkeeper.ContributionShares{
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
	}
)
