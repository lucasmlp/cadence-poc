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

const (
	serviceNameCadenceClient   = "cadence-client"
	serviceNameCadenceFrontend = "cadence-frontend"
	domainName                 = "poc"
)

func main() {

	action := os.Args[1]

	workflowClient, err := helpers.NewWorkflowClient(serviceNameCadenceClient, serviceNameCadenceFrontend)

	if err != nil {
		panic(err)
	}

	triggerClient := helpers.NewCadenceClient(workflowClient)

	workflowID := uuid.New()

	switch name := action; name {
	case "StartHelloWorldWorkflow":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 1 * time.Second,
		}, workflows.HelloWorldWorkflow)

	case "StartSimpleWorkflow":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 45 * time.Second,
		}, workflows.SimpleWorkflow)

	case "StartWaitingSignalWorkflow":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 30 * time.Second,
		}, workflows.WaitingSignalWorkflow, "signalTeste")

	case "SendCancelSignalWorkflow":
		err = triggerClient.SignalWorkflow(context.Background(), os.Args[2], "", "signalTeste", "cancel")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Started Action: ", workflowID)
	fmt.Println("Action: ", action)
}
