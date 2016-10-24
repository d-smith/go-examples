# MTLS Proxy

This project shows how to set up a reverse caching proxy in front of a
golang service where Mutual TLS is required to access to proxied server.

The pieces:

* The http server running in a docker container
* The proxy server running in a docker container
* The client that calls the server through the proxy

The server only allows requests from a known client.

## Certificate Setup

To run the example, we need a root CA and cerficiates for the client, the
proxy server, and the origin server.

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
<pre>