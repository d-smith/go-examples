## Demo Tags

* Config Binding - cfgBinding
* Service Binding - svcBinding

## Consul Setup

* Native Mac OS X docker - pull images while disconnect from the FMN
* Need to specify a client-addr to expose the UI outside the container

<pre>
docker run -p 8500:8500 consul agent -dev -client=0.0.0.0
</pre>

* Dev cluster

<pre>
docker run -d --name=dev-consul -p 8500:8500 consul agent -dev -client=0.0.0.0
docker run -d consul agent -dev -join=172.17.0.2
docker run -d consul agent -dev -join=172.17.0.2
</pre>

* Cluster members

<pre>
docker exec -t dev-consul consul members
</pre>

## Mountebank Setup

<pre>
curl -i -X POST -d@endpoint-setup.json http://127.0.0.1:2525/imposters
curl -i -X POST -d@endpoint2-setup.json http://127.0.0.1:2525/imposters
</pre>

## Run with environment variables

<pre>
docker run -e "endpoint=foo:4545" -e "port=3000" --link mountebank:foo -p 3000:3000  1ac129181e49
</pre>

## Golang Setup for Demo

* golang set up - consul.sh

## Consul Env

* Install go get github.com/mitchellh/gox, clone https://github.com/hashicorp/envconsul,
and make bin

<pre>
./envconsul  -consul=localhost:8500 -once -prefix=sample -pristine -upcase env
</pre>


## Consul Template

* Project - https://github.com/hashicorp/consul-template
* Releases - https://releases.hashicorp.com/consul-template/

## Docker Image

Build thusly:

<pre>
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o service
docker build -t cfgsample .
</pre>

## Add config

<pre>
curl -X PUT localhost:8500/v1/kv/env1/
curl -X PUT localhost:8500/v1/kv/env1/endpoint-host -d foo
curl -X PUT localhost:8500/v1/kv/env1/endpoint-port -d 4545
curl -X PUT localhost:8500/v1/kv/env1/port -d 3000
</pre>

## Consul-template

<pre>
./consul-template -consul localhost:8500 -template /Users/a045103/goex/src/github.com/d-smith/go-examples/consul-template/demo-template.ctmpl -dry -once
</pre>

## Run mountebank

<pre>
docker pull dasmith/mb-server-alpine
docker run -d -p 2525:2525 --name mountebank dasmith/mb-server-alpine
</pre>


