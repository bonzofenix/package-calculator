docker-test:
		docker run -it --rm -v $(PWD):/app -w /app -e GO111MODULE=on golang:1.21 make test
test:
		go test -v ./...

watch:
		ginkgo watch ./...

run:
		go run cmd/main.go

docker-build:
		docker build -t bonzofenix/package-calculator .

docker-run:
		docker run -p 8080:8080 bonzofenix/package-calculator

docker-push:
		docker tag bonzofenix/package-calculator bonzofenix/package-calculator:latest
		docker push bonzofenix/package-calculator:latest

deploy:
	gcloud --quiet app deploy
