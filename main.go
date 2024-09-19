package main

import (
	"net/http"

	"github.com/Petrosz007/go-k8s-search-autocomplete/internal/k8s"
	"github.com/Petrosz007/go-k8s-search-autocomplete/internal/suggestion"
	"github.com/gin-gonic/gin"
)

func main() {
	clientset := k8s.Client()

	r := gin.Default()
	r.GET("/search/autocomplete/pods", func(c *gin.Context) {
		pods := k8s.Pods(clientset)

		c.JSON(http.StatusOK, gin.H{
			"suggestions": suggestion.Suggestions(pods),
		})
	})
	r.GET("/livez", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	r.GET("/readyz", func(c *gin.Context) {
		// Do a request against the k8s api to check if we can still access it
		// It panics if there is any failure, which will be a non 2XX response, which'll fail the readiness probe
		k8s.Pods(clientset)
		c.JSON(http.StatusOK, "ok")
	})
	r.Run()
}
