package broker

import (
	"context"
	"fmt"

	etcdclient "github.com/coreos/etcd/client"
	"github.com/hashicorp/errwrap"
	"github.com/pivotal-cf/brokerapi"
)

// Unbind to remove access to service instance
func (bkr *Broker) Unbind(ctx context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails) (err error) {
	authUserAPI := etcdclient.NewAuthUserAPI(bkr.EtcdClient)
	username := bkr.serviceBindingUser(bindingID)
	err = authUserAPI.RemoveUser(ctx, username)
	if err != nil {
		err = errwrap.Wrapf("Could not remove user: {{err}}", err)
		return
	}
	fmt.Printf("Deleted user %v\n", username)
	return
}
