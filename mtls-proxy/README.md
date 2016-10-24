# MTLS Proxy

This project shows how to set up a reverse caching proxy in front of a
golang service where Mutual TLS is required to access to proxied server.

The pieces:

* The http server running in a docker container
* The proxy server running in a docker container
* The client that calls the server through the proxy

The server only allows requests from a known client.

## Setup

To run the example, we need to set up the server, proxy, and client

### Root CA

<pre>
mkdir openssl
cd openssl/
mkdir ca
openssl genrsa -aes256 -out ca/ca.key 4096 chmod 400 ca/ca.key

openssl req -new -x509 -sha256 -days 730 -key ca/ca.key -out ca/ca.crt
(I used EZ for the org name)
</pre>

### Server CA

For this example, we will run the container using the name 'service' on
a docker network created for this app. Thus the server certificate
will need to be generated with a CN of `service`

<pre>
openssl genrsa -out server/service.key 2048
openssl req -new -key server/service.key -sha256 -out server/service.csr
openssl x509 -req -days 3650 -sha256 -in server/service.csr -CA ca/ca.crt -CAkey ca/ca.key -set_serial 1 -out server/service.crt
</pre>

To build the container, first compile the app:

<pre>
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server .
</pre>

Copy the service.key, service.crt, and ca.crt into the server directory, 
then

<pre>
docker build -t server .
</pre>

### Client Certs

<pre>
mkdir client
openssl genrsa -out client/okguy.key 2048
openssl req -new -key client/okguy.key -out client/okguy.csr
(I used okguy for the CN)
openssl x509 -req -days 3650 -sha256 -in client/okguy.csr -CA ca/ca.crt -CAkey ca/ca.key -set_serial 2 -out client/okguy.crt
</pre>

### Proxy

The proxy needs to have TLS set up for its listening endpoint, and needs
TLS set up for the connection to the origin server.

To generate the certificate for the proxy endpoint for this example:

<pre>
openssl genrsa -out server/localhost.key 2048
openssl req -new -key server/localhost.key -sha256 -out server/localhost.csr
(I used localhost for the CN)
openssl x509 -req -days 3650 -sha256 -in server/localhost.csr -CA ca/ca.crt -CAkey ca/ca.key -set_serial 1 -out server/localhost.crt
</pre>

To run nginx, we need to copy the certs and keys referenced in the rp.conf
file into a directory named certs, then run the following:

<pre>
docker run -p 5000:5000 -v $PWD/rp.conf:/etc/nginx/nginx.conf -v $PWD/certs:/tmp/certs --network foo nginx
</pre>


### Client 

At this point with all the certs generated, it is a simple matter of running the client:

<pre>
go run main.go ../openssl/client/okguy.key ../openssl/client/okguy.crt ../openssl/ca/ca.crt
</pre>