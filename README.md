# Go k8s Search Autocomplete

## Required tools

```sh
brew install minikube just
```

## Local development

First, start minikube with:

```sh
just dev
```

Then, run the app in the local minikube cluster:

```sh
just run-in-cluster
```

This will:
- Build the app
- Build a Docker image inside minikube's docker env
- (Re-)deploy the necessary k8s resources to minikube
- Start port forwarding on 8080 and start showing the logs

You can now query the app, for example with:

```sh
curl localhost:8080/search/autocomplete/pods
```
