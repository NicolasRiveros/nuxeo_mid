package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"

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
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Workflow
// @Param	docID	query	string	true		"ID del documento"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [post]
func (c *WorkflowController) Post() {
	var alertErr models.Alert
	DocumentID := c.GetString("docID")
	Disparo := DisparoFlujo(DocumentID)

	if Disparo != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = Disparo
	} else {
		alertErr.Type = "Failure"
		alertErr.Code = "404"
		alertErr.Body = "Error al disparar el flujo"
		c.Ctx.Output.SetStatus(404)
		// c.Data["json"] = alertErr
		// c.ServeJSON()
	}

	// alertErr.Body = models.GetNuxeo("me")
	c.Data["json"] = alertErr

	c.ServeJSON()
}

// @Title DisparoFlujo
// @Description funcion para dispars flujo
func DisparoFlujo(docID string) interface{} {
	var respuesta interface{}
	requestBody, errBody := json.Marshal(map[string]string{
		"entity-type":         "workflow",
		"workflowModelName":   "SerialDocumentReview",
		"attachedDocumentIds": "[" + docID + "]",
	})
	if errBody != nil {
		logs.Error("fallo el objeto a enviar: ", errBody)
	}
	respuesta = models.PostNuxeo("id", docID, requestBody, "@workflow")
	if respuesta != nil {
		return respuesta
	} else {
		logs.Error("Error al disparar el flujo")
	}
	return nil
}

// Delete ...
// @Title Delete
// @Description delete the Workflow
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *WorkflowController) Delete() {
	idStr := c.GetString("id")
	fmt.Println(idStr)
}
