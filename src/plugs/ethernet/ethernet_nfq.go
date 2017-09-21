package ethernet

//
//import (
//	//"fmt"
//
//	tcpwatcher "DistSysTestPlatform/src/plugs/ethernet/tcpwatcher"
//	transceiver "DistSysTestPlatform/src/plugs/transceiver"
//	queue "DistSysTestPlatform/src/utils/queue"
//	netfilter "github.com/AkihiroSuda/go-netfilter-queue"
//	log "github.com/cihub/seelog"
//	layers "github.com/google/gopacket/layers"
//)
//
//type NFQInspector struct {
//	NFQNumber        uint16
//	EnableTCPWatcher bool
//	trans            transceiver.Transceiver
//	tcpWatcher       *tcpwatcher.TCPWatcher
//}
//
//func (this *NFQInspector) Serve() error {
//	log.Info("Initializing Ethernet Inspector %#v", this)
//	var err error
//
//	if this.EnableTCPWatcher {
//		this.tcpWatcher = tcpwatcher.New()
//	}
//
//	this.trans, err = transceiver.NewTransceiver("local://", "Biztech.DSTP.Agent")
//	if err != nil {
//		return err
//	}
//	this.trans.Start()
//
//	return nil
//}
//
//func (this *NFQInspector) decodeNFPacket(nfp netfilter.NFPacket) (ip *layers.IPv4, tcp *layers.TCP) {
//	if layer := nfp.Packet.Layer(layers.LayerTypeIPv4); layer != nil {
//		ip, _ = layer.(*layers.IPv4)
//	}
//	if layer := nfp.Packet.Layer(layers.LayerTypeTCP); layer != nil {
//		tcp, _ = layer.(*layers.TCP)
//	}
//	if layer := nfp.Packet.Layer(layers.LayerTypeTCP); layer != nil {
//		tcp, _ = layer.(*layers.TCP)
//	}
//
//	return
//}
//
///*
//* ethernet header 0xff.... 0x00....
// */
//func packetBytes(nfp netfilter.NFPacket) []byte {
//	dummyEth := []byte("\xff\xff\xff\xff\xff\xff" +
//		"\x00\x00\x00\x00\x00\x00" +
//		"\x08\x00")
//	payload := nfp.Packet.Data()
//
//	return append(dummyEth[:], payload[:]...)
//}
//
//func (this *NFQInspector) onPacket(nfp netfilter.NFPacket,
//	ip *layers.IPv4,
//	tcp *layers.TCP) error {
//
//	//bytes := packetBytes(nfp)
//
//	//fmt.Println("packet action comming in BYTE:", string(bytes))
//	select {
//	case deq := <-queue.QueueFactory.NetQueueCh:
//		policy := Policy(deq)
//		nfp.SetVerdict(policy)
//	default:
//		nfp.SetVerdict(netfilter.NF_ACCEPT)
//	}
//
//	return nil
//}
