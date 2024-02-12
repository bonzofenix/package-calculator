
docker-test:
		docker run -it --rm -v $(PWD):/app -w /app -e GO111MODULE=on golang:1.21 make test
test:
		go test -v ./...
