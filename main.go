// Devices API
//
// This is a sample devices API.
//
//	Schemes: http
//  Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
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

type Result map[string]Device

var result Result
var devices []Device

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
	_ = json.Unmarshal([]byte(file), &result)

	devices = make([]Device, 0)
	for _, item := range result {
		devices = append(devices, item)
	}

	log.Println("Fetched devices: ", len(devices))
}

// swagger:operation GET /devices devices listDevices
// Returns list of devices
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
func ListDevicesHandler(c *gin.Context) {
	pagination := PaginationRequest(c)
	start, end := (pagination.Page-1)*pagination.Limit, pagination.Page*pagination.Limit

	c.JSON(http.StatusOK, devices[start:end])
}

// swagger:operation POST /devices devices newDevice
// Create a new device
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
func NewDeviceHandler(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	device.ID = xid.New().String()
	devices = append(devices, device)

	c.JSON(http.StatusOK, device)
}

// swagger:operation GET /devices/{id} devices oneDevice
// Get one device
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID of the devices
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Invalid devices ID
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

// swagger:operation PUT /devices/{id} devices updateDevice
// Update an existing device
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the devices
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
//     '404':
//         description: Invalid devices ID
func UpdateDeviceHandler(c *gin.Context) {
	id := c.Param("id")
	var updateDevice Device
	if err := c.ShouldBindJSON(&updateDevice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for idx, item := range devices {
		if item.ID == id {
			updateDevice.ID = id
			devices[idx] = updateDevice
			break
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Device not found",
			})
			return
		}
	}

	c.JSON(http.StatusOK, updateDevice)
}

// swagger:operation GET /devices/type devices findDeviceByType
// Search devices based on type
// ---
// produces:
// - application/json
// parameters:
//   - name: type
//     in: query
//     description: device type
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
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

// swagger:operation GET /devices/status devices findDeviceByStatus
// Search devices based on status
// ---
// produces:
// - application/json
// parameters:
//   - name: status
//     in: query
//     description: device status
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
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

	// Create new devices
	router.POST("/devices", NewDeviceHandler)
	// GET list devices
	router.GET("/devices", ListDevicesHandler)
	// GET device by id
	router.GET("/devices/:id", GetDeviceByIDHandler)
	// Update the device
	router.PUT("/devices/:id", UpdateDeviceHandler)
	// GET list of device by type
	router.GET("/devices/type", GetDeviceByTypeHandler)
	// GET list of device by status
	router.GET("/devices/status", GetDeviceByStatusHandler)

	// The default listen port is 8080
	router.Run()
}
