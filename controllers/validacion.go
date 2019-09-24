package controllers

import (
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/nuxeo_mid/models"
)

// ValidacionController operations for Validacion
type ValidacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ValidacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Create
// @Description create Validacion
// @Param	docID	query	string	true		"ID del documento"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *ValidacionController) Post() {
	var alertErr models.Alert
	DocumentID := c.GetString("docID")
	// fmt.Println(DocumentID)
	alertErr.Type = "OK"
	alertErr.Code = "200"
	alertErr.Body = ObtenerFlujoID(DocumentID)
	// alertErr.Body = models.GetNuxeo("me")
	c.Data["json"] = alertErr
	c.ServeJSON()
}

func ObtenerFlujoID(docID string) interface{} {
	var respuesta interface{}
	respuesta = models.GetNuxeo("id", docID, "@workflow")
	// result := respuesta["Body"]
	logs.Error(respuesta)
	// logs.Error(reflect.ValueOf(respuesta))
	respuesta = models.GetElemento(respuesta, "entries")
	// respuesta = models.GetElemento(respuesta, "attachedDocumentIds")
	logs.Emergency(respuesta)

	value := reflect.ValueOf(respuesta)
	c := value.Index(0).Interface().(map[string]interface{})
	logs.Info(c["attachedDocumentIds"])
	logs.Info(c)

	value2 := reflect.ValueOf(c["attachedDocumentIds"])
	d := value2.Index(0).Interface().(map[string]interface{})
	logs.Info(d["id"])

	return c["attachedDocumentIds"]
}

// GetAll ...
// @Title GetAll
// @Description get Validacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Validacion
// @Failure 403
// @router / [get]
func (c *ValidacionController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Validacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Validacion	true		"body for Validacion content"
// @Success 200 {object} models.Validacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ValidacionController) Put() {

}
