package model

import (
	"time"
)

type Footprint struct {
	ID          uint
	CurrentTime time.Time `gorm:"column:current_time;not null" json:"current_time"`
	IPAddress   string    `gorm:"column:ip_address; type:text" json:"ip_address"`
	DeviceInfo  string    `gorm:"column:device_info; type:text" json:"device_info"`
	BrowserType string    `gorm:"column:browser_type; type:text" json:"browser_type"`
	Longitude   string      `gorm:"column:longitude;type:uint" json:"longitude"`
	Latitude    string      `gorm:"column:latitude;type:uint" json:"latitude"`
	City        string    `gorm:"column:city;type:text" json:"city"`
	Country     string    `gorm:"column:country;type:text" json:"country"`
}
