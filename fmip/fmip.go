package fmip

import (
	"github.com/petitout/ifog/models"
	"net/http"
)

const URL_KEY = "findme"
const ENDPOINT_URL = "/fmipservice/client/web"
const REFRESH_CLIENT_URL = ENDPOINT_URL + "/refreshClient"
const SEND_MESSAGE_URL = ENDPOINT_URL + "/sendMessage"

const Method = http.MethodPost

type ClientContext struct {
	Fmly           bool   `json:"fmly"`
	ShouldLocate   bool   `json:"shouldLocate"`
	SelectedDevice string `json:"selectedDevice"`
}

type RequestCommonBody struct {
	Dsid          string        `json:"dsid"`
	ClientId      string        `json:"clientId"`
}

type RefreshClientRequestBody struct {
	RequestCommonBody
	ClientContext ClientContext `json:"clientContext"`
}

type SendMessageRequestBody struct {
	RequestCommonBody
	Device string `json:"device"`
	Subject string `json:"subject"`
	Sound bool `json:"sound"`
	UserText bool `json:"userText"`
	Text string `json:"text"`
}

type ResponseBody struct {
	Devices []models.Device `json:"content"`
}

var DefaultClientContext = ClientContext{true, true, "all"}
