# Go k8s Search Autocomplete

HTTP service which provides search autocompletion for Kubernetes resources running in the cluster.

## Required tools

Assuming Docker and Go is installed on your machine, you'll need these project specific packages:

```sh
brew install minikube just
```

## Running it locally

> [!WARNING]  
> Minikube will edit your Kubeconfig and add its configuration there. It should add it next to your current config, but it's never a bad idea to back it up, just in case.

To run the app in the local minikube cluster:

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

## Notes

I've added `// TODO` comments for parts which would be nice to write, but would require more effort than it's worth for this POC.

I've added `// ?` comments for parts, where I'm not too familiar with the pragmatic Go way of writing things, and I'm not sure what's the best choice.
