PROJECT_ROOT=$(dirname "$0")/..
KUBECTL_PROXY_PID=-1 # Global variable, so we can terminate it after other scripts are run

function start_in_cluster {
  # Start minikube if it isn't running
  minikube status || minikube start

  GOOS=linux go build -o "$PROJECT_ROOT/target/search-autocomplete" "$PROJECT_ROOT"
  eval $(minikube docker-env --shell bash)
  docker build -t search-autocomplete:local-run "$PROJECT_ROOT"

  kubectl delete -f "$PROJECT_ROOT/local-run.yaml" --wait=true || true # We ignore kubectl's error with "|| true" when there is nothing to delete
  kubectl apply -f "$PROJECT_ROOT/local-run.yaml"
  kubectl wait --for=condition=ready pod search-autocomplete-local-run

  kubectl port-forward --pod-running-timeout=1m0s service/search-autocomplete-local-run 8080:8080 > /dev/null & # Running the kubectl proxy in the backround, silently
  KUBECTL_PROXY_PID=$!
}
