package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetNuxeo(endpoint string) interface{} {
	url := "https://" + beego.AppConfig.String("urlNuxeo") + endpoint
	// if &opcionendpoint != nil {
	// 	url = url + "/" + ID + "/" + opcionendpoint
	// }

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+BasicAuth(beego.AppConfig.String("user"), beego.AppConfig.String("password")))
	if err != nil {
	}
	resp, err3 := client.Do(req)

	if err3 != nil {
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			logs.Error("fallo el leer el body de la peticion")
		}
		bodyString := string(bodyBytes)
		var data interface{}
		json.Unmarshal([]byte(bodyString), &data)
		return data
	}
	return nil

}

func PostNuxeo(endpoint string, ID string, objeto []byte, opcionendpoint string) interface{} {
	url := "https://" + beego.AppConfig.String("urlNuxeo") + endpoint
	if &opcionendpoint != nil {
		url = url + "/" + ID + "/" + opcionendpoint
	}

	var client http.Client

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(objeto))
	req.Header.Add("Authorization", "Basic "+BasicAuth(beego.AppConfig.String("user"), beego.AppConfig.String("password")))
	if err != nil {
		logs.Error("error en post de nuxeo: ", err)
	}
	resp, err3 := client.Do(req)

	if err3 != nil {
		logs.Info("error en client.Do", err3)
	}

	defer resp.Body.Close()
	logs.Info(resp.StatusCode)
	if resp.StatusCode == 200 || resp.StatusCode == 201 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			logs.Error("fallo el leer el body de la peticion")
		}
		bodyString := string(bodyBytes)
		var data interface{}
		json.Unmarshal([]byte(bodyString), &data)
		return data
	}
	return nil

}

func PutNuxeo(endpoint string, objeto string) interface{} {
	url := "https://" + beego.AppConfig.String("urlNuxeo") + endpoint
	var client http.Client
	payload := strings.NewReader(objeto)
	req, err := http.NewRequest("PUT", url, payload)
	req.Header.Add("Authorization", "Basic "+BasicAuth(beego.AppConfig.String("user"), beego.AppConfig.String("password")))
	if err != nil {
		logs.Error("error en post de nuxeo: ", err)
	}
	resp, err3 := client.Do(req)
	if err3 != nil {
		logs.Info("error en client.Do", err3)
	}

	defer resp.Body.Close()
	logs.Error(resp.StatusCode)
	if resp.StatusCode == 200 || resp.StatusCode == 201 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			logs.Error("fallo el leer el body de la peticion")
		}
		bodyString := string(bodyBytes)
		var data interface{}
		json.Unmarshal([]byte(bodyString), &data)
		return data
	} else {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			logs.Error("fallo el leer el body de la peticion")
		}
		bodyString := string(bodyBytes)
		logs.Error(bodyString)
	}
	return nil

}

func DeleteNuxeo(endpoint string) interface{} {
	url := "https://" + beego.AppConfig.String("urlNuxeo") + endpoint
	var client http.Client

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", "Basic "+BasicAuth(beego.AppConfig.String("user"), beego.AppConfig.String("password")))
	if err != nil {
	}
	resp, err3 := client.Do(req)

	if err3 != nil {
	}

	defer resp.Body.Close()

	if resp.StatusCode == 204 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			logs.Error("fallo el leer el body de la peticion")
		}
		bodyString := string(bodyBytes)
		var data interface{}
		json.Unmarshal([]byte(bodyString), &data)
		return data
	}
	return nil

}
