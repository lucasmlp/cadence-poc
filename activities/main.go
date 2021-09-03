package activities

import (
	"context"
	"fmt"
	"time"
)

func PrintCurrentTime(ctx context.Context) error {
	//dayMonthYearLayout := "02-01-2006"

	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	tokyoTime := time.Now().In(loc) //.Format(dayMonthYearLayout)

	fmt.Printf("tokyoTime: %v\n", tokyoTime)

	loc, err = time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	saoPauloTime := time.Now().In(loc) //.Format(dayMonthYearLayout)
	fmt.Printf("saoPauloTime: %v\n", saoPauloTime)

	return nil
}
