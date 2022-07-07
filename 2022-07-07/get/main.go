package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	myName := os.Getenv("NAME")
	cache := map[string]gin.H{}

	fmt.Println(myName)
	r := gin.Default()
	r.GET("/hw", func(c *gin.Context) {
		clientName := c.Query("name")
		if _, ok := cache[clientName]; !ok {
			cache[clientName] = gin.H{
				"server":  myName,
				"client":  clientName,
				"message": fmt.Sprintf("hi %s, I'm %s", clientName, myName),
			}
			fmt.Println("write to cache")
		}

		c.JSON(http.StatusOK, cache[clientName])
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
