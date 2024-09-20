package suggestion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPods(t *testing.T) {
	// Arrange
	pods := []corev1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "production",
				Labels: map[string]string{
					"label1": "key1-1",
					"label2": "key2-1",
				}},
			Status: corev1.PodStatus{Phase: corev1.PodFailed},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "staging",
				Labels: map[string]string{
					"label1": "key1-2",
					"label3": "key3",
				}},
			Status: corev1.PodStatus{Phase: corev1.PodUnknown},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "production",
				Labels:    map[string]string{}},
			Status: corev1.PodStatus{Phase: corev1.PodFailed},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Labels: map[string]string{
					"label4": "key4",
					"label2": "key2-2",
				}},
			Status: corev1.PodStatus{Phase: corev1.PodPhase(corev1.PodReady)},
		},
	}
	expected := []Suggestion{
		{Key: "namespace", Value: []string{"production", "staging", "default"}},
		{Key: "phase", Value: []string{"Failed", "Unknown", "Ready"}},
		{Key: "labels:label1", Value: []string{"key1-1", "key1-2"}},
		{Key: "labels:label2", Value: []string{"key2-1", "key2-2"}},
		{Key: "labels:label3", Value: []string{"key3"}},
		{Key: "labels:label4", Value: []string{"key4"}},
	}

	// Act
	result := Pods(pods)

	// Assert
	// TODO: This test can be flaky, because the value keys can also have a different order, which makes the assertion fail
	// ?: I'm not sure what's the best Go way to write an "ElementsMatch, but deeply, so if each element contains a list somewhere make sure the order doesn't matter" function
	// Maybe thsi package could be used: https://github.com/r3labs/diff
	assert.ElementsMatch(t, expected, result)
}
