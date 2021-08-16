package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/lucasmachadolopes/cadencePoc/workflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

const (
	serviceNameCadenceClient   = "cadence-client"
	serviceNameCadenceFrontend = "cadence-frontend"
)

func NewWorkflowClient() (workflowserviceclient.Interface, error) {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(serviceNameCadenceClient))
	if err != nil {
		return nil, err
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: serviceNameCadenceClient,
		Outbounds: yarpc.Outbounds{
			serviceNameCadenceFrontend: {Unary: ch.NewSingleOutbound("127.0.0.1:7933")},
		},
	})

	if dispatcher == nil {
		return nil, errors.New("failed to create dispatcher")
	}

	if err := dispatcher.Start(); err != nil {
		panic(err)
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(serviceNameCadenceFrontend)), nil
}

func NewCadenceClient(workflowClient workflowserviceclient.Interface) client.Client {
	return client.NewClient(workflowClient, "poc", &client.Options{})
}

func main() {

	workflowName := os.Args[1]

	wfClient, err := NewWorkflowClient()

	if err != nil {
		panic(err)
	}

	triggerClient := NewCadenceClient(wfClient)

	workflowID := uuid.New()

	switch name := workflowName; name {
	case "HelloWorld":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 3 * time.Second,
		}, workflows.HelloWorldWorkflow)

	case "SimpleWorkflow":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 10 * time.Second,
		}, workflows.SimpleWorkflow)

	case "WaitingSignalWorkflow":
		_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:                           workflowID,
			TaskList:                     "pocTasklist",
			ExecutionStartToCloseTimeout: 1 * time.Minute,
		}, workflows.WaitingSignalWorkflow, "signalTeste")

	case "SendSignalWorkflow":
		err = triggerClient.SignalWorkflow(context.Background(), os.Args[2], "", "signalTeste", "SOME_VALUE")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Started workflow: ", workflowID)
	fmt.Println("Workflow: ", workflowName)
}
