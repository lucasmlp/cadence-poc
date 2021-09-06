package main

import (
	"os"

	"github.com/machado-br/cadence-poc/activities"
	"github.com/machado-br/cadence-poc/helpers"
	"github.com/machado-br/cadence-poc/workflows"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func main() {
	serviceNameCadenceClient := os.Getenv("CADENCE_CLIENT_NAME")
	serviceNameCadenceFrontend := os.Getenv("CADENCE_FRONTEND_NAME")
	domainName := os.Getenv("CADENCE_DOMAIN_NAME")

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

	workflow.Register(workflows.HelloWorldWorkflow)
	workflow.Register(workflows.WaitingSignalWorkflow)
	workflow.Register(workflows.ActivityWorkflow)
	workflow.Register(workflows.VersionWorkflow)
	workflow.Register(workflows.VersionWorkflow2)

	activity.Register(activities.PrintCurrentTime)
	activity.Register(activities.ActivityA)
	activity.Register(activities.ActivityB)
	activity.Register(activities.ActivityC)

	err = w.Run()

	if err != nil {
		panic(err)
	}
}
