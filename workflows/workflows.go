package workflows

import (
	"fmt"
	"time"

	"github.com/lucasmachadolopes/cadencePoc/activities"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func HelloWorldWorkflow(ctx workflow.Context) error {
	fmt.Println("Hello world")
	fmt.Println("------------------------------------------------------------------")

	return nil
}

func WaitingSignalWorkflow(ctx workflow.Context, signalName string) error {
	fmt.Println("Started workflow: WaitingSignalWorkflow")
	internalContext, cancelfunc := workflow.WithCancel(externalContext)

	var signalVal string
	signalChan := workflow.GetSignalChannel(internalContext, signalName)

	workflow.Go(internalContext, func(ctx workflow.Context) {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Started Go Routine")
		s := workflow.NewSelector(ctx)
		s.AddReceive(signalChan, func(c workflow.Channel, more bool) {
			c.Receive(ctx, &signalVal)
			fmt.Println("Received signal: ", signalVal)
			cancelfunc()
		})
		s.Select(ctx)
		fmt.Println("Ended Go Routine")
		fmt.Println("------------------------------------------------------------------")
	})

	workflow.Sleep(internalContext, time.Second*30)
	fmt.Println("Ended workflow: WaitingSignalWorkflow")
	fmt.Println("------------------------------------------------------------------")
	return nil
}

func ActivityWorkflow(ctx workflow.Context) error {

	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    time.Second * 30,
		ScheduleToStartTimeout: time.Second * 120,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	fmt.Println("Started workflow: ActivityWorkflow")

	workflow.Sleep(ctx, time.Second*5)

	workflow.ExecuteActivity(ctx, activities.PrintCurrentTime, "")

	workflow.Sleep(ctx, time.Second*5)

	fmt.Println("Ended workflow: ActivityWorkflow")
	fmt.Println("------------------------------------------------------------------")

	return nil
}

func VersionWorkflow(ctx workflow.Context) error {
	fmt.Println("Started workflow: VersionWorkflow")
	workflow.Sleep(ctx, time.Minute*3)
	fmt.Println("New message")
	fmt.Println("Ended workflow: VersionWorkflow")
	fmt.Println("------------------------------------------------------------------")
	return nil
}
