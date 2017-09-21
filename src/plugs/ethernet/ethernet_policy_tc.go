package ethernet

import (
	"fmt"
	"net"
	//	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"

	tc "DistSysTestPlatform/src/plugs/ethernet/throttler"
	config "DistSysTestPlatform/src/utils/config"
	queue "DistSysTestPlatform/src/utils/queue"
)

func PolicyTc(item_ interface{}) {
	item := item_.(*(queue.BasicTimeoutQueueItem))
	value := item.Value()

	duration := item.Duration()
	var latency string
	var loss string
	var duplicate string
	var reorder string
	var targetip string

	switch value.(type) {
	case config.EthernetConfig:
		_struct := value.(config.EthernetConfig)
		p := _struct.Parameter

		latency = p.Latency
		loss = p.Loss
		duplicate = p.Duplicate
		reorder = p.Reorder

		targetip = p.Targetip
		break
	default:
		break
	}

	//
	//var ip_list string = config.DEFAULT_IPADDR
	var ip_list string = targetip
	targetIPv4, targetIPv6 := parseAddrs(ip_list) //depend host list.....

	cfg := tc.Config{
		Device:           config.DEFAULT_ETHERNET, //rand eth0, eth1
		Stop:             false,
		Latency:          parseLatency(latency),
		TargetBandwidth:  -1,
		DefaultBandwidth: -1,
		Reorder:          parseReorder(reorder),
		Duplicate:        parseDuplicate(duplicate),
		PacketLoss:       parseLoss(loss),
		TargetIps:        targetIPv4,
		TargetIps6:       targetIPv6,
		TargetPorts:      parsePorts(""),
		TargetProtos:     parseProtos("tcp,udp,icmp"),
		DryRun:           false,
	}
	go tc.Run(&cfg)

	timer1 := time.NewTicker(time.Duration(duration) * time.Second)
	<-timer1.C
	timer1.Stop()

	go tc.Stop(&cfg)
}

func parseLatency(delay string) int {
	val := delay
	l, err := strconv.Atoi(val)
	if err != nil {
		log.Error(fmt.Sprintf("Incorrectly specified packet latency:%s", delay))
		//os.Exit(1)
	}
	return l
}

func parseReorder(order string) float64 {
	val := order
	if strings.Contains(order, "%") {
		val = order[:len(order)-1]
	}
	l, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Error(fmt.Sprintf("Incorrectly specified packet reorder:%s", order))
		//os.Exit(1)
	}
	return l
}

func parseDuplicate(dup string) float64 {
	val := dup
	if strings.Contains(dup, "%") {
		val = dup[:len(dup)-1]
	}
	l, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Error(fmt.Sprintf("Incorrectly specified packet duplicate:%s", dup))
		//os.Exit(1)
	}
	return l
}

func parseLoss(loss string) float64 {
	val := loss
	if strings.Contains(loss, "%") {
		val = loss[:len(loss)-1]
	}
	l, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Error(fmt.Sprintf("Incorrectly specified packet loss:%s", loss))
		//os.Exit(1)
	}
	return l
}

func parseAddrs(addrs string) ([]string, []string) {
	adrs := strings.Split(addrs, ",")
	parsedIPv4 := []string{}
	parsedIPv6 := []string{}

	if addrs != "" {
		for _, adr := range adrs {
			ip := net.ParseIP(adr)
			if ip != nil {
				if ip.To4() != nil {
					parsedIPv4 = append(parsedIPv4, adr)
				} else {
					parsedIPv6 = append(parsedIPv6, adr)
				}
			} else { //Not a valid single IP, could it be a CIDR?
				parsedIP, net, err := net.ParseCIDR(adr)
				if err == nil {
					if parsedIP.To4() != nil {
						parsedIPv4 = append(parsedIPv4, net.String())
					} else {
						parsedIPv6 = append(parsedIPv6, net.String())
					}
				} else {
					log.Error(fmt.Sprintf("Incorrectly specified target IP or CIDR:%s", adr))
					//os.Exit(1)
				}
			}
		}
	}

	return parsedIPv4, parsedIPv6
}

func parsePorts(ports string) []string {
	prts := strings.Split(ports, ",")
	parsed := []string{}

	if ports != "" {
		for _, prt := range prts {
			if strings.Contains(prt, ":") {
				if validRange(prt) {
					parsed = append(parsed, prt)
				} else {
					log.Error(fmt.Sprintf("Incorrectly specified port range:%s", prt))
					//os.Exit(1)
				}
			} else { //Isn't a range, check if just a single port
				if validPort(prt) {
					parsed = append(parsed, prt)
				} else {
					log.Error(fmt.Sprintf("Incorrectly specified port:%s", prt))
					//os.Exit(1)
				}
			}
		}
	}

	return parsed
}

func parsePort(port string) int {
	prt, err := strconv.Atoi(port)
	if err != nil {
		return 0
	}

	return prt
}

func validPort(port string) bool {
	prt := parsePort(port)
	return prt > 0 && prt < 65536
}

func validRange(ports string) bool {
	pr := strings.Split(ports, ":")

	if len(pr) == 2 {
		if !validPort(pr[0]) || !validPort(pr[1]) {
			return false
		}

		if portHigher(pr[0], pr[1]) {
			return false
		}
	} else {
		return false
	}

	return true
}

func portHigher(prt1, prt2 string) bool {
	p1 := parsePort(prt1)
	p2 := parsePort(prt2)

	return p1 > p2
}

func parseProtos(protos string) []string {
	ptcs := strings.Split(protos, ",")
	parsed := []string{}

	if protos != "" {
		for _, ptc := range ptcs {
			p := strings.ToLower(ptc)
			if p == "udp" ||
				p == "tcp" ||
				p == "icmp" {
				parsed = append(parsed, p)
			} else {
				log.Error(fmt.Sprintf("Incorrectly specified protocol:%s", p))
				//os.Exit(1)
			}
		}
	}

	return parsed
}
