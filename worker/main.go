package main

import (
	"github.com/lucasmachadolopes/cadencePoc/activities"
	"github.com/lucasmachadolopes/cadencePoc/helpers"
	"github.com/lucasmachadolopes/cadencePoc/workflows"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	serviceNameCadenceClient   = "cadence-client"
	serviceNameCadenceFrontend = "cadence-frontend"
	domainName                 = "poc"
)

func main() {

	workflowClient, err := helpers.NewWorkflowClient(serviceNameCadenceClient, serviceNameCadenceFrontend)
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	w := worker.New(workflowClient, domainName, "pocTasklist",
		worker.Options{
			Logger: logger,
		})
	workflow.RegisterWithOptions(workflows.HelloWorldWorkflow, workflow.RegisterOptions{
		Name: "HelloWorldWorkflow",
	})
	workflow.RegisterWithOptions(workflows.SimpleWorkflow, workflow.RegisterOptions{
		Name: "SimpleWorkflow",
	})
	workflow.RegisterWithOptions(workflows.WaitingSignalWorkflow, workflow.RegisterOptions{
		Name: "WaitingSignalWorkflow",
	})
	activity.Register(activities.PrintCurrentTime)

	err = w.Run()

	if err != nil {
		panic(err)
	}
}
