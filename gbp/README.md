gbp - go build pack

Push and run a simple go app.

cf push gbp


To use golang 1.9.2, use the build pack from github - see
https://github.com/cloudfoundry/go-buildpack


cf push gbp -b https://github.com/cloudfoundry/go-buildpack.git#v1.8.13

