package broker

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/cloudfoundry-community/etcd-cf-service-broker/assets"
	"github.com/frodenas/brokerapi"
)

// Services is the catalog of services offered by the broker
func (bkr *Broker) Services() brokerapi.CatalogResponse {
	defaultCatalog := bytes.NewBuffer(assets.MustAsset("data/default_catalog.json"))

	catalog := &brokerapi.CatalogResponse{}
	dec := json.NewDecoder(defaultCatalog)
	if err := dec.Decode(catalog); err != nil {
		log.Fatal(err)
	}

	return *catalog
}
