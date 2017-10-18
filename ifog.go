package ifog

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/petitout/ifog/fmip"
	"github.com/petitout/ifog/login"
	"github.com/petitout/ifog/models"
)

// Session ICloud session structure
type Session struct {
	client      *http.Client
	cookies     []*http.Cookie
	User        models.User
	Webservices map[string]models.WebService
	Devices     []models.Device
}

// NewSession creates a new Icloud session
func NewSession() *Session {
	return new(Session)
}

// Login allows to login into ICloud
func (s *Session) Login(loginURL string, requestBody login.RequestBody) error {
	s.client = &http.Client{}
	responseBody := login.ResponseBody{}
	err := s.sendRequest(requestBody, loginURL, &responseBody, true)
	if err != nil {
		return err
	}
	s.User = responseBody.User
	s.Webservices = responseBody.Webservices
	return nil
}

// PopulateDevices retrieves device information associated with the currently logged in ICloud user
func (s *Session) PopulateDevices() error {
	requestBody := fmip.RefreshClientRequestBody{RequestCommonBody: fmip.RequestCommonBody{Dsid: s.User.Dsid, ClientId: "whatever"}, ClientContext: fmip.DefaultClientContext}
	responseBody := fmip.ResponseBody{}
	err := s.sendRequest(requestBody, s.Webservices[fmip.URLKey].Url+fmip.RefreshClientURL, &responseBody, false)
	if err != nil {
		return err
	}
	s.Devices = responseBody.Devices
	return nil
}

// SendMessage sends a message to the given device
func (s *Session) SendMessage(msg string, deviceID string, subject string) error {
	requestBody := fmip.SendMessageRequestBody{RequestCommonBody: fmip.RequestCommonBody{Dsid: s.User.Dsid, ClientId: "whatever"}, Device: deviceID, Subject: subject, Sound: false, UserText: true, Text: msg}
	err := s.sendRequest(requestBody, s.Webservices[fmip.URLKey].Url+fmip.SendMessageURL, nil, false)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) sendRequest(body interface{}, url string, responseBody interface{}, saveCookies bool) error {
	if s.client == nil {
		return errors.New("You have to login first")
	}
	data, e := json.Marshal(body)
	if e != nil {
		return e
	}
	req, e := http.NewRequest(fmip.Method, url, bytes.NewReader(data))
	if e != nil {
		return e
	}
	for _, cookie := range s.cookies {
		if strings.Index(cookie.String(), "X-APPLE-WEBAUTH-LOGIN") != -1 || strings.Index(cookie.String(), "X-APPLE-WEBAUTH-USER") != -1 {
			req.Header.Add("Cookie", cookie.Name+"=\""+cookie.Value+"\"")
		}
	}
	setCommonHeaders(req)
	var resp *http.Response
	if resp, e = s.client.Do(req); e != nil {
		return e
	}
	defer resp.Body.Close()
	rawBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	if saveCookies {
		s.cookies = resp.Cookies()
	}
	if responseBody != nil {
		err := json.Unmarshal(rawBody, responseBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func setCommonHeaders(request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", "https://www.icloud.com")
}

func (s *Session) Devices() []models.Device {
	return s.Devices
}
