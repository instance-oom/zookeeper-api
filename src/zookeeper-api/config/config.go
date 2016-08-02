package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// Config info
var (
	APPVersion string
	ListenPort int
	LogLevel   int
	ZKHosts    []string
	ENV        = os.Getenv("ZA_ENV")
)

// LoadConfig - Load config info from .ini file
func LoadConfig() {
	APPVersion = beego.AppConfig.String("version")

	if listenPort := os.Getenv("ZA_LISTEN_PORT"); listenPort != "" {
		ListenPort, _ = strconv.Atoi(listenPort)
	} else {
		ListenPort = beego.BConfig.Listen.HTTPPort
	}

	hosts := os.Getenv("ZA_ZK_HOSTS")
	if hosts == "" {
		hosts = beego.AppConfig.String("zk_hosts")

	}
	ZKHosts = strings.Split(hosts, ",")
	if len(ZKHosts) == 0 {
		fmt.Println("Fail to load zk hosts from config file.")
		os.Exit(2)
	}

	if ENV == "" {
		ENV = "gdev"
	} else {
		ENV = strings.ToLower(ENV)
	}

	tmpLogLevel := os.Getenv("ZA_LOG_LEVEL")
	if tmpLogLevel == "" {
		LogLevel = beego.AppConfig.DefaultInt("log_level", 3)
	} else {
		var err error
		LogLevel, err = strconv.Atoi(tmpLogLevel)
		if err != nil {
			LogLevel = 3
		}
	}
}

// GetConfig - Return config info with map
func GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"Version":    APPVersion,
		"ENV":        ENV,
		"ListenPort": ListenPort,
		"ZKHosts":    ZKHosts,
		"LogLevel":   LogLevel,
	}
}
