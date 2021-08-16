package workflows

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func HelloWorldWorkflow(ctx workflow.Context) error {
	fmt.Println("Hello world")
	fmt.Println("------------------------------------------------------------------")
	return nil
}

func SimpleWorkflow(ctx workflow.Context) error {

	fmt.Println("Started workflow: SimpleWorkflow")
	workflow.Sleep(ctx, time.Second*5)
	fmt.Println("Ended workflow: SimpleWorkflow")
	fmt.Println("------------------------------------------------------------------")

	return nil
}

func WaitingSignalWorkflow(ctx workflow.Context, signalName string) error {
	fmt.Println("Started workflow: WaitingSignalWorkflow")

	var signalVal string
	signalChan := workflow.GetSignalChannel(ctx, signalName)

	s := workflow.NewSelector(ctx)
	s.AddReceive(signalChan, func(c workflow.Channel, more bool) {
		c.Receive(ctx, &signalVal)
		workflow.GetLogger(ctx).Info("Received signal!", zap.String("signal", signalName), zap.String("value", signalVal))
	})
	s.Select(ctx)

	if len(signalVal) > 0 && signalVal != "SOME_VALUE" {
		return errors.New("signalVal")
	}

	fmt.Println("Ended workflow: WaitingSignalWorkflow")
	fmt.Println("------------------------------------------------------------------")
	return nil
}
