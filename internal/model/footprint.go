package model

import (	
	"time"
)

type Footprint struct {
	CurrentTime     time.Time 	
    IPAddress		string 		
    DeviceInfo      string		
	BrowserType     string		
	Location       struct {
		Longitude string `json:"longitude"`
		Latitude string `json:"latitude"`
		City    string `json:"city"`
		Country string `json:"country"`
	}
}
