package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lucasmachadolopes/cadencePoc/helpers"
	"github.com/lucasmachadolopes/cadencePoc/workflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
)

const (
	serviceNameCadenceClient   = "cadence-client"
	serviceNameCadenceFrontend = "cadence-frontend"
)

func main() {
	wfClient, err := helpers.NewWorkflowClient(serviceNameCadenceClient, serviceNameCadenceFrontend)
	if err != nil {
		panic(err)
	}

	triggerClient := helpers.NewCadenceClient(wfClient)

	workflowID := uuid.New()

	_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:                           workflowID,
		TaskList:                     "pocTasklist",
		ExecutionStartToCloseTimeout: 3 * time.Second,
	}, workflows.HelloWorldWorkflow)

	if err != nil {
		panic(err)
	}

	fmt.Println("Started workflow:", workflowID)
}
