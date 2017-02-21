package broker

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cloudfoundry-community/etcd-cf-service-broker/utils"
	etcdclient "github.com/coreos/etcd/client"
	"github.com/hashicorp/errwrap"
	"github.com/pivotal-cf/brokerapi"
)

// EtcdCredentials is the set of credentials passed back to user to access their
// slice of etcd cluster
type EtcdCredentials struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	BasePath string `json:"base_path"`
	URI      string `json:"uri"`
}

// Bind returns access credentials for a service instance
//
// Previously, provisioning a new service instance:
// 1. Create role
// 2. Grant path access to role
// If they all already existed, then provisioning fails
//
// Now, for each binding:
// 3. Create user
// 4. Assign the service instance role to user
func (bkr *Broker) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails) (binding brokerapi.Binding, err error) {
	authUserAPI := etcdclient.NewAuthUserAPI(bkr.EtcdClient)
	username := fmt.Sprintf("user-%s", bindingID)
	password := utils.NewPassword(10)
	err = authUserAPI.AddUser(ctx, username, password)
	if err != nil {
		err = errwrap.Wrapf("Could not add user: {{err}}", err)
		return
	}

	serviceInstanceRoles := []string{bkr.serviceInstanceRole(instanceID)}
	user, err := authUserAPI.GrantUser(ctx, username, serviceInstanceRoles)
	if err != nil {
		err = errwrap.Wrapf("Could not assign user to role: {{err}}", err)
		return
	}
	fmt.Printf("Created user %v\n", user)

	basePath := bkr.serviceInstancePath(instanceID)
	serviceInstanceURL := fmt.Sprintf("%s%s", bkr.etcdBaseURL(), basePath)
	u, err := url.Parse(serviceInstanceURL)
	if err != nil {
		err = errwrap.Wrapf(fmt.Sprintf("Could not parse URL %s: {{err}}", serviceInstanceURL), err)
		return
	}
	uri := fmt.Sprintf("%s://%s:%s@%s%s", u.Scheme, username, password, u.Host, u.Path)

	creds := EtcdCredentials{
		URI:      uri,
		Host:     bkr.etcdBaseURL(),
		Username: username,
		Password: password,
		BasePath: basePath,
	}
	return brokerapi.Binding{
		Credentials: creds,
	}, nil
}
