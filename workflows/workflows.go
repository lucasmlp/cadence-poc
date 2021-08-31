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

func SimpleWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	workflowInfo := workflow.GetInfo(ctx)
	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    time.Second * 30,
		ScheduleToStartTimeout: time.Second * 120,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	fmt.Println("Started workflow: SimpleWorkflow")
	workflow.Sleep(ctx, time.Second*5)
	var returnedError error

	if err := workflow.ExecuteActivity(ctx, activities.PrintCurrentTime).Get(ctx, &returnedError); err != nil {
		logger.Error(
			"Failed to get log-id in notifications-service",
			zap.String("WorkflowID", workflowInfo.WorkflowExecution.ID),
			zap.Int("ChannelID", 5),
			zap.Error(err),
		)
		return err
	}

	workflow.Sleep(ctx, time.Second*5)
	fmt.Println("Ended workflow: SimpleWorkflow")
	fmt.Println("------------------------------------------------------------------")

	return nil
}

func WaitingSignalWorkflow(externalContext workflow.Context, signalName string) error {
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
