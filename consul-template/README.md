

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
<pre>

* Cluster members

<pre>
docker exec -t dev-consul consul members
</pre>

## Mountebank Setup

<pre>
curl -i -X POST -d@endpoint-setup.json http://127.0.0.1:2525/imposters
curl -i -X POST -d@endpoint2-setup.json http://127.0.0.1:2525/imposters
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



