dev:
	rm microapp
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o microapp .
build:
	rm microapp
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o microapp .
	docker build -t ct-cli .
