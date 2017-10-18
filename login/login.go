package login

import (
	"net/http"

	"github.com/petitout/ifog/models"
)

// URL ICloud login endpoint URL
const URL = "https://setup.icloud.com/setup/ws/1/login"

// Method HTTP method to use to trigger the login endpoint
const Method = http.MethodPost

// RequestBody structure to use in the login request
type RequestBody struct {
	AppleId  string `json:"apple_id"`
	Password string `json:"password"`
}

// ResponseBody structure received in the login response
type ResponseBody struct {
	User        models.User                  `json:"dsInfo"`
	Webservices map[string]models.WebService `json:"webservices"`
}
