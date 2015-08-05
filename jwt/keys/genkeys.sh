#!/bin/sh

openssl genrsa -out app.rsa 1024
openssl rsa -in app.rsa -pubout > app.rsa.pub
openssl genrsa -out otherkey.rsa 1024
openssl rsa -in otherkey.rsa -pubout > otherkey.rsa.pub
