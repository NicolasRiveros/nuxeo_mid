package models

import (
	"fmt"
	"reflect"
)

func GetElemento(objeto interface{}, item string) interface{} {
	var subobjeto interface{}
	subobjeto = objeto.(map[string]interface{})[item]
	return subobjeto
}

func GetElementoMaptoString(objeto interface{}, item string) string {
	value := reflect.ValueOf(objeto)
	var resuesta string
	aux := value.Index(0).Interface().(map[string]interface{})
	// logs.Info(aux["id"])
	// logs.Info(aux)
	// logs.Error(reflect.TypeOf(aux["id"]))
	// aux2 := reflect.ValueOf(aux["id"]).String
	// resuesta = reflect.String(aux2)
	// resuesta = strconv.Quote(aux["id"])
	resuesta = fmt.Sprintf("%v", aux["id"])
	return resuesta
}
