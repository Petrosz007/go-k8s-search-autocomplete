package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Client() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

// ?: My instinct to pass the k8s client as a parameter to this function,
// but we could use a global var for the client and that way we can easily use it in this file, no need to call the Client() function outside this file.
// Not sure what's the best practice for this in Go.
func Pods(clientset *kubernetes.Clientset) []corev1.Pod {
	pod_list, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// TODO: In this POC panicing is fine, but proper error handling could be added
		// If we can't list pods, we'll try to backtrack, check if it's a permission issue, or if the connection to k8s fail
		panic(err.Error())
	}

	var pods []corev1.Pod
	pods = append(pods, pod_list.Items...)

	return pods
}
