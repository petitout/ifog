package fmip

import (
	"github.com/petitout/ifog/models"
	"net/http"
)

const URL_KEY = "findme"

const Method = http.MethodPost

type ClientContext struct {
	Fmly           bool   `json:"fmly"`
	ShouldLocate   bool   `json:"shouldLocate"`
	SelectedDevice string `json:"selectedDevice"`
}

type RequestBody struct {
	Dsid          string        `json:"dsid"`
	ClientId      string        `json:"clientId"`
	ClientContext ClientContext `json:"clientContext"`
}

type ResponseBody struct {
	Devices []models.Device `json:"content"`
}

var DefaultClientContext = ClientContext{true, true, "all"}
