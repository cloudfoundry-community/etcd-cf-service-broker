package broker

import (
	"context"

	etcdclient "github.com/coreos/etcd/client"
	"github.com/pivotal-cf/brokerapi"
)

// Provision a new service instance
func (bkr *Broker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (resp brokerapi.ProvisionedServiceSpec, err error) {
	authRoleAPI := etcdclient.NewAuthRoleAPI(bkr.EtcdClient)

	authRoleAPI.ListRoles(ctx)
	return brokerapi.ProvisionedServiceSpec{}, nil
}
