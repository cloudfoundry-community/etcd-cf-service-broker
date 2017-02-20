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
