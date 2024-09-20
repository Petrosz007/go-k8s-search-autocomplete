# Go k8s Search Autocomplete

HTTP service which provides search autocompletion for Kubernetes resources running in the cluster.

## Required tools

Assuming Docker and Go is installed on your machine, you'll need these project specific packages:

```sh
brew install minikube just bruno golangci-lint

# Optional: to use bruno on the cli install bruno-cli with npm, the intergation test script uses this
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

- Write more tests for the Go code
  - Mock out the k8s API to test the my k8s module
  - Mock out my k8s module to test the API layer
- Config loading
  - Right now only the port could be a config
  - k8s client could be configured with other options, timeouts could be set, ... 
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
