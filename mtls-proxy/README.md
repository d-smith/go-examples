# MTLS Proxy

This project shows how to set up a reverse caching proxy in front of a
golang service where Mutual TLS is required to access to proxied server.

The pieces:

* The http server running in a docker container
* The proxy server running in a docker container
* The client that calls the server through the proxy

The server only allows requests from a known client.

## The Server