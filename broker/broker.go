package broker

import "github.com/pivotal-golang/lager"

// Broker holds config for Etcd service broker API endpoints
type Broker struct {
	Logger lager.Logger
}

// NewBroker constructs Broker
func NewBroker(logger lager.Logger) (bkr *Broker, err error) {
	bkr = &Broker{
		Logger: logger,
	}
	return
}
