package login

import (
	"github.com/petitout/ifog/models"
	"net/http"
)

const URL = "https://setup.icloud.com/setup/ws/1/login"

const Method = http.MethodPost

type RequestBody struct {
	AppleId  string `json:"apple_id"`
	Password string `json:"password"`
}

type ResponseBody struct {
	User        models.User                  `json:"dsInfo"`
	Webservices map[string]models.WebService `json:"webservices"`
}
