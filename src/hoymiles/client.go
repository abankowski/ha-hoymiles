package hoymiles

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	// BaseURL = "https://global.hoymiles.com/platform/api/gateway/"
	BaseURL = "http://127.0.0.1:4010/"
)

type Client struct {
	BaseURL    *url.URL
	user       string
	password   string
	ctx        context.Context
	HTTPClient *http.Client
	session    string
}

type ResponseEnvelope struct {
	Message      string  `json:"message"`
	Status       string  `json:"status"`
	SystemNotice *string `json:"systemNotice,omitempty"`
}

func NewClient(user string, password string, ctx context.Context, eager ...bool) (*Client, error) {

	u, e := url.Parse(BaseURL)

	if e != nil {
		return nil, e
	}

	c := Client{
		BaseURL:  u,
		user:     user,
		password: fmt.Sprintf("%x", md5.Sum([]byte(password))), // Sic! Seriously...
		ctx:      ctx,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	if len(eager) > 0 && eager[0] {
		c.refreshSession()
	}

	return &c, nil
}

func (c *Client) makeRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.AddCookie(&http.Cookie{Name: "hm_token_language", Value: "en_us"})

	if c.session != "" {
		req.AddCookie(&http.Cookie{Name: "hm_token", Value: c.session})
	}

	res, err := c.HTTPClient.Do(req)

	body, err := io.ReadAll(res.Body)

	if err == nil {
		env := &ResponseEnvelope{}

		if err = json.Unmarshal(body, env); err == nil {
			if env.Status != "0" { // string 0 is ok and 1 is error
				err = fmt.Errorf("request failed - %s", env.Message)
				log.Fatalf("Response status=%s indicates erro - %sr", env.Status, err)

				return nil, err
			}
		} else {
			fmt.Println(string(body))
			log.Fatal("Cannot decode response due to unmarshalling error", err)

			return nil, err
		}
	}

	return body, nil
}

func (c *Client) refreshSession() error {

	fmt.Println("refresh session with pass", c.password)

	var buf bytes.Buffer

	payload := AuthLoginRequest{
		ERRORBACK:      true,
		WAITINGPROMISE: true,
		Body: struct {
			Password string `json:"password"`
			UserName string `json:"user_name"`
		}{
			Password: c.password,
			UserName: c.user,
		},
	}

	json.NewEncoder(&buf).Encode(payload)

	req, err := http.NewRequest(http.MethodPost, c.url("/iam/auth_login"), &buf)

	if err != nil {
		return err
	}

	body, err := c.makeRequest(req)

	fmt.Println(string(body))

	if err != nil {
		return err
	}

	resp := &AuthLoginResponse{}
	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}

	c.session = resp.Data.Token

	return nil
}

func (c *Client) url(path string) string {
	u := *c.BaseURL
	u.Path = path
	return u.String()
}

func (c *Client) GetUserMe() (*GetUserMeResponse, error) {

	if c.session == "" {
		c.refreshSession()
	}

	var buf bytes.Buffer

	payload := GetUserMeRequest{
		ERRORBACK:      true,
		WAITINGPROMISE: true,
	}

	json.NewEncoder(&buf).Encode(payload)

	req, _ := http.NewRequest(http.MethodPost, c.url("/iam/user_me"), &buf)

	body, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	resp := GetUserMeResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetDeviceAll(station int) (*StationSelectDeviceAllResponse, error) {

	if c.session == "" {
		c.refreshSession()
	}

	var buf bytes.Buffer

	payload := StationRequest{
		WAITINGPROMISE: true,
		Body: struct {
			Id int `json:"id"`
		}{Id: station},
	}

	json.NewEncoder(&buf).Encode(payload)

	req, _ := http.NewRequest(http.MethodPost, c.url("/pvm/station_select_device_all"), &buf)

	body, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	resp := StationSelectDeviceAllResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetDataDetails(station int, mi_id int, mi_sn string, time string) (*DataFindDetailsResponse, error) {
	if c.session == "" {
		c.refreshSession()
	}

	var buf bytes.Buffer

	payload := DataFindDetailsRequest{
		WAITINGPROMISE: true,
		Body: struct {
			MiId     int    `json:"mi_id"`
			MiSn     string `json:"mi_sn"`
			Port     int    `json:"port"`
			Sid      int    `json:"sid"`
			Time     string `json:"time"`
			WarnCode int    `json:"warn_code"`
		}{
			MiId:     mi_id,
			MiSn:     mi_sn,
			Port:     1,
			Sid:      station,
			Time:     time,
			WarnCode: 0,
		},
	}

	json.NewEncoder(&buf).Encode(payload)

	req, _ := http.NewRequest(http.MethodPost, c.url("/pvm-data/data_find_details"), &buf)

	body, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	resp := DataFindDetailsResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetDeviceTree(station int) (*StationSelectDeviceOfTreeResponse, error) {
	if c.session == "" {
		c.refreshSession()
	}

	var buf bytes.Buffer

	payload := StationRequest{
		WAITINGPROMISE: true,
		Body: struct {
			Id int `json:"id"`
		}{Id: station},
	}

	json.NewEncoder(&buf).Encode(payload)

	req, _ := http.NewRequest(http.MethodPost, c.url("/pvm/station_select_device_of_tree"), &buf)

	body, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	resp := StationSelectDeviceOfTreeResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetStation(station int) (*StationFindResponse, error) {
	if c.session == "" {
		c.refreshSession()
	}

	var buf bytes.Buffer

	payload := StationRequest{
		WAITINGPROMISE: true,
		Body: struct {
			Id int `json:"id"`
		}{Id: station},
	}

	json.NewEncoder(&buf).Encode(payload)

	req, _ := http.NewRequest(http.MethodPost, c.url("/pvm/station_find"), &buf)

	body, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	resp := StationFindResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetRealData(station int) (*DataCountStationRealDataResponse, error) {
	if c.session == "" {
		c.refreshSession()
	}

	var buf bytes.Buffer

	payload := StationDataRequest{
		WAITINGPROMISE: true,
		Body: struct {
			Sid int `json:"sid"`
		}{Sid: station},
	}

	json.NewEncoder(&buf).Encode(payload)

	req, _ := http.NewRequest(http.MethodPost, c.url("/pvm-data/data_count_station_real_data"), &buf)

	body, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	resp := DataCountStationRealDataResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
