package config

import (
	//"fmt"
	"encoding/xml"
	"io/ioutil"
)

type XmlConfig struct {
	Base Base_t `xml:"base"`
}

type Base_t struct {
	//Name  string `xml:"name,attr"`
	UserGroup   string `xml:"usergroup"`
	Listen      string `xml:"listen"`
	ScheListen  string `xml:"sche_listen"`
	AgentListen string `xml:"agent_listen"`
	Logpath     string `xml:"log"`
}

type Init_t struct {
	Hostname string `xml:"hostname"`
}

type Verify_t struct {
	Hostname string `xml:"hostname"`
}

func (this *XmlConfig) Parse(config_path string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	content, err := ioutil.ReadFile(config_path)
	if err != nil {
		return false
	}

	err = xml.Unmarshal(content, &this)
	if err != nil {
		return false
	}
	//fmt.Println(this)

	return true
}

func NewXmlConfigSingleton() *XmlConfig {
	return &XmlConfig{}
}

var SingletonXmlConfig = NewXmlConfigSingleton()

//////// CONST ///////////
const (
	//request response status
	OK          = "ok"
	OKTOO       = `{"flag":1}`
	FAIL        = "fail"
	REPEAT_TASK = "duplicate"

	//scheduler status
	SCHEDULER_STATUS_CREATE            = "create"
	SCHEDULER_STATUS_RUNNING           = "running"
	SCHEDULER_STATUS_RUNNING_INTERRUPT = "running-interrupt"
	SCHEDULER_STATUS_RUNNING_ETHERNET  = "running-ethernet"
	SCHEDULER_STATUS_RUNNING_DISK      = "running-disk"
	SCHEDULER_STATUS_RUNNING_PROC      = "running-proc"
	SCHEDULER_STATUS_RUNNING_SYSCPU    = "running-sys-cpu"
	SCHEDULER_STATUS_RUNNING_SYSMEM    = "running-sys-mem"

	SCHEDULER_STATUS_FINISH           = "finish"
	SCHEDULER_STATUS_FINISH_INTERRUPT = "finish-interrupt"
	SCHEDULER_STATUS_FINISH_ETHERNET  = "finish-ethernet"
	SCHEDULER_STATUS_FINISH_DISK      = "finish-disk"
	SCHEDULER_STATUS_FINISH_PROC      = "finish-proc"
	SCHEDULER_STATUS_FINISH_SYSCPU    = "finish-sys-cpu"
	SCHEDULER_STATUS_FINISH_SYSMEM    = "finish-sys-mem"

	//task exec status
	TASK_STATUS_SUCCESS = "success"
	TASK_STATUS_FAIL    = "fail"
	TASK_STATUS_RUNNING = "running"

	//error step(error class)
	ERROR_STEP_NONE                 = "none"
	ERROR_STEP_PARSER               = "parser"
	ERROR_STEP_FLOW                 = "flow"
	ERROR_STEP_CASE                 = "autocase"
	ERROR_STEP_INIT                 = "init"
	ERROR_STEP_VERIFY               = "verify"
	ERROR_STEP_INTERRUPT            = "interrupt"
	ERROR_STEP_ETHERNET             = "ethernet"
	ERROR_STEP_DISK                 = "disk"
	ERROR_STEP_SYSCPU               = "syscpu"
	ERROR_STEP_SYSMEM               = "sysmem"
	ERROR_STEP_PROC                 = "proc"
	ERROR_STEP_CONNECTION_AGENT     = "connection_agent"
	ERROR_STEP_CONNECTION_SCHE      = "connection_sche"
	ERROR_STEP_CONNECTION_BIZTREE   = "connection_biztree"
	ERROR_STEP_CONNECTION_BIZCONF   = "connection_bizconf"
	ERROR_STEP_CONNECTION_BIZDEPLOY = "connection_bizdeploy"
)

//service type
const (
	SERVICE_TYPE_CPLUSPLUS = "cplusplus"
	SERVICE_TYPE_JAVA      = "java"
	SERVICE_TYPE_GOLANG    = "golang"
)

//
const (
	DEFAULT_ETHERNET  = "eth0"
	DEFAULT_IPADDR    = "10.134.99.128" //for testing //127.0.0.1
	DEFAULT_PARAMETER = `{
			"duration":"60" ,
			"interval":"10" ,
			"ethernet":{
				"type":"yes" ,
				"status":"success",
				"parameter":{
					"duration":"20",
					"latency":"2000" ,
					"duplicate":"0" ,
					"loss":"0" ,
					"reorder":"20"
					"latency_duration":"10" ,
					"duplicate_duration":"10" ,
					"loss_duration":"10" ,
					"reorder_duration":"10"
				}
			},
			"disk":{
				"type":"yes" ,
				"status":"success",
				"parameter":{
					"duration":"10",
					"dir":"/opt/logs"
				}
			},
			"proc":{
				"type":"yes" ,
				"status":"success" ,
				"parameter":{
					"signal":"15",
					"duration":"2"
				}
			},
			"syscpu":{
				"type":"yes" ,
				"status":"success",
				"parameter":{
					"duration":"5"
				}
			},
			"sysmem":{
				"type":"yes" ,
				"status":"success",
				"parameter":{
					"duration":"5"
				}
			}
		}`
)

type TypeClass int

const (
	_ TypeClass = iota
	TYPE_ETHERNET
	TYPE_DISK
	TYPE_PROC
	TYPE_SYSCPU
	TYPE_SYSMEM
)
