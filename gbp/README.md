gbp - go build pack

Push and run a simple go app.

cf push gbp

Or, to push a binary:

* Build the binary e.g. GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
* Push via cf push my_app -c './executable-file' -b binary-buildpack
