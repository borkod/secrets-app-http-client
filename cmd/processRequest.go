package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

func (r *request) processRequest() (string, error) {
	const ConnectMaxWaitTime = 1 * time.Second
	const RequestMaxWaitTime = 5 * time.Second

	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: ConnectMaxWaitTime,
			}).DialContext,
		},
	}

	var req *http.Request

	if r.action == "GET" {
		req, _ = http.NewRequest(http.MethodGet, r.url+"/"+r.id, strings.NewReader(""))
	} else {
		var jsonStr = []byte(`{"plain_text": "` + r.data + `"}`)
		req, _ = http.NewRequest(http.MethodPost, r.url+"/", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	var respData string

	if r.action == "GET" {
		var data secretData
		err = json.Unmarshal(buf.Bytes(), &data)

		if err != nil {
			return "", err
		}
		respData = data.Data
	} else {
		var id secretID
		err = json.Unmarshal(buf.Bytes(), &id)

		if err != nil {
			return "", err
		}
		respData = r.url + "/" + id.Id
	}

	return respData, nil
}
