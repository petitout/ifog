package fmip

import (
	"net/http"

	"github.com/petitout/ifog/models"
)

// URLKey key in the ICloud map allowing to retrieve the value (URL) for the find my iphone end point
const URLKey = "findme"

// EndpointURL endpoint url
const EndpointURL = "/fmipservice/client/web"

// RefreshClientURL URL suffix to access to the refreshClient endpoint
const RefreshClientURL = EndpointURL + "/refreshClient"

// SendMessageURL URL suffix allowing to send a message
const SendMessageURL = EndpointURL + "/sendMessage"

// Method HTTP method to use for the find my iphone endpoint
const Method = http.MethodPost

// ClientContext context to pass to the fmip endpoint
type ClientContext struct {
	Fmly           bool   `json:"fmly"`
	ShouldLocate   bool   `json:"shouldLocate"`
	SelectedDevice string `json:"selectedDevice"`
}

// RequestCommonBody json structure to use in the request
type RequestCommonBody struct {
	Dsid     string `json:"dsid"`
	ClientId string `json:"clientId"`
}

// RefreshClientRequestBody json struct to use in the refresh client request
type RefreshClientRequestBody struct {
	RequestCommonBody
	ClientContext ClientContext `json:"clientContext"`
}

// SendMessageRequestBody json struct to use in the send message request
type SendMessageRequestBody struct {
	RequestCommonBody
	Device   string `json:"device"`
	Subject  string `json:"subject"`
	Sound    bool   `json:"sound"`
	UserText bool   `json:"userText"`
	Text     string `json:"text"`
}

// ResponseBody json struct received in the response
type ResponseBody struct {
	Devices []models.Device `json:"content"`
}

// DefaultClientContext default client context
var DefaultClientContext = ClientContext{true, true, "all"}
