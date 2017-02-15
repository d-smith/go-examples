image:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o pfactors
	docker build -t xtracdev/pfactors .

clean:
	rm -f pfactors
