package workflows

import (
	"fmt"
	"time"

	"go.uber.org/cadence/workflow"
)

func HelloWorldWorkflow(ctx workflow.Context) error {
	fmt.Println("Hello world")
	return nil
}

func SimpleWorkflow(ctx workflow.Context) error {

	fmt.Println("Started workflow")
	workflow.Sleep(ctx, time.Second*5)
	fmt.Println("Ended workflow")

	return nil
}
