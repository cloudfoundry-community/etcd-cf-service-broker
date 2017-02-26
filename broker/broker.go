package broker

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"code.cloudfoundry.org/lager"
	etcdclient "github.com/coreos/etcd/client"
)

// Broker holds config for Etcd service broker API endpoints
type Broker struct {
	Logger     lager.Logger
	EtcdClient etcdclient.Client
}

// NewBroker constructs Broker
func NewBroker(logger lager.Logger) (bkr *Broker, err error) {
	bkr = &Broker{
		Logger: logger,
	}
	bkr.setupEtcdClient()
	return
}

func (bkr *Broker) etcdBaseURL() string {
	return bkr.EtcdClient.Endpoints()[0]
}

func (bkr *Broker) setupEtcdClient() {
	etcdURI := os.Getenv("ETCD_URI")
	if etcdURI == "" {
		fmt.Fprintf(os.Stderr, "Require $ETCD_URI\n")
		os.Exit(1)
	}
	endpoint, err := url.Parse(etcdURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse $ETCD_URI: %s\n", etcdURI)
		os.Exit(1)
	}
	user := endpoint.User
	password, _ := user.Password()
	endpoint.User = nil

	cfg := etcdclient.Config{
		Endpoints: []string{endpoint.String()},
		Transport: etcdclient.DefaultTransport,
		Username:  user.Username(),
		Password:  password,
	}
	ctx := context.Background()

	c, err := etcdclient.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to server: %s\n", err)
		os.Exit(1)
	}
	bkr.EtcdClient = c

	etcdclient.EnablecURLDebug()

	fmt.Println("\nEnabling auth, if not already enabled...")
	authAPI := etcdclient.NewAuthAPI(bkr.EtcdClient)
	err = authAPI.Enable(ctx)
	if err != nil {
		fmt.Printf("%s... continuing...\n", err)
	}

	fmt.Println("\nList existing auth users...")
	authUserAPI := etcdclient.NewAuthUserAPI(bkr.EtcdClient)
	users, err := authUserAPI.ListUsers(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Old etcd warning: Failed to get existing auth users: %s\n", err)
	} else {
		fmt.Printf("%#v\n\n", users)
	}

	fmt.Println("\nList existing auth roles...")
	authRoleAPI := etcdclient.NewAuthRoleAPI(bkr.EtcdClient)
	roles, err := authRoleAPI.ListRoles(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Old etcd warning: Failed to get existing auth roles: %s\n", err)
	} else {
		fmt.Printf("%#v\n\n", roles)
	}
}
