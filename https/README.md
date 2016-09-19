Cert and key gen - use built in src/crypto/tls/generate_cert.go

<pre>
cd $GOROOT
cd src/crypto/tls
go build generate_cert.go
</pre>

To run from your working directory:

<pre>
$GOROOT/src/crypto/tls/generate_cert -ca -host localhost
</pre>


Once the cert and key are in place, you can hit it via curl

<pre>
curl -k https://localhost:10443
</pre>


