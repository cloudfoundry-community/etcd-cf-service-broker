package broker

import "github.com/frodenas/brokerapi"

// Services is the catalog of services offered by the broker
func (bkr *Broker) Services() brokerapi.CatalogResponse {
	return brokerapi.CatalogResponse{}
}

// Provision a new service instance
func (bkr *Broker) Provision(instanceID string, details brokerapi.ProvisionDetails, acceptsIncomplete bool) (resp brokerapi.ProvisioningResponse, async bool, err error) {
	return brokerapi.ProvisioningResponse{}, false, nil
}

// Update service instance
func (bkr *Broker) Update(instanceID string, updateDetails brokerapi.UpdateDetails, acceptsIncomplete bool) (async bool, err error) {
	return false, nil
}

// Bind returns access credentials for a service instance
func (bkr *Broker) Bind(instanceID string, bindingID string, details brokerapi.BindDetails) (brokerapi.BindingResponse, error) {
	return brokerapi.BindingResponse{
		Credentials: map[string]interface{}{},
	}, nil
}

// Unbind to remove access to service instance
func (bkr *Broker) Unbind(instanceID string, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

// Deprovision service instance
func (bkr *Broker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, acceptsIncomplete bool) (async bool, err error) {
	return false, nil
}

// LastOperation returns the status of the last operation on a service instance
func (bkr *Broker) LastOperation(instanceID string) (resp brokerapi.LastOperationResponse, err error) {
	return brokerapi.LastOperationResponse{
		State:       brokerapi.LastOperationSucceeded,
		Description: "Succeeded",
	}, nil
}
