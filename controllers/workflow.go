package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/nuxeo_mid/models"
)

// WorkflowController operations for Workflow
type WorkflowController struct {
	beego.Controller
}

// URLMapping ...
func (c *WorkflowController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Create
// @Description create Workflow
// @Param	docID	query	string	true		"ID del documento"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *WorkflowController) Post() {
	var alertErr models.Alert
	// alertas := append([]interface{}{"Response:"})
	// var resultadoGet map[string]interface{}

	DocumentID := c.GetString("docID")
	fmt.Println(DocumentID)
	alertErr.Type = "OK"
	alertErr.Code = "200"
	// models.BasicAuth()
	// request.SetHeader()
	// request.SetHeader("Authorization", "Basic "+models.BasicAuth("username1", "password123"))
	// errGET := request.GetJson("https://"+beego.AppConfig.String("NovedadesApiMongoService")+"novedad", &resultadoGet)
	// fmt.Println(errGET)
	// req, err := http.NewRequest("GET", "https://"+beego.AppConfig.String("urlNuxeo")+"me", nil)
	// // req.Header.Add("Authorization", "Basic "+models.BasicAuth())
	// // http.SetBasicAuth()
	// req.SetBasicAuth(beego.AppConfig.String("user"), beego.AppConfig.String("password"))
	// // fmt.Println(err)
	// logs.Error(err)
	// logs.Info(req)
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// logs.Info(resp)
	// logs.Error(err)
	alertErr.Body = peticion()
	c.Data["json"] = alertErr
	c.ServeJSON()
}

func peticion() interface{} {
	url := "https://" + beego.AppConfig.String("urlNuxeo") + "me"

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(beego.AppConfig.String("user"), beego.AppConfig.String("password")))
	if err != nil {
	}
	resp, err3 := client.Do(req)

	if err3 != nil {
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		var data interface{}
		json.Unmarshal([]byte(bodyString), &data)
		return data
		if err2 != nil {
		}
	}
	return nil

}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// GetAll ...
// @Title GetAll
// @Description get Workflow
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Workflow
// @Failure 403
// @router / [get]
func (c *WorkflowController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Workflow
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Workflow	true		"body for Workflow content"
// @Success 200 {object} models.Workflow
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WorkflowController) Put() {

}
