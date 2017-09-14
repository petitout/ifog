package models

import "strconv"

type WebService struct {
	Url         string `json:"url"`
	Status      string `json:"status"`
	PcsRequired bool   `json:"pcsRequired"`
}

type User struct {
	FirstName                 string `json:"firtName"`
	FullName                  string `json:"fullName"`
	Locked                    bool   `json:"locked"`
	StatusCode                uint32 `json:"statusCode"`
	PrimaryEmail              string `json:"primaryEmail"`
	Dsid                      string `json:"dsid"`
	AppleId                   string `json:"appleId"`
	IsPaidDeveloper           bool   `json:"isPaidDeveloper"`
	HasICloudQualifyingDevice bool   `json:"hasICloudQualifyingDevice"`
	Locale                    string `json:"locale"`
	AppleIdAlias              string `json:"appleIdAlias"`
	LastName                  string `json:"lastName"`
	ICloudAppleIdAlias        string `json:"iCloudAppleIdAlias"`
	PrimaryEmailVerified      bool   `json:"primaryEmailVerified"`
}

type Location struct {
	IsOld              bool    `json:"isOld"`
	IsInaccurate       bool    `json:"isInaccurate"`
	Altitude           float64 `json:"altitude"`
	PositionType       string  `json:"positionType"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	HorizontalAccuracy float64 `json:"horizontalAccuracy"`
	VerticalAccuracy   float64 `json:"verticalAccuracy"`
	Timestamp          uint64  `json:"timestamp"`
	LocationFinished   bool    `json:"locationFinished"`
}

func (this *Location) String() string {
	return strconv.FormatFloat(this.Latitude, 'f', -1, 64) + "," + strconv.FormatFloat(this.Longitude, 'f', -1, 64)
}

type Device struct {
	Location    Location `json:"location"`
	DeviceClass string   `json:"deviceClass"`
	Id          string   `json:"Id"`
}

func (this *Device) IsIPhone() bool {
	return this.DeviceClass == "iPhone"
}

func (this *Device) IsIMac() bool {
	return this.DeviceClass == "iMac"
}
