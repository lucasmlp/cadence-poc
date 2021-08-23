package workflows

import (
	"fmt"
	"time"

	"github.com/lucasmachadolopes/cadencePoc/activities"

	"go.uber.org/cadence/workflow"
)

func HelloWorldWorkflow(ctx workflow.Context) error {
	fmt.Println("Hello world")
	fmt.Println("------------------------------------------------------------------")
	return nil
}

func SimpleWorkflow(ctx workflow.Context) error {

	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    time.Second * 30,
		ScheduleToStartTimeout: time.Second * 120,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	fmt.Println("Started workflow: SimpleWorkflow")
	workflow.Sleep(ctx, time.Second*5)
	workflow.ExecuteActivity(ctx, activities.PrintCurrentTime, "")
	workflow.Sleep(ctx, time.Second*5)
	fmt.Println("Ended workflow: SimpleWorkflow")
	fmt.Println("------------------------------------------------------------------")

	return nil
}

func WaitingSignalWorkflow(ctx workflow.Context, signalName string) error {
	fmt.Println("Started workflow: WaitingSignalWorkflow")

	var signalVal string
	signalChan := workflow.GetSignalChannel(ctx, signalName)

	workflow.Go(ctx, func(ctx workflow.Context) {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Started Go Routine")
		s := workflow.NewSelector(ctx)
		s.AddReceive(signalChan, func(c workflow.Channel, more bool) {
			c.Receive(ctx, &signalVal)
			fmt.Println("Received signal: ", signalVal)
		})
		s.Select(ctx)
		fmt.Println("Ended Go Routine")
		fmt.Println("------------------------------------------------------------------")
	})

	workflow.Sleep(ctx, time.Second*30)
	fmt.Println("Ended workflow: WaitingSignalWorkflow")
	fmt.Println("------------------------------------------------------------------")
	return nil
}
