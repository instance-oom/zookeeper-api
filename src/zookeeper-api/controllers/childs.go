package controllers

import "github.com/samuel/go-zookeeper/zk"

// ChildsController - Handle the http request with childs
type ChildsController struct {
	baseController
}

// Get - List children for the path
func (child *ChildsController) Get() {
	childs, _, err := zkClient.Children(child.path)
	if err != nil {
		if err == zk.ErrNoNode {
			child.Error(404, err.Error())
		} else {
			child.Error(500, err.Error())
		}
	} else {
		result := map[string]interface{}{
			"Path":   child.path,
			"Childs": childs,
		}
		child.JSON(result)
	}
}
