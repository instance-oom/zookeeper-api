package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"zookeeper-api/config"

	"github.com/astaxie/beego"
	"github.com/samuel/go-zookeeper/zk"
)

var zkClient *zk.Conn

// Base - Provider some common func, like 'Error()' .etc
type baseController struct {
	beego.Controller
	path string
}

// Init - init basic info
func Init() {
	var err error
	zkHosts := config.ZKHosts
	zkClient, _, err = zk.Connect(zkHosts, time.Second*10)
	if err != nil {
		fmt.Printf("Connect zk cluster failed! \n Error: %s \n Hosts: %v", err.Error(), zkHosts)
		os.Exit(2)
	}
}

// Prepare - Format path before exec real action
func (base *baseController) Prepare() {
	path := base.Ctx.Input.Param(":splat")
	if path == "" {
		path = "/"
	} else {
		if index := strings.Index(path, "/"); index != 0 {
			path = "/" + path
		}
	}
	base.path = path
}

// Json - Return json data
func (base *baseController) JSON(data interface{}) {
	base.Data["json"] = data
	base.ServeJSON()
}

// Error - Unified processing error
func (base *baseController) Error(status int, msg string) {
	errData := map[string]interface{}{
		"Code":   status,
		"Detail": msg,
	}

	base.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	base.Ctx.ResponseWriter.WriteHeader(status)
	body, _ := json.Marshal(errData)
	base.Ctx.ResponseWriter.Write(body)
}
