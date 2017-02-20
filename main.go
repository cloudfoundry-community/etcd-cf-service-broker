package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-community/etcd-cf-service-broker/broker"
	"github.com/pivotal-cf/brokerapi"
)

func main() {
	rand.Seed(5000)
	run()
}

func run() {
	logger := setupLogger()
	credentials := brokerapi.BrokerCredentials{
		Username: os.Getenv("BROKER_USERNAME"),
		Password: os.Getenv("BROKER_PASSWORD"),
	}
	if credentials.Username == "" {
		fmt.Fprintf(os.Stderr, "Require $BROKER_USERNAME, defaulting 'broker'\n")
		credentials.Username = "broker"
	}
	if credentials.Password == "" {
		fmt.Fprintf(os.Stderr, "Require $BROKER_PASSWORD, defaulting 'broker'\n")
		credentials.Password = "broker"
	}
	portStr := os.Getenv("PORT")
	if portStr == "" {
		fmt.Fprintf(os.Stderr, "Require $PORT, defaulting 6000\n")
		portStr = "6000"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "$PORT must be an integer, was '%s'\n", portStr)
	}

	bkr, _ := broker.NewBroker(logger)
	brokerAPI := brokerapi.New(bkr, logger, credentials)
	http.Handle("/v2/", brokerAPI)

	fmt.Printf("Running on :%d\n", port)
	logger.Fatal("http-listen", http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

func setupLogger() lager.Logger {
	logger := lager.NewLogger("etcd-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))
	return logger
}
