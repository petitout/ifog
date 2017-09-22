package ifog

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/petitout/ifog/fmip"
	"github.com/petitout/ifog/login"
	"github.com/petitout/ifog/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type session struct {
	client      *http.Client
	cookies     []*http.Cookie
	User        models.User
	Webservices map[string]models.WebService
	Devices     []models.Device
}

func NewSession() *session {
	return new(session)
}

func (this *session) Login1(requestBody login.RequestBody) error {
	this.client = &http.Client{}
	data, e := json.Marshal(requestBody)
	if e != nil {
		return e
	}
	req, e := http.NewRequest(login.Method, login.URL, bytes.NewReader(data))
	if e != nil {
		return e
	}
	setCommonHeaders(req)
	var resp *http.Response
	if resp, e = this.client.Do(req); e != nil {
		return e
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	this.cookies = resp.Cookies()
	var responseBody login.ResponseBody
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		return err
	}
	this.User = responseBody.User
	this.Webservices = responseBody.Webservices
	return nil
}

func (this *session) Login(requestBody login.RequestBody) error {
	this.client = &http.Client{}
	responseBody := login.ResponseBody{}
	err := this.sendRequest(requestBody, login.URL, &responseBody, true)
	if err != nil {
		return err
	}
	this.User = responseBody.User
	this.Webservices = responseBody.Webservices
	return nil
}



func (this *session) PopulateDevices() error {
	requestBody := fmip.RefreshClientRequestBody{fmip.RequestCommonBody{this.User.Dsid, "whatever"}, fmip.DefaultClientContext}
	responseBody := fmip.ResponseBody{}
	err := this.sendRequest(requestBody, this.Webservices[fmip.URL_KEY].Url+fmip.REFRESH_CLIENT_URL, &responseBody, false)
	if err != nil {
		return err
	}
	this.Devices = responseBody.Devices
	return nil
}

func (this *session) SendMessage(msg string, deviceId string, subject string) error {
	requestBody := fmip.SendMessageRequestBody{fmip.RequestCommonBody{this.User.Dsid, "whatever"}, deviceId, subject, false, true, msg}
	err := this.sendRequest(requestBody, this.Webservices[fmip.URL_KEY].Url+fmip.SEND_MESSAGE_URL, nil, false)
	if err != nil {
		return err
	}
	return nil
}


func (this *session) sendRequest(body interface{}, url string, responseBody interface{}, saveCookies bool) error {
	if this.client == nil {
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
	for _, cookie := range this.cookies {
		if strings.Index(cookie.String(), "X-APPLE-WEBAUTH-LOGIN") != -1 || strings.Index(cookie.String(), "X-APPLE-WEBAUTH-USER") != -1 {
			req.Header.Add("Cookie", cookie.Name+"=\""+cookie.Value+"\"")
		}
	}
	setCommonHeaders(req)
	var resp *http.Response
	if resp, e = this.client.Do(req); e != nil {
		return e
	}
	defer resp.Body.Close()
	rawBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	if (saveCookies) {
		this.cookies = resp.Cookies()
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