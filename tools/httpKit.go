package tools

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"io/ioutil"
	"net/http"
)

//get get data from givin url and return map as value
func Get(url string) (map[string]interface{}, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	rstMap := make(map[string]interface{})
	json.Unmarshal(resp, &rstMap)
	return rstMap, nil
}

//get get data from givin url and return map as value
func PostAuth(url, reqBody, user, pass string) (map[string]interface{}, error) {
	buf := bytes.NewBuffer([]byte(reqBody))
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if user != "" && pass != "" {
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.Conf.Ela.JsonrpcUser+":"+config.Conf.Ela.JsonrpcPassword)))
	}
	r, err := http.DefaultClient.Do(req)
	//r, err := http.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	rstMap := make(map[string]interface{})
	json.Unmarshal(resp, &rstMap)
	if rstMap == nil || len(rstMap) == 0 {
		return nil, errors.New("Invalid Post Request to Server")
	}
	return rstMap, nil
}

//get get data from givin url and return map as value
func Post(url string, reqBody string) (map[string]interface{}, error) {
	return PostAuth(url, reqBody, "", "")
}
