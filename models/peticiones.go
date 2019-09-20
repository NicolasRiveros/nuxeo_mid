package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetNuxeo(endpoint string) interface{} {
	url := "https://" + beego.AppConfig.String("urlNuxeo") + endpoint

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
