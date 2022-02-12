package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

/**************ENDPOINTS******************/

// ESP32 MCU Interacts with These Values
type pot struct {
	PotID    string  `json:"id"`
	PotColor []int8  `json:"potcolor"`
	Temp     float64 `json:"temp"`
	Water    float64 `json:"water"`
	Light    float64 `json:"light"`
}

// iOS Device Interacts with These Values
type user struct {
	// Device	string	`json:"device"`
	UserID    string `json:"id"`
	UserColor []int8 `json:"usercolor"`
}

/***************METHODS******************/

// Sensor Data Methods

// getTemp responds with the float64 temp sens value
func getTemp(c *gin.Context) {
	// TODO: How to retrieve specific value from JSON?
}
