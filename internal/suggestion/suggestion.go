package suggestion

import (
	"github.com/Petrosz007/go-k8s-search-autocomplete/internal/k8s"
	corev1 "k8s.io/api/core/v1"
)

type Suggestion struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func label_suggestions(labels map[string][]string) []Suggestion {
	var suggestions []Suggestion
	for key, values := range labels {
		suggestions = append(suggestions, Suggestion{Key: "labels:" + key, Value: values})
	}

	return suggestions
}

func Suggestions(pods []corev1.Pod) []Suggestion {
	suggestions := []Suggestion{
		{Key: "namespace", Value: k8s.Namespaces(pods)},
		{Key: "phase", Value: k8s.Phases(pods)},
	}

	suggestions = append(suggestions, label_suggestions(k8s.Labels(pods))...)

	return suggestions
}
