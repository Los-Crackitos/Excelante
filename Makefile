init:
	cp docker-compose.override.yml.dist docker-compose.override.yml
	cp entrypoint.sh.dist entrypoint.sh

coding-style:
	go mod tidy
	gofmt -l -w ./
	go vet ./...
	goimports -l -w ./
	golint -set_exit_status ./...
