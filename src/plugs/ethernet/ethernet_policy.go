package ethernet

//
//import (
//	"fmt"
//	"math/rand"
//	//"reflect"
//	"time"
//
//	cmd "DistSysTestPlatform/src/utils/cmd"
//	config "DistSysTestPlatform/src/utils/config"
//	convert "DistSysTestPlatform/src/utils/convert"
//	queue "DistSysTestPlatform/src/utils/queue"
//	netfilter "github.com/AkihiroSuda/go-netfilter-queue"
//	//log "github.com/cihub/seelog"
//)
//
//func Policy(item_ interface{}) netfilter.Verdict {
//	ret := netfilter.NF_ACCEPT
//
//	item := item_.(*(queue.BasicTimeoutQueueItem))
//	value := item.Value()
//	enqueue_time := item.EnqueuedTime()
//	duration := item.Duration()
//	/*
//		fmt.Println(reflect.TypeOf(item_)) //*queue.BasicTimeoutQueueItem
//		fmt.Println("VALUE:", value) //(this *config.HookNetPacket)
//		fmt.Println("ENTIME:", enqueue_time)
//		fmt.Println("DURATION:", duration)
//		fmt.Println(reflect.TypeOf(value))
//		fmt.Println(reflect.TypeOf(enqueue_time))
//		fmt.Println(reflect.TypeOf(duration))
//		now := time.Now()
//		fmt.Println("NOW:", now)
//	*/
//	time_diff := time.Since(enqueue_time).Seconds()
//	//fmt.Println("DIFF:", time_diff)
//	if int(time_diff) > duration {
//		return ret
//	}
//
//	//case config.EthernetConfig:
//	switch value.(type) {
//	case config.EthernetConfig:
//		_struct := value.(config.EthernetConfig)
//		//fmt.Println(reflect.TypeOf(_struct))
//		p := _struct.Parameter
//		latency := convert.String2Int(p.Latency)
//		loss := convert.String2Float64(p.Loss)
//		duplicate := convert.String2Float64(p.Duplicate)
//		unorder := convert.String2Float64(p.Reorder)
//
//		//fmt.Println(p)
//		if 0 != latency {
//			time.Sleep((time.Duration)(latency) * time.Millisecond)
//		}
//		if 0 != loss {
//			if r := rand.Int63n(100000); r < int64(loss*10000) {
//				ret = netfilter.NF_DROP
//			}
//		}
//		if 0 != duplicate {
//			if r := rand.Int63n(100000); r < int64(duplicate*10000) {
//				ret = netfilter.NF_REPEAT
//			}
//		}
//		if 0 != unorder {
//			if r := rand.Int63n(100000); r < int64(unorder*10000) {
//				ret = netfilter.NF_QUEUE
//			}
//		}
//
//		if probab(30) {
//			nfq_cmd := fmt.Sprintf("ifdown eth0;ifdown eth1;ifup eth0;ifup eth1")
//			res := cmd.CommandFactory.RunCmd(nfq_cmd)
//			if !res.GetStatus() {
//				//log.Critical(res.GetMsg())
//			}
//		}
//
//		//again enqueue
//		queue.QueueFactory.NetQueue.UpdateEnqueue(item, enqueue_time, duration)
//		break
//	default:
//		break
//	}
//
//	return ret
//}
//
//func probab(percentage int) bool {
//	return rand.Intn(99) < percentage
//}
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
