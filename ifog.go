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

func (this *session) Login(requestBody login.RequestBody) error {
	this.client = &http.Client{}
	data, e := json.Marshal(requestBody)
	if e != nil {
		return e
	}
	req, e := http.NewRequest(login.Method, login.URL, bytes.NewReader(data))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://www.icloud.com")
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

func (this *session) PopulateDevices() error {
	if this.client == nil {
		return errors.New("You have to login first")
	}
	requestBody := fmip.RequestBody{this.User.Dsid, "whatever", fmip.DefaultClientContext}
	data, e := json.Marshal(requestBody)
	if e != nil {
		return e
	}
	req, e := http.NewRequest(fmip.Method, this.Webservices[fmip.URL_KEY].Url+"/fmipservice/client/web/refreshClient", bytes.NewReader(data))
	if e != nil {
		return e
	}
	for _, cookie := range this.cookies {
		if strings.Index(cookie.String(), "X-APPLE-WEBAUTH-LOGIN") != -1 || strings.Index(cookie.String(), "X-APPLE-WEBAUTH-USER") != -1 {
			req.Header.Add("Cookie", cookie.Name+"=\""+cookie.Value+"\"")
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Origin", "https://www.icloud.com")
	var resp *http.Response
	if resp, e = this.client.Do(req); e != nil {
		return e
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	var responseBody fmip.ResponseBody
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		return err
	}
	this.Devices = responseBody.Devices
	return nil
}
