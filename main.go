package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

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

func GetDeviceHandler(c *gin.Context) {
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

func main() {
	router := gin.Default()

	// Set the memory limit
	router.Use(limits.RequestSizeLimiter(200))

	// GET list devices
	router.GET("/devices", ListDevicesHandler)
	// GET device by id
	router.GET("/devices/:id", GetDeviceHandler)
	// The default listen port is 8080
	router.Run()
}
