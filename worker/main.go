package main

import (
	"fmt"
	"os"

	"github.com/lucasmachadolopes/cadencePoc/activities"
	"github.com/lucasmachadolopes/cadencePoc/helpers"
	"github.com/lucasmachadolopes/cadencePoc/workflows"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func main() {
	serviceNameCadenceClient := os.Getenv("CADENCE_CLIENT_NAME")
	serviceNameCadenceFrontend := os.Getenv("CADENCE_FRONTEND_NAME")
	domainName := os.Getenv("CADENCE_DOMAIN_NAME")

	fmt.Printf("serviceNameCadenceClient: %v\n", serviceNameCadenceClient)

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
	workflow.RegisterWithOptions(workflows.WaitingSignalWorkflow, workflow.RegisterOptions{
		Name: "WaitingSignalWorkflow",
	})
	workflow.RegisterWithOptions(workflows.ActivityWorkflow, workflow.RegisterOptions{
		Name: "ActivityWorkflow",
	})
	workflow.RegisterWithOptions(workflows.VersionWorkflow, workflow.RegisterOptions{
		Name: "VersionWorkflow",
	})
	workflow.RegisterWithOptions(workflows.VersionWorkflow2, workflow.RegisterOptions{
		Name: "VersionWorkflow2",
	})
	activity.Register(activities.PrintCurrentTime)
	activity.Register(activities.ActivityA)
	activity.Register(activities.ActivityB)
	activity.Register(activities.ActivityC)

	err = w.Run()

	if err != nil {
		panic(err)
	}
}
