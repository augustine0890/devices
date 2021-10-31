package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

// Define the data model
type Device struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
	Status      string    `json:"status"`
	Timezone    string    `json:"timezone"`
}

type Results map[string]Device

var devices Results

var ctx context.Context
var err error

// Initialization code
func init() {
	// devices = make([]Results, 0)
	// Read JSON file and convert into an array of devices
	file, err := ioutil.ReadFile("devices.json")
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal([]byte(file), &devices)

}

func ListDevicesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, devices)
}

func GetDeviceByIDHandler(c *gin.Context) {
	id := c.Param("id")

	for _, item := range devices {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Device not found",
	})
}

func GetDeviceByTypeHandler(c *gin.Context) {
	typeOf := c.Query("type")
	listOfRecipes := make([]Device, 0)

	for _, item := range devices {
		found := false
		for _, t := range item.Type {
			if strings.EqualFold(t, typeOf) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, item)
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}

func GetDeviceByStatusHandler(c *gin.Context) {
	status := c.Query("status")
	listOfRecipes := make([]Device, 0)

	for _, item := range devices {
		found := false
		for _, s := range item.Status {
			if strings.EqualFold(s, status) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, item)
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}

func main() {
	router := gin.Default()

	// Set the memory limit
	router.Use(limits.RequestSizeLimiter(200))

	// GET list devices
	router.GET("/devices", ListDevicesHandler)
	// GET device by id
	router.GET("/devices/:id", GetDeviceByIDHandler)
	// GET list of device by type
	router.GET("/devices/type", GetDeviceByTypeHandler)
	// GET list of device by status
	router.GET("/devices/status", GetDeviceByStatusHandler)

	// The default listen port is 8080
	router.Run()
}
