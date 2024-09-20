dev:
  minikube start

build:
  go build -o ./target/search-autocomplete .

test:
  go test ./...

run:
  go run .

fmt:
  go fmt ./...

lint:
  golangci-lint run ./...

run-in-cluster:
  ./bin/run-in-cluster.sh

integration-test:
  ./bin/integration-test-in-cluster.sh
