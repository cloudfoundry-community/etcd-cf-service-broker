package broker

import (
	"context"
	"fmt"

	etcdclient "github.com/coreos/etcd/client"
	"github.com/hashicorp/errwrap"
	"github.com/pivotal-cf/brokerapi"
)

// Deprovision service instance
func (bkr *Broker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (resp brokerapi.DeprovisionServiceSpec, err error) {
	roleName := bkr.serviceInstanceRole(instanceID)

	authRoleAPI := etcdclient.NewAuthRoleAPI(bkr.EtcdClient)
	err = authRoleAPI.RemoveRole(ctx, roleName)
	if err != nil {
		err = errwrap.Wrapf("Could not remove role: {{err}}", err)
		return
	}

	keysAPI := etcdclient.NewKeysAPI(bkr.EtcdClient)
	deleteOpts := &etcdclient.DeleteOptions{
		Recursive: true,
		Dir:       true,
	}
	deletePath := bkr.serviceInstanceKeyPath(instanceID)
	_, err = keysAPI.Delete(ctx, deletePath, deleteOpts)
	if err != nil {
		err = errwrap.Wrapf(fmt.Sprintf("Could not recursive delete '%s': {{err}}", deletePath), err)
		return
	}
	return brokerapi.DeprovisionServiceSpec{}, nil
}
