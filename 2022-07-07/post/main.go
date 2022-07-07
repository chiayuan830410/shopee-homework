package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Req struct {
	Name string `json:"name"`
}

func main() {
	godotenv.Load()

	myName := os.Getenv("NAME")
	cache := map[string]gin.H{}

	fmt.Println(myName)
	r := gin.Default()
	r.POST("/hw", func(c *gin.Context) {
		var req Req
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "please request name",
			})
			return
		}
		if _, ok := cache[req.Name]; !ok {
			cache[req.Name] = gin.H{
				"server":  myName,
				"client":  req.Name,
				"message": fmt.Sprintf("hi %s, I'm %s", req.Name, myName),
			}
			fmt.Println("write to cache")
		}

		c.JSON(http.StatusOK, cache[req.Name])
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
