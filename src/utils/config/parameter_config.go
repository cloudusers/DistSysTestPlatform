package config

/**
* para config
 */
type ParaConfig struct {
	Duration string `json:"duration"`
	Interval string `json:"interval"`

	Ethernet EthernetConfig `json:"ethernet"`
	Disk     DiskConfig     `json:"disk"`
	Proc     ProcConfig     `json:"proc"`
	SysCpu   SysCpuConfig   `json:"syscpu"`
	SysMem   SysMemConfig   `json:"sysmem"`
}

/**
* ethernet config
 */
type EthernetConfig struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Parameter struct {
		Targetip          string `json:"targetip"`
		Duration          string `json:"duration"`
		Latency           string `json:"latency"`
		LatencyDuration   string `json:"latency_duration"`
		Loss              string `json:"loss"`
		LossDuration      string `json:"loss_duration"`
		Duplicate         string `json:"duplicate"`
		DuplicateDuration string `json:"duplicate_duration"`
		Reorder           string `json:"reorder"`
		ReorderDuration   string `json:"reorder_duration"`
	} `json:"parameter"`
}

/**
* disk config
 */
type DiskConfig struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Parameter struct {
		Dir      string `json:"dir"`
		Duration string `json:"duration"`
	} `json:"parameter"`
}

/**
* proc config
 */
type ProcConfig struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Parameter struct {
		Signal   string `json:"signal"`
		Duration string `json:"duration"`
	} `json:"parameter"`
}

/**
* system cpu config
 */
type SysCpuConfig struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Parameter struct {
		Duration string `json:"duration"`
	} `json:"parameter"`
}

/**
* system mem config
 */
type SysMemConfig struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Parameter struct {
		Duration string `json:"duration"`
	} `json:"parameter"`
}
