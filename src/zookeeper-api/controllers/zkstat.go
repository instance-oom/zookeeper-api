package controllers

import (
	"regexp"
	"strings"
	"time"
	"zookeeper-api/models"

	"github.com/samuel/go-zookeeper/zk"
)

// ZKStatController - ZK cluster state
type ZKStatController struct {
	baseController
	zkhosts []string
}

// Prepare - Override baseController
func (zkStat *ZKStatController) Prepare() {
	ipStr := zkStat.Ctx.Input.Param(":ip")
	if ipStr != "" {
		zkStat.zkhosts = strings.Split(ipStr, ",")
	} else {
		zkStat.Error(400, "Zookeeper ip is required")
		zkStat.StopRun()
	}
	pattern := "^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[[1-9][0-9]|[0-9])\\.){3}((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[[1-9][0-9]|[1-9])):[0-9]{1,5}$"
	for _, host := range zkStat.zkhosts {
		isMatch, err := regexp.MatchString(pattern, host)
		if !isMatch || err != nil {
			zkStat.Error(400, host+" is invalid")
			zkStat.StopRun()
		}
	}
}

// Get - List cluster server's state
func (zkStat *ZKStatController) Get() {
	states, _ := zk.FLWSrvr(zkStat.zkhosts, time.Second*10)
	result := make([]models.ZKStat, len(states))
	for i, stat := range states {
		tempStat := models.ZKStat{
			Server:      zkStat.zkhosts[i],
			NodeCount:   stat.NodeCount,
			MinLatency:  stat.MinLatency,
			AvgLatency:  stat.AvgLatency,
			MaxLatency:  stat.MaxLatency,
			Connections: stat.Connections,
			Mode:        stat.Mode.String(),
			Version:     stat.Version,
			Error:       stat.Error,
		}
		result[i] = tempStat
	}
	zkStat.JSON(result)
}
