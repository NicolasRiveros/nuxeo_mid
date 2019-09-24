package controllers

import (
	"github.com/astaxie/beego"
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
	alertErr.Body = validar(DocumentID)
	// alertErr.Body = models.GetNuxeo("me")
	c.Data["json"] = alertErr
	c.ServeJSON()
}

func validar(docID string) string {
	flujoID := ObtenerFlujoID(docID)
	tareaID := ObtenerTareaID(docID, flujoID, beego.AppConfig.String("user"), "SerialDocumentReview")
	return tareaID
}

func ObtenerFlujoID(docID string) string {
	var respuesta interface{}
	var StringRespueta string
	endpoint := "id/" + docID + "/@workflow"
	respuesta = models.GetNuxeo(endpoint)
	// result := respuesta["Body"]
	// logs.Error(respuesta)
	// logs.Error(reflect.ValueOf(respuesta))
	respuesta = models.GetElemento(respuesta, "entries")
	StringRespueta = models.GetElementoMaptoString(respuesta, "id")

	// value2 := reflect.ValueOf(c["attachedDocumentIds"])
	// d := value2.Index(0).Interface().(map[string]interface{})
	// logs.Info(d["id"])

	return StringRespueta
	// return c["attachedDocumentIds"]
}

func ObtenerTareaID(docID string, flujoID string, userID string, nameFlujo string) string {
	var respuesta interface{}
	var StringRespueta string
	endpoint := "id/" + docID + "/@task?userId=" + userID + "&workflowInstanceId=" + flujoID + "&workflowModelName=" + nameFlujo
	respuesta = models.GetNuxeo(endpoint)
	respuesta = models.GetElemento(respuesta, "entries")
	StringRespueta = models.GetElementoMaptoString(respuesta, "id")

	return StringRespueta
}

func IniciarReview(docID string) string {
	var respuesta interface{}
	var StringRespueta string
	endpoint := "id/" + docID + "/@workflow"
	respuesta = models.GetNuxeo(endpoint)
	// result := respuesta["Body"]
	// logs.Error(respuesta)
	// logs.Error(reflect.ValueOf(respuesta))
	respuesta = models.GetElemento(respuesta, "entries")
	StringRespueta = models.GetElementoMaptoString(respuesta, "id")

	// value2 := reflect.ValueOf(c["attachedDocumentIds"])
	// d := value2.Index(0).Interface().(map[string]interface{})
	// logs.Info(d["id"])

	return StringRespueta
	// return c["attachedDocumentIds"]
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
