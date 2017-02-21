# Service Broker for CoreOS Etcd

This project gives users of a Cloud Foundry installation access to slices of a shared CoreOS etcd deployment.

This project does not describe the deployment of etcd. Consider https://github.com/cloudfoundry-incubator/etcd-release/ as one way to bring up single- and multi-node etcd clusters with or without https support.


## Development

If any new files or changes are made to `data/` directory, then need to run `go-bindata` to update `assets/bindata.go`:

```
go get -u github.com/jteeuwen/go-bindata/...
go-bindata -pkg assets -o assets/bindata.go ./data/...
```

Access the assets using:

```go
contents, err := assets.Asset("data/filename")
```

## Tests

Since this broker interacts with etcd, the integration test suite uses `docker-compose` to launch the broker and an etcd server. In addition, we use [`delmo`](https://github.com/bodymindarts/delmo) to run tests against the broker and its provisioned etcd service instances.

```
go get -u github.com/bodymindarts/delmo
delmo -m <docker-machine-name>
```

This will build two Docker images: one for the broker (see `Dockerfile`) and one for the test scripts (see `tests/Dockerfile`), then run the cluster of containers and iterate through the `delmo.yml` test cases.

The output might look like:

```
Running test 'provision-read-write'...
Starting 'provision-read-write' Runtime...
Creating network "provisionreadwrite_default" with the default driver
Creating etcd
Creating etcd-cf-service-broker
Creating provisionreadwrite_tests_1
Executing - <Exec: show-catalog>
show-catalog | + curl -s http://broker:password@54.159.121.202:6000/v2/catalog
show-catalog | {"services":[{"id":"5b0ad2fe-f7c0-11e6-8e76-7fc33eaeccd4","name":"etcd","description":"Etcd as a service","bindable":true,"tags":["etcd","etcd2","keyvalue"],"plan_updateable":false,"plans":[{"id":"5bcfa502-f7c0-11e6-bd06-e323138af97b","name":"shared","description":"Shared slice of etcd cluster","metadata":{"displayName":"Shared","bullets":["etcd v2"]}}],"metadata":{"displayName":"displayname","longDescription":"Distributed reliable key-value store for the most critical data of a distributed system","providerDisplayName":"Stark \u0026 Wayne","supportUrl":"https://github.com/cloudfoundry-community/etcd-cf-service-broker/issues"}}]}
Stopping 'provision-read-write' Runtime...
Stopping etcd-cf-service-broker ... done
Stopping etcd ... done
Test 'provision-read-write' completed sucessfully!
```
