This is a simple example of packaging an executable produced with
golang into a very small container.

To use:

1. Build the executable via GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o main .
2. Package it using docker build .
3. Run via docker run


