package broker

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/cloudfoundry-community/etcd-cf-service-broker/assets"
	"github.com/pivotal-cf/brokerapi"
)

// Catalog describes a set of CF Services
type Catalog struct {
	Services []brokerapi.Service `json:"services"`
}

// Services is the catalog of services offered by the broker
func (bkr *Broker) Services(context context.Context) []brokerapi.Service {
	defaultCatalog := bytes.NewBuffer(assets.MustAsset("data/default_catalog.json"))

	catalog := &Catalog{}
	dec := json.NewDecoder(defaultCatalog)
	if err := dec.Decode(catalog); err != nil {
		log.Fatal(err)
	}

	return catalog.Services
}
