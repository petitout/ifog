package models

import (
	"testing"
)

func TestLocation_String(t *testing.T) {
	type fields struct {
		IsOld              bool
		IsInaccurate       bool
		Altitude           float64
		PositionType       string
		Latitude           float64
		Longitude          float64
		HorizontalAccuracy float64
		VerticalAccuracy   float64
		Timestamp          uint64
		LocationFinished   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Nominal", fields: fields{Latitude: 1.1, Longitude: 2.2}, want: "1.1,2.2"},
	}
	for _, tt := range tests {
		this := &Location{
			IsOld:              tt.fields.IsOld,
			IsInaccurate:       tt.fields.IsInaccurate,
			Altitude:           tt.fields.Altitude,
			PositionType:       tt.fields.PositionType,
			Latitude:           tt.fields.Latitude,
			Longitude:          tt.fields.Longitude,
			HorizontalAccuracy: tt.fields.HorizontalAccuracy,
			VerticalAccuracy:   tt.fields.VerticalAccuracy,
			Timestamp:          tt.fields.Timestamp,
			LocationFinished:   tt.fields.LocationFinished,
		}
		if got := this.String(); got != tt.want {
			t.Errorf("Location.String() = %v, want %v", got, tt.want)
		}
	}
}

func TestDevice_IsIPhone(t *testing.T) {
	type fields struct {
		Location    Location
		DeviceClass string
		Id          string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Actually an iphone", fields: fields{DeviceClass: "iPhone"}, want: true},
		{name: "Actually an imac", fields: fields{DeviceClass: "iMac"}, want: false},
		{name: "Unknown device class", fields: fields{DeviceClass: "blahblah"}, want: false},
	}
	for _, tt := range tests {
		d := &Device{
			Location:    tt.fields.Location,
			DeviceClass: tt.fields.DeviceClass,
			Id:          tt.fields.Id,
		}
		if got := d.IsIPhone(); got != tt.want {
			t.Errorf("Device.IsIPhone() = %v, want %v", got, tt.want)
		}
	}
}

func TestDevice_IsImac(t *testing.T) {
	type fields struct {
		Location    Location
		DeviceClass string
		Id          string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Actually an iphone", fields: fields{DeviceClass: "iPhone"}, want: false},
		{name: "Actually an imac", fields: fields{DeviceClass: "iMac"}, want: true},
		{name: "Unknown device class", fields: fields{DeviceClass: "blahblah"}, want: false},
	}
	for _, tt := range tests {
		d := &Device{
			Location:    tt.fields.Location,
			DeviceClass: tt.fields.DeviceClass,
			Id:          tt.fields.Id,
		}
		if got := d.IsIMac(); got != tt.want {
			t.Errorf("Device.IsIMac() = %v, want %v", got, tt.want)
		}
	}
}
