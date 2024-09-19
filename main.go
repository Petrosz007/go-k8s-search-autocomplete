package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Suggestion struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

// TODO: Move this to another file
func Namespaces(clientset *kubernetes.Clientset) []string {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var namespace_names []string
	for _, namespace := range namespaces.Items {
		namespace_names = append(namespace_names, namespace.Name)
	}

	return namespace_names
}

func main() {
	// TODO: Move this setup to another file
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.GET("/search/autocomplete/pods", func(c *gin.Context) {
		suggestions := []Suggestion{
			{Key: "namespace", Value: Namespaces(clientset)},
			{Key: "phase", Value: []string{"Pending", "Running", "Failed"}},
			{Key: "labels:app.kubernetes.io/name", Value: []string{"nginx", "mysql", "elasticsearch"}},
			{Key: "labels:app.kubernetes.io/component", Value: []string{"webserver", "database"}},
			{Key: "labels:team", Value: []string{"backend", "frontend"}},
		}

		c.JSON(http.StatusOK, gin.H{
			"suggestions": suggestions,
		})
	})
	r.Run()
}
