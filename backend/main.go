package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api/introduction", func(c *gin.Context) {
		serveJSONFile(c, "introduction.json")
	})

	r.GET("/api/projects", func(c *gin.Context) {
		serveJSONFile(c, "projects.json")
	})

	r.GET("/api/research", func(c *gin.Context) {
		serveJSONFile(c, "research.json")
	})

	r.Run("0.0.0.0:8080")
}

func serveJSONFile(c *gin.Context, fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
		return
	}

	var parsedData interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		fmt.Println("Failed to parse the data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}

	c.JSON(http.StatusOK, parsedData)
}
