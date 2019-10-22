package controllers

import (
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
}

// Post ...
// @Title Create
// @Description create Validacion
// @Param	docID	query	string	true		"ID del documento"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [post]
func (c *ValidacionController) Post() {
	var alertErr models.Alert
	DocumentID := c.GetString("docID")
	validacion := validar(DocumentID)
	if validacion != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = validacion
	} else {
		alertErr.Type = "Failure"
		alertErr.Code = "404"
		alertErr.Body = "Error al validar el documento el flujo"
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()
}

func validar(docID string) interface{} {
	flujoID := ObtenerFlujoID(docID)
	if flujoID != "nil" {
		tareaID := ObtenerTareaID(docID, flujoID, beego.AppConfig.String("user"), "SerialDocumentReview")
		if tareaID != "nil" {
			review := IniciarReview(tareaID)
			if review != "nil" {
				// return review
				tareaID2 := ObtenerTareaID(docID, flujoID, beego.AppConfig.String("user"), "SerialDocumentReview")
				if tareaID != "nil" {
					ResValidacion := RealizarValidacion(tareaID2)
					if ResValidacion != nil {
						// logs.Info(ResValidacion)
						return ResValidacion
					}
				}

			}
		}

	}
	return nil
}

func ObtenerFlujoID(docID string) string {
	var respuesta interface{}
	var StringRespueta string
	endpoint := "id/" + docID + "/@workflow"
	respuesta = models.GetNuxeo(endpoint)
	if respuesta != nil {
		respuesta = models.GetElemento(respuesta, "entries")
		StringRespueta = models.GetElementoMaptoString(respuesta, "id")
		return StringRespueta
	} else {
		logs.Error("Error al obtener el ID del flujo")
	}
	return "nil"
}

func ObtenerTareaID(docID string, flujoID string, userID string, nameFlujo string) string {
	var respuesta interface{}
	var StringRespueta string
	endpoint := "id/" + docID + "/@task?userId=" + userID + "&workflowInstanceId=" + flujoID + "&workflowModelName=" + nameFlujo
	respuesta = models.GetNuxeo(endpoint)
	if respuesta != nil {
		respuesta = models.GetElemento(respuesta, "entries")
		StringRespueta = models.GetElementoMaptoString(respuesta, "id")
		return StringRespueta
	} else {
		logs.Error("Error al obtener el ID de la tarea")
	}
	return "nil"
}

func IniciarReview(tareaID string) interface{} {
	var respuesta interface{}
	endpoint := "task/" + tareaID + "/start_review"
	requestBody := "{\"entity-type\":\"task\",\n\"id\":\"" + tareaID +
		"\",\n\"variables\":\n{\"comment\":\"Se validan los documentos.\",\n\"participants\":[\"user:" +
		beego.AppConfig.String("user") + "\"],\n\"validationOrReview\":\"validation\"}\n}"
	respuesta = models.PutNuxeo(endpoint, requestBody)
	if respuesta != nil {
		respuesta = models.GetElemento(respuesta, "id")
		return respuesta
	} else {
		logs.Error("Error el review de la tarea")
	}
	return nil
}

func RealizarValidacion(tareaID string) interface{} {
	var respuesta interface{}
	endpoint := "task/" + tareaID + "/validate"
	requestBody := "{\"entity-type\":\"task\",\n\"id\":\"" + tareaID +
		"\",\n\"variables\":\n{\"comment\":\"El supervisor aprueba los documentos.\",\n\"participants\":[\"user:" +
		beego.AppConfig.String("user") + "\"],\n\"initiatorComment\":\"Se validan los documentos.\"}\n}"
	respuesta = models.PutNuxeo(endpoint, requestBody)
	if respuesta != nil {
		return respuesta
	} else {
		logs.Error("Error al validar el documento")
	}
	return nil
}

func Comprobante(docID string) interface{} {
	var respuesta interface{}
	var StringRespueta string
	endpoint := "id/" + docID
	respuesta = models.GetNuxeo(endpoint)
	if respuesta != nil {
		// respuesta = models.GetElemento(respuesta, "entries")
		// StringRespueta = models.GetElementoMaptoString(respuesta, "id")
		return StringRespueta
	} else {
		logs.Error("Error al obtener el ID del flujo")
	}
	return nil
}
