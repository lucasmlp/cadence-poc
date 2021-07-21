package workflows

import (
	"fmt"

	"go.uber.org/cadence/workflow"
)

func HelloWorldWorkflow(ctx workflow.Context) error {
	fmt.Println("Hello world")
	return nil
}
