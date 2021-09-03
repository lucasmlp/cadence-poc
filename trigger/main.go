package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lucasmachadolopes/cadencePoc/helpers"
	"github.com/lucasmachadolopes/cadencePoc/workflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
)

func main() {
	serviceNameCadenceClient := os.Getenv("CADENCE_CLIENT_NAME")
	serviceNameCadenceFrontend := os.Getenv("CADENCE_FRONTEND_NAME")

	action := os.Args[1]

	workflowClient, err := helpers.NewWorkflowClient(serviceNameCadenceClient, serviceNameCadenceFrontend)

	if err != nil {
		panic(err)
	}

	triggerClient := helpers.NewCadenceClient(workflowClient)

	workflowID := uuid.New()

	switch name := action; name {
	case "HelloWorld":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 1 * time.Second,
		}, workflows.HelloWorldWorkflow)

	case "Activity":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 45 * time.Second,
		}, workflows.ActivityWorkflow)

	case "WaitingSignal":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{}, workflows.WaitingSignalWorkflow, "signalTeste")

	case "Version":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 6 * time.Minute,
		}, workflows.VersionWorkflow)

	case "CancelSignal":
		err = triggerClient.SignalWorkflow(context.Background(), os.Args[2], "", "signalTeste", "cancel")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Started Action: ", workflowID)
	fmt.Println("Action: ", action)
}
