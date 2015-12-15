Build a simple docker image that has a single executable for the purpose of leveraging it in the code ship
docker pipeline when building and running container steps based on the scratch image.

Build for docker via `GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o synopsis`