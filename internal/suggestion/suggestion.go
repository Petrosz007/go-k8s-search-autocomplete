package suggestion

import (
	"github.com/Petrosz007/go-k8s-search-autocomplete/internal/utils"
	corev1 "k8s.io/api/core/v1"
)

type Suggestion struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func extract_values[T any](xs []T, lens func(T) string) []string {
	var values []string
	for _, x := range xs {
		values = append(values, lens(x))
	}

	return utils.Uniques(values)
}

func extract_maps[T any](xs []T, lens func(T) map[string]string) map[string][]string {
	extracted := make(map[string][]string)
	for _, x := range xs {
		for key, value := range lens(x) {
			extracted[key] = append(extracted[key], value)
		}
	}

	for key, values := range extracted {
		extracted[key] = utils.Uniques(values)
	}

	return extracted
}

// TODO: Add function extractor for values inside array, for example container names
// It'd require two lens functions: One lens for the array, and then another lens for looking inside the array element

func prefix_map_suggestions(prefix string, maps map[string][]string) []Suggestion {
	var suggestions []Suggestion
	for key, values := range maps {
		suggestions = append(suggestions, Suggestion{Key: prefix + ":" + key, Value: values})
	}

	return suggestions
}

func Pods(pods []corev1.Pod) []Suggestion {
	suggestions := []Suggestion{
		{Key: "namespace", Value: extract_values(pods, func(pod corev1.Pod) string { return pod.Namespace })},
		{Key: "phase", Value: extract_values(pods, func(pod corev1.Pod) string { return string(pod.Status.Phase) })},
		// Example: Add Name as well
		// {Key: "name", Value: extract_values(pods, func(pod corev1.Pod) string { return string(pod.Name) })},
	}

	suggestions = append(suggestions, prefix_map_suggestions("labels", extract_maps(pods, func(pod corev1.Pod) map[string]string { return pod.Labels }))...)
	// Example: to add Annotations as well:
	// suggestions = append(suggestions, prefix_map_suggestions("annotations", extract_maps(pods, func(pod corev1.Pod) map[string]string { return pod.Annotations }))...)

	return suggestions
}

// Example: Add suggestions for Deployments
// This would need a k8s client client call in k8s.go to list the deployments

// func Deployments(deployments []appsv1.Deployment) []Suggestion {
// 	suggestions := []Suggestion{
// 		{Key: "namespace", Value: extract_values(deployments, func(deployment appsv1.Deployment) string { return deployment.Namespace })},
// 		{Key: "name", Value: extract_values(deployments, func(deployment appsv1.Deployment) string { return deployment.Name })},
// 	}

// 	suggestions = append(suggestions, prefix_map_suggestions("labels", extract_maps(deployments, func(deployment appsv1.Deployment) map[string]string { return deployment.Labels }))...)

// 	return suggestions
// }
