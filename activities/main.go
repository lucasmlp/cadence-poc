package activities

import (
	"context"
	"fmt"
	"time"
)

func PrintCurrentTime(ctx context.Context) error {
	dayMonthYearLayout := "02-01-2006"

	utcTime := fmt.Sprintf("UTC time: %v\n", time.Now().Format(dayMonthYearLayout))
	fmt.Printf("utcTime: %v\n", utcTime)
	loc, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return err
	}

	correctedTime := fmt.Sprintf("UTC-3 time: %v\n", time.Now().In(loc).Format(dayMonthYearLayout))
	fmt.Printf("correctedTime: %v\n", correctedTime)
	return nil
}
