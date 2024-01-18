package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"mythical-mysfit-api/initializers"
	"mythical-mysfit-api/models"
	"os"
)

type Mysfits struct {
	Mysfits []models.Mysfit `json:"mysfits"`
}

func init() {
	// Load the .env file
	initializers.LoadEnvVars()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "mythical-mysfit-api is running!",
		})
	})

	r.GET("/mysfits", func(c *gin.Context) {
		// Read the JSON file
		jsonFile, err := os.ReadFile("data/data.json")
		if err != nil {
			log.Fatal(err)
		}

		// Parse the JSON file into the Mysfits struct
		var mysfits Mysfits
		err = json.Unmarshal(jsonFile, &mysfits)
		if err != nil {
			log.Fatal(err)
		}

		// Return the parsed JSON data as a response
		c.JSON(200, mysfits)
	})

	r.Run()
}
