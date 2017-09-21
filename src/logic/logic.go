package logic

import (
	"fmt"

	cmd "DistSysTestPlatform/src/utils/cmd"
	config "DistSysTestPlatform/src/utils/config"
	log "github.com/cihub/seelog"
	"github.com/manucorporat/try"
)

func CheckValidity(tc config.TypeClass,
	cfg config.ParaConfig) bool {

	var ret bool = false

	switch tc {
	case config.TYPE_ETHERNET:
		if cfg.Ethernet.Type == "yes" {
			ret = true
		}
		break
	case config.TYPE_DISK:
		if cfg.Disk.Type == "yes" {
			ret = true
		}
		break
	case config.TYPE_PROC:
		if cfg.Proc.Type == "yes" {
			ret = true
		}
		break
	case config.TYPE_SYSCPU:
		if cfg.SysCpu.Type == "yes" {
			ret = true
		}
		break
	case config.TYPE_SYSMEM:
		if cfg.SysMem.Type == "yes" {
			ret = true
		}
		break
	default:
		break
	}

	return ret
}

func TearDown(cfg config.ParaConfig) {
	try.This(func() {
		if CheckValidity(config.TYPE_ETHERNET, cfg) {
			qdis_cmd := "tc qdisc del dev eth0 handle 10: root"
			iptable_cmd := "iptables -F"
			cmd.CommandFactory.RunCmd(qdis_cmd)
			cmd.CommandFactory.RunCmd(iptable_cmd)
		}

		if CheckValidity(config.TYPE_DISK, cfg) {
			disk_cmd := fmt.Sprintf("umount -l %s", cfg.Disk.Parameter.Dir)
			cmd.CommandFactory.RunCmd(disk_cmd)
		}

	}).Finally(func() {

	}).Catch(func(e try.E) {
		log.Error(e)
	})
}
