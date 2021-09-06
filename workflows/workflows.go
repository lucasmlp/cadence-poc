package workflows

import (
	"fmt"
	"time"

	"github.com/machado-br/cadence-poc/activities"

	"go.uber.org/cadence/workflow"
)

func HelloWorldWorkflow(ctx workflow.Context) error {
	fmt.Println("Hello world")
	fmt.Println("------------------------------------------------------------------")

	return nil
}

func WaitingSignalWorkflow(ctx workflow.Context, signalName string) error {
	fmt.Println("Started workflow: WaitingSignalWorkflow")
	internalContext, cancelfunc := workflow.WithCancel(ctx)

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

func VersionWorkflow(ctx workflow.Context, workflowId string) error {
	fmt.Println("Started workflow: VersionWorkflow")

	// workflow.Sleep(ctx, time.Minute*2)
	// fmt.Println("New message")
	// fmt.Println("Ended workflow: ", workflowId)
	// fmt.Println("------------------------------------------------------------------")

	v := workflow.GetVersion(ctx, "Step1", workflow.DefaultVersion, 1)
	if v == workflow.DefaultVersion {
		workflow.Sleep(ctx, time.Minute*2)
		fmt.Println("New message")
		fmt.Println("Ended workflow: ", workflowId)
		fmt.Println("------------------------------------------------------------------")
	} else {
		workflow.Sleep(ctx, time.Minute*1)
		fmt.Println("New message")
		fmt.Println("Ended new workflow: ", workflowId)
		fmt.Println("------------------------------------------------------------------")
	}

	return nil
}

func VersionWorkflow2(ctx workflow.Context, data string) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	var err error
	var result1 string

	// workflow.Sleep(ctx, time.Minute*2)
	// err = workflow.ExecuteActivity(ctx, activities.ActivityA, data).Get(ctx, &result1)

	v := workflow.GetVersion(ctx, "Step1", workflow.DefaultVersion, 1)
	if v == workflow.DefaultVersion {
		workflow.Sleep(ctx, time.Minute*2)
		err = workflow.ExecuteActivity(ctx, activities.ActivityA, data).Get(ctx, &result1)
	} else {
		workflow.Sleep(ctx, time.Minute*1)
		err = workflow.ExecuteActivity(ctx, activities.ActivityC, data).Get(ctx, &result1)
	}

	if err != nil {
		return err
	}
	fmt.Println(result1)
	var result2 string
	err = workflow.ExecuteActivity(ctx, activities.ActivityB, result1).Get(ctx, &result2)
	if err != nil {
		return err
	}

	fmt.Println(result2)
	return err
}
