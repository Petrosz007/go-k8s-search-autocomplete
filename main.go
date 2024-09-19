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
	r.Run()
}
