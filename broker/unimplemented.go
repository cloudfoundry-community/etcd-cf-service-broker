package broker

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

// Deprovision service instance
func (bkr *Broker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (resp brokerapi.DeprovisionServiceSpec, err error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

// Update service instance
func (bkr *Broker) Update(ctx context.Context, instanceID string, updateDetails brokerapi.UpdateDetails, asyncAllowed bool) (resp brokerapi.UpdateServiceSpec, err error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

// LastOperation returns the status of the last operation on a service instance
func (bkr *Broker) LastOperation(ctx context.Context, instanceID string, operationData string) (resp brokerapi.LastOperation, err error) {
	return brokerapi.LastOperation{
		State:       brokerapi.Succeeded,
		Description: "Succeeded",
	}, nil
}
