package models

import (
	"github.com/astaxie/beego/logs"
)

func GetFlujoID(docID string) string {
	var respuesta interface{}
	var StringRespueta string
	var estadoFlujo string
	endpoint := "id/" + docID + "/@workflow"
	respuesta = GetNuxeo(endpoint)
	if respuesta != nil {
		respuesta = GetElemento(respuesta, "entries")
		estadoFlujo = GetElementoMaptoString(respuesta, "state")
		if(estadoFlujo != "Objeto de longitud cero") {
			StringRespueta = GetElementoMaptoString(respuesta, "id")
			return StringRespueta
		}
		if(estadoFlujo == "Objeto de longitud cero") {			
			return "NoID"
		}
		
	} else {
		logs.Error("Error al obtener el ID del flujo")
	}
	return "nil"
}

