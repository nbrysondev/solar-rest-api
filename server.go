// build
// env GOOS=linux GOARCH=arm GOARM=5 go build
package main

import (
	"log"
	"encoding/json"
	"os/exec"
	"github.com/gin-gonic/gin"
)

type SolarData struct {
	RealTime string `json:"realtime"`
	Today string `json:"today"`
	Yesterday string `json:"yesterday"`
	Month string `json:"month"`
	Year string `json:"year"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		out, err := exec.Command("/bin/sh", "/home/pi/get-solar-data.sh").Output()
		if err != nil {
			log.Println(err)
		} else {
			data := SolarData{}
			json.Unmarshal(out, &data)
			c.JSON(200,data)
		}
	})

	/*
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List of users",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Created a new user",
		})
	})
*/
	r.Run() // listen and serve on 0.0.0.0:8080
}