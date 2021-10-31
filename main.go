package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func main() {
	router := gin.Default()

	router.GET("/devices", ListDevicesHandler)

	router.Run()
}
