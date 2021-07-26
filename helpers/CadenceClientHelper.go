package helpers

import (
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
)

func NewCadenceClient(workflowClient workflowserviceclient.Interface) client.Client {
	return client.NewClient(workflowClient, "cadence-poc", &client.Options{})
}
