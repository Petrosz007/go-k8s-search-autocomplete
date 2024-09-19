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

run-in-cluster:
  ./bin/run-in-cluster.sh
