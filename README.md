# Service Broker for CoreOS Etcd

This project gives users of a Cloud Foundry installation access to slices of a shared CoreOS etcd deployment.

This project does not describe the deployment of etcd. Consider https://github.com/cloudfoundry-incubator/etcd-release/ as one way to bring up single- and multi-node etcd clusters with or without https support.

## Usage

Users will access this service broker via their `cf` CLI using the standard commands for `cf create-service`, `cf bind-service`, `cf unbind-service` and `cf delete-service`; as well as `cf create-service-key`, etc.

Bound applications/service keys will receive credentials that look similar to:

```json
{
  "credentials": {
    "uri": "http://user-8c0d5822-f806-11e6-84f6-a79e3c5fe739:WVmH8Qr365@23.159.100.202:4001",
    "keypath": "/service_instances/92122d42-f806-11e6-bcb2-4b77b9de2108",
    "host": "http://23.159.100.202:4001",
    "username": "user-8c0d5822-f806-11e6-84f6-a79e3c5fe739",
    "password": "WVmH8Qr365"
  }
}
```

## Deployment

The broker is a Golang application that can be deployed to Cloud Foundry.

It requires the following environment variables:

* `ETCD_URI` - full URI including basic auth for a user with `root` auth permissions (to create other users/roles)

Optional env vars:

* `BROKER_USERNAME` and `BROKER_USERNAME` to configure the service broker's basic auth (both default to `broker`)
* `PORT` is the port to which the application will bind for incoming HTTP requests (defaults to `6000`)

## Performance

There is a performance hit to etcd when you enable etcd auth to use this service broker on your etcd cluster. It was discovered by the compose.io team that etcd with auth enabled is slower for every request due to the use of `bcrypt` within etcd.

https://www.compose.com/articles/the-mystery-of-etcds-rationed-requests/

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
...
provision-bind-use | deleting instance etcd-4631667841
provision-bind-use | + curl -f -XDELETE 'http://broker:password@54.159.121.202:6000/v2/service_instances/etcd-4631667841?service_id=5b0ad2fe-f7c0-11e6-8e76-7fc33eaeccd4&plan_id=5bcfa502-f7c0-11e6-bd06-e323138af97b'
provision-bind-use | {}
provision-bind-use | + set +x
provision-bind-use | User's /hello should now be deleted
Stopping 'provision-read-write' Runtime...
Stopping etcd-cf-service-broker ... done
Stopping etcd ... done
Test 'provision-read-write' completed sucessfully!
```
