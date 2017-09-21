package ethernet

//
//import (
//	"errors"
//	"fmt"
//
//	cmd "DistSysTestPlatform/src/utils/cmd"
//	config "DistSysTestPlatform/src/utils/config"
//	netfilter "github.com/AkihiroSuda/go-netfilter-queue"
//	log "github.com/cihub/seelog"
//)
//
//func Run() {
//	log.Info("Ethernet Inspector Start")
//
//	var num uint16 = 1
//	for {
//		num += 1
//		if num >= 65535 {
//			log.Critical("Ethernet Inspector FAIL, queue number exceed 65535")
//			return
//		}
//
//		insp := &NFQInspector{
//			NFQNumber:        num,
//			EnableTCPWatcher: true,
//		}
//
//		SetNfq(insp)
//		defer UnSetNfq(insp)
//
//		err := insp.Serve()
//		if err != nil {
//			log.Critical("Ethernet Inspector FAIL, err:", err)
//			return
//		}
//
//		nfq, err := netfilter.NewNFQueue(insp.NFQNumber, 1000000, netfilter.NF_DEFAULT_PACKET_SIZE)
//		if err != nil {
//			log.Critical("Ethernet Inspector FAIL, err:", err)
//			continue
//		}
//
//		defer nfq.Close()
//		nfpChan := nfq.GetPackets()
//		for {
//			nfp := <-nfpChan
//			ip, tcp := insp.decodeNFPacket(nfp)
//			/*
//				if insp.EnableTCPWatcher && insp.tcpWatcher.IsTCPRetrans(ip, tcp) {
//					//nfp.SetVerdict(netfilter.NF_DROP)
//					nfp.SetVerdict(netfilter.NF_ACCEPT)
//					continue
//				}
//			*/
//
//			go func() {
//				if err := insp.onPacket(nfp, ip, tcp); err != nil {
//					log.Error(err)
//				}
//			}()
//		}
//	}
//
//	log.Info("Ethernet Inspector Stop")
//}
//
//func SetNfq(insp *NFQInspector) error {
//	log.Info("Ethernet NFQ Setting.......")
//
//	nfq_cmd := fmt.Sprintf("iptables -F")
//	res := cmd.CommandFactory.RunCmd(nfq_cmd)
//
//	nfq_cmd = fmt.Sprintf("iptables -A OUTPUT -p tcp -m owner --uid-owner $(id -u %s) -j NFQUEUE --queue-num %d", config.SingletonXmlConfig.Base.UserGroup, insp.NFQNumber)
//	res = cmd.CommandFactory.RunCmd(nfq_cmd)
//	if !res.GetStatus() {
//		log.Info("Ethernet Inspector NFQ Failed")
//		log.Critical(res.GetMsg())
//		return errors.New("Ethernet Inspector NFQ Failed")
//	}
//	log.Info("Ethernet NFQ set success")
//
//	return nil
//}
//
//func UnSetNfq(insp *NFQInspector) error {
//	log.Info("Ethernet NFQ UnSetting.......")
//
//	nfq_cmd := fmt.Sprintf("iptables -D OUTPUT -p tcp -m owner --uid-owner $(id -u %s) -j NFQUEUE --queue-num %d", config.SingletonXmlConfig.Base.UserGroup, insp.NFQNumber)
//	res := cmd.CommandFactory.RunCmd(nfq_cmd)
//	if !res.GetStatus() {
//		log.Info("Ethernet Inspector NFQ Failed")
//		log.Critical(res.GetMsg())
//		return errors.New("Ethernet Inspector NFQ Failed")
//	}
//	log.Info("Ethernet NFQ resume success")
//
//	return nil
//}
