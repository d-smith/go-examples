Misc go examples. The structure of this repo is in line with the 
'How to Write Go Code' guidelines.

Packages

* actors - Experiment with the actor model using golang (incomplete)
* aq - Experments calling Oracle AQ via the golang sql driver
* awsreg - Function to return the AWS region from the environment, or a
default value if nothing specified in the environments
* base64 - simple base64 encode and decode example
* bowling - the bowling kata
* certs - example of decoding a PEM encoded cert, parsing the cert, and extracting the public key
* cfun - simple example of how to use cgo to call a C function.
* clang - Another C language sample
* cli - Example use of the mitchelh cli library
* client - Simple http client that can do a get on a URL. Can also use to show how the http_proxy is picked
up automatically by go.
* compile-string-init - Initialize a string variable via a compile time flag
* concurrency - working example code from Rob Pike's [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1) talk
* concurrent - Two concurrency examples, one using shared memory and synchronization, one using message passing.
* consul-kv-read - Simple consul api sample for reading a value by key
* consul-service-reg - Read a service definition from consul
* cookies - simple http cookie jar sample
* ctx-hierarchy - example of wiring in context aware http handlers
* custom-handler - package for context aware handlers and samples of how to create and use context aware handlers
* decorate-http-hf - decorator pattern on http handlers
* docker-list-containers - example use of the docker api to list docker containers
* docker-pkg - simple minded example of packaging a golang program as a Docker image
* dynamo - AWS DynamodDB golang sample
* errors - customer error sample
* es - simple event sourcing prototype
* es2 - more complete event sourcing example
* flags - basic use of the flags package
* fn-method-rev - example showing how to use a function type as a method signature.
* genkeys - generate a private/public key pair
* godash - start of some lodash-like functional programming support
* hello - hello world, plus an example of an interface and implementation too.
* hexdump - read from the command line and run it through the encoding/hex Dumper.
* http - Very basic http server that echos back the last uri component beyond /echo. Includes two unit test
samples using the built in testing package and using the testify assert package.
* http-panic - Sample panic handler for http
* httpgzip - Sample that shows how to use gzip compression with HTTP on both the client and server sides.
* https - Sample that shows how to configure a server to use HTTPS
* id-and-secret - Sample to show how to generate a application id and secret
* ignite - Simple REST get and put to Apache Ignite/Grid Gain cache




