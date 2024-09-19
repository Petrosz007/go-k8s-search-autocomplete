package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Suggestion struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func main() {
	r := gin.Default()
	r.GET("/search/autocomplete/pods", func(c *gin.Context) {
		suggestions := []Suggestion{
			{Key: "namespace", Value: []string{"default", "prod", "staging"}},
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
