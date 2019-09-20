package models

import (
	"encoding/base64"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func BasicAuth() string {
	fmt.Println("tiempo antes de correccion")
	username := beego.AppConfig.String("user")
	password := beego.AppConfig.String("password")
	auth := username + ":" + password
	logs.Error(base64.StdEncoding.EncodeToString([]byte(auth)))

	// ctx := *context.Context{
	// 	Request:        req,
	// 	ResponseWriter: rw,
	// }

	// request.SetHeader()

	return base64.StdEncoding.EncodeToString([]byte(auth))
}
