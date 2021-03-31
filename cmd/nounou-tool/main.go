package main

import (
	"fmt"
	"strconv"

	"github.com/gen2brain/dlgs"

	nounou "github.com/moutoum/go-nounou"
)

func getFloat(text string, defaultValue float64) (float64, bool) {
	value, _, err := dlgs.Entry("Nounou Tool", text, fmt.Sprintf("%.2f", defaultValue))
	if err != nil {
		return 0, false
	}

	parsedValue, err := strconv.ParseFloat(value, 64)
	return parsedValue, err == nil
}

func getInt(text string, defaultValue uint) (uint, bool) {
	value, _, err := dlgs.Entry("Nounou Tool", text, strconv.FormatUint(uint64(defaultValue), 10))
	if err != nil {
		return 0, false
	}

	parsedValue, err := strconv.ParseUint(value, 10, 64)
	return uint(parsedValue), err == nil
}

func main() {
	hourlySalary, ok := getFloat("Enter the hourly net salary", 3.76)
	if !ok {
		return
	}

	hoursPerWeek, ok := getFloat("Enter the average number of hours per week", 42.75)
	if !ok {
		return
	}

	daysPerWeek, ok := getFloat("Enter the average number of days per week", 4.5)
	if !ok {
		return
	}

	dailyAllowances, ok := getFloat("Enter the daily allowances", 4.5)
	if !ok {
		return
	}

	workingDays, ok := getInt("Enter the number of working days during the month", 0)

	core := nounou.NewCore(hourlySalary, hoursPerWeek, daysPerWeek, dailyAllowances, workingDays)
	declaration := core.Normal()

	dlgs.Info("Nounou Tool", fmt.Sprintf("Hours: %d hours\nDays: %d days\nSalary: %.2f €\nAllowances: %.2f €", declaration.Hours, declaration.Days, declaration.Salary, declaration.Allowances))
}
