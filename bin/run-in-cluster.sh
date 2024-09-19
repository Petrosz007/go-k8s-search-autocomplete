#!/usr/bin/env bash

set -euxo pipefail

# Start minikube if it isn't running
minikube status || minikube start

GOOS=linux go build -o ./target/search-autocomplete .
eval $(minikube docker-env --shell bash)
docker build -t search-autocomplete:local-run .

kubectl delete -f local-run.yaml --wait=true || true # We ignore kubectl's error with "|| true" when there is nothing to delete
kubectl apply -f local-run.yaml
kubectl wait --for=condition=ready pod search-autocomplete-local-run

kubectl port-forward --pod-running-timeout=1m0s service/search-autocomplete-local-run 8080:8080 > /dev/null & # Running the kubectl proxy in the backround, silently
kubectl logs -f search-autocomplete-local-run
