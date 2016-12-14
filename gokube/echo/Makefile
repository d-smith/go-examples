image:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o echo .
	docker build -t dsmith/echo .

clean:
	rm -f echo
