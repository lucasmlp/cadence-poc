package activities

import (
	"context"
	"fmt"
	"time"
)

func PrintCurrentTime(ctx context.Context) error {
	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	tokyoTime := time.Now().In(loc)

	fmt.Printf("tokyoTime: %v\n", tokyoTime)

	loc, err = time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	saoPauloTime := time.Now().In(loc)
	fmt.Printf("saoPauloTime: %v\n", saoPauloTime)

	return nil
}

func ActivityA(data string) (string, error) {
	return data + " antigo", nil
}

func ActivityB(data string) (string, error) {
	return data + " + final", nil
}

func ActivityC(data string) (string, error) {
	return data + " novo", nil
}
