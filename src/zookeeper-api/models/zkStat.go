package models

// ZKStat - zookeeper server state
type ZKStat struct {
	Server      string `json:'Server'`
	NodeCount   int64  `json:'NodeCount'`
	MinLatency  int64  `json:'MinLatency'`
	AvgLatency  int64  `json:'AvgLatency'`
	MaxLatency  int64  `json:'MaxLatency'`
	Connections int64  `json:'Connections'`
	Mode        string `json:'Mode'`
	Version     string `json:'Version'`
	Error       error  `json:'Error'`
}
