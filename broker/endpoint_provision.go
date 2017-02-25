package broker

import (
	"context"
	"fmt"

	etcdclient "github.com/coreos/etcd/client"
	"github.com/hashicorp/errwrap"
	"github.com/pivotal-cf/brokerapi"
)

// Provision a new service instance
// 1. Create role
// 2. Grant path access to role
// If they all already existed, then provisioning fails
//
// Binding will:
// 3. Create user
// 4. Assign role to user
func (bkr *Broker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (resp brokerapi.ProvisionedServiceSpec, err error) {
	roleName := bkr.serviceInstanceRole(instanceID)
	rolePaths := []string{
		fmt.Sprintf("%s/*", bkr.serviceInstanceKeyPath(instanceID)),
	}

	authRoleAPI := etcdclient.NewAuthRoleAPI(bkr.EtcdClient)
	err = authRoleAPI.AddRole(ctx, roleName)
	if err != nil {
		err = errwrap.Wrapf("Could not add role: {{err}}", err)
		return
	}

	grantedRole, err := authRoleAPI.GrantRoleKV(ctx, roleName, rolePaths, etcdclient.ReadWritePermission)
	if err != nil {
		err = errwrap.Wrapf("Could not grant role: {{err}}", err)
		return
	}
	fmt.Printf("Created role %v\n", grantedRole)
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (bkr *Broker) serviceInstanceRole(instanceID string) string {
	return fmt.Sprintf("instance-%s", instanceID)
}

func (bkr *Broker) serviceInstanceKeyPath(instanceID string) string {
	return fmt.Sprintf("/service_instances/%s", instanceID)
}
