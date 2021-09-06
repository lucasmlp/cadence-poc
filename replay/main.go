package main

import (
	"context"
	"os"

	"github.com/machado-br/cadence-poc/helpers"
	"github.com/machado-br/cadence-poc/workflows"
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

	replayer := worker.NewWorkflowReplayer()

	replayer.RegisterWorkflowWithOptions(workflows.VersionWorkflow, workflow.RegisterOptions{
		Name: "VersionWorkflow",
	})

	execution := workflow.Execution{
		ID:    "",
		RunID: "",
	}
	// jsonFileName := "95796113 44a9 41f5 ae9f aeee7d1586b4 - 44824689-c10c-4250-b944-42e859ad1772.json"

	// if workflow history has been loaded into memory
	// err := replayer.ReplayWorkflowHistory(logger, history)

	// if workflow history is stored in a json file
	// err = replayer.ReplayWorkflowHistoryFromJSONFile(logger, jsonFileName)

	// if workflow history is stored in a json file and you only want to replay part of it
	// NOTE: lastEventID can't be set arbitrarily. It must be the end of of a history events batch
	// when in doubt, set to the eventID of decisionTaskStarted events.
	// err = replayer.ReplayPartialWorkflowHistoryFromJSONFile(logger, jsonFileName, lastEventID)

	// if you want to fetch workflow history directly from cadence server
	// please check the Worker Service page for how to create a cadence service client
	err = replayer.ReplayWorkflowExecution(context.Background(), workflowClient, logger, domainName, execution)

}
