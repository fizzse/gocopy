package tdengine

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"
)

type DeviceData struct {
	Mac             string    `gorm:"column:mac" json:"mac"`
	DeviceModel     string    `gorm:"column:device_model" json:"device_model,omitempty"`
	EventType       string    `gorm:"column:event_type" json:"event_type,omitempty"`
	Battery         *int64    `gorm:"column:battery" json:"battery,omitempty"`
	Humidity        *float64  `gorm:"column:humidity" json:"humidity,omitempty"`
	Temperature     *float64  `gorm:"column:temperature" json:"temperature,omitempty"`
	Pressure        *float64  `gorm:"column:pressure" json:"pressure,omitempty"`
	Signal          *float64  `gorm:"column:signal" json:"signal,omitempty"`
	Lumen           *float64  `gorm:"column:lumen" json:"lumen,omitempty"`
	Tvoc            *int64    `gorm:"column:tvoc" json:"tvoc,omitempty"`
	Co2             *int64    `gorm:"column:co2" json:"co2,omitempty"`
	Pm10            *float64  `gorm:"column:pm10" json:"pm10,omitempty"`
	Pm25            *float64  `gorm:"column:pm25" json:"pm25,omitempty"`
	Pm100           *float64  `gorm:"column:pm100" json:"pm100,omitempty"`
	Pomodoro        *int64    `gorm:"column:pomodoro" json:"pomodoro,omitempty"`
	ProbTemperature *float64  `gorm:"column:prob_temperature" json:"prob_temperature,omitempty"`
	ProbHumidity    *float64  `gorm:"column:prob_humidity" json:"prob_humidity,omitempty"`
	ReportTime      time.Time `gorm:"column:report_time" json:"report_time"`
}

func TestQuery(t *testing.T) {
	db, err := gorm.Open(Open("root:taosdata@http(127.0.0.1:6041)/qingping"))
	if err != nil {
		fmt.Println("new conn failed: ", err)
		return
	}

	data := &DeviceData{}
	err = db.Table("snow_data").Where("mac = ?", "582D3400ADCC").Find(&data).Error
	if err != nil {
		fmt.Println("query failed failed: ", err)
		return
	}
}
