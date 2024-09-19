# Go k8s Search Autocomplete

HTTP service which provides search autocompletion for Kubernetes resources running in the cluster.

## Required tools

Assuming Docker and Go is installed on your machine, you'll need these project specific packages:

```sh
brew install minikube just bruno

# Optional: to use bruno on the cli install bruno-cli with npm
npm install -g @usebruno/cli
```

## Running it locally

> [!WARNING]  
> Minikube will edit your Kubeconfig and add its configuration there. It should add it next to your current config, but it's never a bad idea to back it up, just in case.

To run the app in the local minikube cluster:

```sh
just run-in-cluster
```

This will:
- Start minikube
- Build the app
- Build a Docker image inside minikube's docker env
- (Re-)deploy the necessary k8s resources to minikube
- Start port forwarding on 8080 and start showing the logs

You can now query the app, for example with:

```sh
curl localhost:8080/search/autocomplete/pods
```

## Integration testing

Integration tests are set up with [Bruno](https://docs.usebruno.com/introduction/what-is-bruno) and bruno-cli.

To run them automatically:

```sh
just integration-test
```

## Notes

I've added `// TODO` comments for parts which would be nice to write, but would require more effort than it's worth for this POC.

I've added `// ?` comments for parts, where I'm not too familiar with the pragmatic Go way of writing things, and I'm not sure what's the best choice.

## Future improvement ideas

### Design alternatives
The current code with the go k8s client is extendible, if new resource types or fields need to be surfaced on the API only a few new functions have to be written.

If we expect that a lot of new resource types would have to be added, then the k8s Go client can hold us back. It is type safe, but because of the type safety we need to explicitly query each resource (ex.: pod.Status.Phase). An alternative would be, to use the JSON API of k8s and provide the autocompletable fields as JSON path selectors in a config file. That way, we only define how to access specific fields in a JSON structure, and that can be scaled more easily with overall less code.

I didn't choose this approach, because for the POC it would have required much more time and lines of code, and in this phase evaluating the functionality and iterating faster is more important, than an all-extensible perfect solution for the first POC.

### TODOs
- Write more tests for the Go code
  - Mock out the k8s API to test the my k8s module
  - Mock out my k8s module to test the API layer 
- Better integration testing with Bruno
  - The tests should better test the values returned by the API. These current tests are enough to check if the API is alive and responds with some correctly formatted data.
- Prometheus metrics
- Tracing with OpenTelemetry
  - It's very cool, but probably far down the priority list
- Swagger UI
- Automated CI/CD pipeline
  - Building and packaging is easy
  - Minikube setup might be a bit tricky
  - The `bin/integration-test-in-cluster.sh` script runs the tests automatically already, we just need to call it in the pipeline
