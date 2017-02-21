FROM alpine:3.5

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN apk add --no-cache --update gcc g++ curl bash sed jq go git

EXPOSE 6000
COPY tests/scripts /scripts
CMD ["/scripts/start_broker.sh"]

COPY . /go/src/github.com/cloudfoundry-community/etcd-cf-service-broker
RUN set -x \
    && go install github.com/cloudfoundry-community/etcd-cf-service-broker
