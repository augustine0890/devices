package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
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

func PaginationRequest(c *gin.Context) Pagination {
	limit := 10
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return Pagination{
		Limit: limit,
		Page:  page,
	}
}

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
	listOfDevices := make([]Device, 0)

	for _, item := range devices {
		listOfDevices = append(listOfDevices, item)
	}
	pagination := PaginationRequest(c)
	start, end := (pagination.Page-1)*pagination.Limit, pagination.Page*pagination.Limit

	c.JSON(http.StatusOK, listOfDevices[start:end])
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
	listOfDevices := make([]Device, 0)

	for _, item := range devices {
		found := false
		if strings.EqualFold(item.Type, typeOf) {
			found = true
		}
		if found {
			listOfDevices = append(listOfDevices, item)
		}
	}
	pagination := PaginationRequest(c)
	start, end := (pagination.Page-1)*pagination.Limit, pagination.Page*pagination.Limit

	c.JSON(http.StatusOK, listOfDevices[start:end])
}

func GetDeviceByStatusHandler(c *gin.Context) {
	status := c.Query("status")
	listOfDevices := make([]Device, 0)

	for _, item := range devices {
		found := false
		if strings.EqualFold(item.Status, status) {
			found = true
		}
		if found {
			listOfDevices = append(listOfDevices, item)
		}
	}
	pagination := PaginationRequest(c)
	start, end := (pagination.Page-1)*pagination.Limit, pagination.Page*pagination.Limit

	c.JSON(http.StatusOK, listOfDevices[start:end])

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
