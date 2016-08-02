package controllers

import (
	"encoding/json"
	"strings"
	"zookeeper-api/models"

	"github.com/samuel/go-zookeeper/zk"
)

// NodeController - Handle the http request for node
type NodeController struct {
	baseController
	znode models.ZNode
}

// Prepare - Format request body beform exec real action
func (node *NodeController) Prepare() {
	path := node.Ctx.Input.Param(":splat")
	if path == "" {
		path = "/"
	} else {
		if index := strings.Index(path, "/"); index != 0 {
			path = "/" + path
		}
	}
	node.path = path

	if node.Ctx.Request.Method == "POST" || node.Ctx.Request.Method == "PUT" {
		err := json.Unmarshal(node.Ctx.Input.RequestBody, &node.znode)
		if err != nil {
			node.Error(422, "Invalid request body")
			node.StopRun()
		}
		node.znode.Path = path
	}
}

// Get - Get node's value
func (node *NodeController) Get() {
	value, stat, err := zkClient.Get(node.path)
	if err != nil {
		if err == zk.ErrNoNode {
			node.Error(404, err.Error())
		} else {
			node.Error(500, err.Error())
		}
	} else {
		result := map[string]interface{}{
			"Path":         node.path,
			"Value":        string(value),
			"Version":      stat.Version,
			"InDate":       stat.Ctime,
			"LastEditDate": stat.Mtime,
			"ChildNum":     stat.NumChildren,
		}
		node.JSON(result)
	}
}

// Post - Create node
func (node *NodeController) Post() {
	path, err := zkClient.Create(node.znode.Path, []byte(node.znode.Value), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		if err == zk.ErrNodeExists {
			node.Error(409, err.Error())
		} else {
			node.Error(500, err.Error())
		}
	} else {
		result := map[string]string{
			"Path": path,
		}
		node.JSON(result)
	}
}

// Put - Update node
func (node *NodeController) Put() {
	_, err := zkClient.Set(node.znode.Path, []byte(node.znode.Value), -1)
	if err != nil {
		if err == zk.ErrNoNode {
			node.Error(404, err.Error())
		} else {
			node.Error(500, err.Error())
		}
	}
	node.Ctx.ResponseWriter.Write([]byte(""))
}

// Delete - Delete node
func (node *NodeController) Delete() {
	err := zkClient.Delete(node.path, -1)
	if err != nil {
		if err == zk.ErrNoNode {
			node.Error(404, err.Error())
		} else {
			node.Error(500, err.Error())
		}
	}
	node.Ctx.ResponseWriter.Write([]byte(""))
}
