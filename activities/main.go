package activities

import (
	"context"
	"fmt"
	"time"
)

func PrintCurrentTime(ctx context.Context) (string, string, error) {
	dayMonthYearLayout := "02-01-2006"

	utcTime := fmt.Sprintf("UTC time: %v\n", time.Now().Format(dayMonthYearLayout))

	loc, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return "", "", err
	}

	correctedTime := fmt.Sprintf("UTC-3 time: %v\n", time.Now().In(loc).Format(dayMonthYearLayout))

	return utcTime, correctedTime, nil
}
