package broker

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

// Provision a new service instance
func (bkr *Broker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (resp brokerapi.ProvisionedServiceSpec, err error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

// Deprovision service instance
func (bkr *Broker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (resp brokerapi.DeprovisionServiceSpec, err error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

// Bind returns access credentials for a service instance
func (bkr *Broker) Bind(context context.Context, instanceID string, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{
		Credentials: map[string]interface{}{},
	}, nil
}

// Unbind to remove access to service instance
func (bkr *Broker) Unbind(context context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

// Update service instance
func (bkr *Broker) Update(context context.Context, instanceID string, updateDetails brokerapi.UpdateDetails, asyncAllowed bool) (resp brokerapi.UpdateServiceSpec, err error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

// LastOperation returns the status of the last operation on a service instance
func (bkr *Broker) LastOperation(context context.Context, instanceID string, operationData string) (resp brokerapi.LastOperation, err error) {
	return brokerapi.LastOperation{
		State:       brokerapi.Succeeded,
		Description: "Succeeded",
	}, nil
}
