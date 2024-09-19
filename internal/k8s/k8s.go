package k8s

import (
	"context"

	"github.com/Petrosz007/go-k8s-search-autocomplete/internal/utils"
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
		panic(err.Error())
	}

	var pods []corev1.Pod
	for _, pod := range pod_list.Items {
		pods = append(pods, pod)
	}

	return pods
}

func Namespaces(pods []corev1.Pod) []string {
	var namespaces []string
	for _, pod := range pods {
		namespaces = append(namespaces, pod.Namespace)
	}

	return utils.Uniques(namespaces)
}

func Phases(pods []corev1.Pod) []string {
	var phases []string
	for _, pod := range pods {
		phases = append(phases, string(pod.Status.Phase))
	}

	return utils.Uniques(phases)
}

func Labels(pods []corev1.Pod) map[string][]string {
	labels := make(map[string][]string)
	for _, pod := range pods {
		for label_key, label_value := range pod.Labels {
			labels[label_key] = append(labels[label_key], label_value)
		}
	}

	for key, values := range labels {
		labels[key] = utils.Uniques(values)
	}

	return labels
}
