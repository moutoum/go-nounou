package go_nounou

import (
	"fmt"
	"math"
)

type Core struct {
	hourlySalary  float64
	hoursPerMonth float64
	baseSalary    float64
	daysPerMonth  uint
	allowances    float64
}

type Declaration struct {
	Salary     float64
	Hours      uint
	Days       uint
	Allowances float64
}

func NewCore(hourlySalary, hoursPerWeek, daysPerWeek, dailyAllowances float64, workingDays uint) *Core {
	hoursPerMonth := hoursPerWeek * 52 / 12
	daysPerMonth := daysPerWeek * 52 / 12

	return &Core{
		hourlySalary:  hourlySalary,
		hoursPerMonth: hoursPerMonth,
		baseSalary:    hoursPerMonth * hourlySalary,
		daysPerMonth:  uint(math.Ceil(daysPerMonth)),
		allowances:    dailyAllowances * float64(workingDays),
	}
}

func (c *Core) Normal() *Declaration {
	fmt.Println("base salary: ", c.baseSalary*100)
	salary := math.Ceil(c.baseSalary*100) / 100

	return &Declaration{
		Salary:     salary,
		Hours:      uint(math.Round(c.hoursPerMonth)),
		Days:       c.daysPerMonth,
		Allowances: c.allowances,
	}
}

func (c *Core) Absence(absenceHours, potentialHours float64, absenceDays uint) *Declaration {
	salary := c.baseSalary - ((c.baseSalary * absenceHours) / potentialHours)

	return &Declaration{
		Salary:     salary,
		Hours:      uint(math.Round(salary / c.hourlySalary)),
		Days:       c.daysPerMonth - absenceDays,
		Allowances: c.allowances,
	}
}
