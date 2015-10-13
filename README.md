Misc go examples. The structure of this repo is in line with the 
'How to Write Go Code' guidelines.

Packages

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
* cookies - simple http cookie jar sample
* docker-list-containers - example use of the docker api to list docker containers
* decorate-http-hf - add a decorator method to functions with the http handler signature
* dynamo - AWS DynamodDB golang sample
* errors - customer error sample
* flags - basic use of the flags package
* fn-method-rev - example showing how to use a function type as a method signature.
* genkeys - generate a private/public key pair
* hello - hello world, plus an example of an interface and implementation too.
* hexdump - read from the command line and run it through the encoding/hex Dumper.
* http - Very basic http server that echos back the last uri component beyond /echo. Includes two unit test
samples using the built in testing package and using the testify assert package.
* httpgzip - Sample that shows how to use gzip compression with HTTP on both the client and server sides.
* https - Sample that shows how to configure a server to use HTTPS



