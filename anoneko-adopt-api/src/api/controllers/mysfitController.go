package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"anoneko-adopt-api/src/api/models"
	"os"
)

type Mysfits struct {
	Mysfits []models.Mysfit `json:"mysfits"`
}

func GetMysFits(c *gin.Context) {
	jsonFile, err := os.ReadFile("src/database/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON file into the Mysfits struct
	var mysfits Mysfits
	err = json.Unmarshal(jsonFile, &mysfits)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, mysfits)
}
