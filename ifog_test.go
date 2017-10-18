package ifog

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/petitout/ifog/fmip"

	"github.com/petitout/ifog/login"

	"github.com/petitout/ifog/models"
)

func TestSession_Login(t *testing.T) {
	s := NewSession()
	response := new(login.ResponseBody)
	response.User.AppleId = "myAppleId"
	response.Webservices = make(map[string]models.WebService)
	response.Webservices[fmip.URLKey] = models.WebService{Url: "http://falseFmipURL.com"}
	var requestHeader http.Header
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHeader = r.Header
		w.Header().Set("Content-Type", "application/json")
		responseToSend, _ := json.Marshal(*response)
		w.Write(responseToSend)
	}))
	s.LoginURL = ts.URL
	defer ts.Close()
	err := s.Login(login.RequestBody{AppleId: "myAppleId", Password: "mySuperPassword"})
	if err != nil {
		t.Errorf("Session.Login() error = %v, wantErr %v", err, nil)
	}
	if !reflect.DeepEqual(response.User, s.User) {
		t.Errorf("Session.Login() error = %v, wantErr %v", s.User, response.User)
	}
	if !reflect.DeepEqual(response.Webservices, s.Webservices) {
		t.Errorf("Session.Login() error = %v, wantErr %v", s.User, response.User)
	}
	if !checkRequestHeader(requestHeader) {
		t.Errorf("Session.Login() error = Bad request header")
	}
}

func TestSession_PopulateDevices(t *testing.T) {
	response := new(fmip.ResponseBody)
	response.Devices = []models.Device{
		models.Device{Location: models.Location{}, DeviceClass: "iPhone", Id: "thisIsADeviceId"},
	}
	var requestHeader http.Header
	var requestURL url.URL
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHeader = r.Header
		requestURL = *r.URL
		w.Header().Set("Content-Type", "application/json")
		responseToSend, _ := json.Marshal(*response)
		w.Write(responseToSend)
	}))
	defer ts.Close()
	s := NewSession()
	s.client = new(http.Client)
	s.Webservices = make(map[string]models.WebService)
	s.Webservices[fmip.URLKey] = models.WebService{Url: ts.URL}
	err := s.PopulateDevices()
	if err != nil {
		t.Errorf("Session.Login() error = %v, wantErr %v", err, nil)
	}
	if requestURL.Path != fmip.RefreshClientURL {
		t.Errorf("Session.Login() error = %v, wantErr %v", requestURL.Path, fmip.RefreshClientURL)
	}
	if !reflect.DeepEqual(response.Devices, s.Devices) {
		t.Errorf("Session.Login() error = %v, wantErr %v", s.Devices, response.Devices)
	}
	if !checkRequestHeader(requestHeader) {
		t.Errorf("Session.Login() error = Bad request header")
	}
}

func TestSession_SendMessage(t *testing.T) {
	response := new(fmip.ResponseBody)
	var requestHeader http.Header
	var requestURL url.URL
	expectedBody := fmip.SendMessageRequestBody{RequestCommonBody: fmip.RequestCommonBody{Dsid: "Dsid", ClientId: "whatever"}, Device: "deviceID", Subject: "subject", Sound: false, UserText: true, Text: "this is a message"}
	var requestBody fmip.SendMessageRequestBody
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHeader = r.Header
		requestURL = *r.URL
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		responseToSend, _ := json.Marshal(*response)
		w.Write(responseToSend)
	}))
	defer ts.Close()
	s := NewSession()
	s.client = new(http.Client)
	s.Webservices = make(map[string]models.WebService)
	s.Webservices[fmip.URLKey] = models.WebService{Url: ts.URL}
	s.User.Dsid = expectedBody.Dsid
	err := s.SendMessage(expectedBody.Text, expectedBody.Device, expectedBody.Subject)
	if err != nil {
		t.Errorf("Session.Login() error = %v, wantErr %v", err, nil)
	}
	if requestURL.Path != fmip.SendMessageURL {
		t.Errorf("Session.Login() error = %v, wantErr %v", requestURL.Path, fmip.SendMessageURL)
	}
	if !reflect.DeepEqual(requestBody, expectedBody) {
		t.Errorf("Session.Login() error = %v, wantErr %v", requestBody, expectedBody)
	}
	if !checkRequestHeader(requestHeader) {
		t.Errorf("Session.Login() error = Bad request header")
	}
}

func checkRequestHeader(h http.Header) bool {
	return h.Get("Content-Type") == "application/json" && h.Get("Origin") == "https://www.icloud.com"
}
