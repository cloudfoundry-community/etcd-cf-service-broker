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

	c, err := etcdclient.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to server: %s\n", err)
		os.Exit(1)
	}
	bkr.EtcdClient = c

	fmt.Println("Test credentials by listing member nodes...")
	m := etcdclient.NewMembersAPI(c)
	nodes, err := m.List(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to access etcd server: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n\n", nodes)
}
