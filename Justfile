dev:
  minikube start

run:
  go run .

fmt:
  go fmt ./...

run-in-cluster:
  ./bin/run-in-cluster.sh
