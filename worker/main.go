package main

import (
	"github.com/lucasmachadolopes/cadencePoc/helpers"
	"github.com/lucasmachadolopes/cadencePoc/workflows"

	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

const (
	serviceNameCadenceClient   = "cadence-client"
	serviceNameCadenceFrontend = "cadence-frontend"
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

	w := worker.New(workflowClient, "cadence-poc", "pocTasklist",
		worker.Options{
			Logger: logger,
		})

	w.RegisterWorkflow(workflows.HelloWorldWorkflow)

	err = w.Run()

	if err != nil {
		panic(err)
	}
}
