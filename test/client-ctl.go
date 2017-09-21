package main

import (
	"flag"
	"fmt"
	//"io/ioutil"
	//"os"
	"strings"

	request "DistSysTestPlatform/src/utils/requests"
)

func main() {
	//basic info
	host := flag.String("host", "127.0.0.1", "host")
	targetip := flag.String("targetip", "10.134.99.128", "ip address")
	duration := flag.String("duration", "30", "hack duration")
	interval := flag.String("interval", "10", "hack interval")

	//ethernet inject
	ethernet := flag.String("ethernet", "no", "ethernet hack")
	eth_duration := flag.String("eth_duration", "10", "ethernet duration")
	eth_latency := flag.String("eth_latency", "100", "ethernet latency")
	eth_duplicate := flag.String("eth_duplicate", "0", "ethernet duplicate")
	eth_loss := flag.String("eth_loss", "0", "ethernet loss")
	eth_reorder := flag.String("eth_reorder", "0", "ethernet reorder")
	//ignore [latency,duplicate,loss,reorder]_duration

	//disk inject
	disk := flag.String("disk", "no", "disk hack")
	disk_duration := flag.String("disk_duration", "10", "disk duration")
	diskdir := flag.String("diskdir", "/tmp", "disk hack dir")

	//proc inject
	proc := flag.String("proc", "no", "proc hack")
	proc_duration := flag.String("proc_duration", "10", "proc duration")
	proc_signal := flag.String("proc_signal", "15", "proc signal")

	//syscpu inject
	syscpu := flag.String("syscpu", "no", "syscpu hack")
	syscpu_duration := flag.String("syscpu_duration", "10", "syscpu duration")

	//sysmem inject
	sysmem := flag.String("sysmem", "no", "sysmem hack")
	sysmem_duration := flag.String("sysmem_duration", "10", "sysmem duration")

	flag.Parse()
	/*
		if len(os.Args) < 5 {
			fmt.Printf("%s -help\n", os.Args[0])
			return
		}
	*/
	str := `{
			"duration":"` + *duration + `" ,
			"interval":"` + *interval + `" ,
			"ethernet":{
				"type":"` + *ethernet + `" ,
				"status":"success",
				"parameter":{
					"duration":"` + *eth_duration + `",
					"latency":"` + *eth_latency + `" ,
					"duplicate":"` + *eth_duplicate + `" ,
					"loss":"` + *eth_loss + `" ,
					"reorder":"` + *eth_reorder + `" ,
					"targetip":"` + *targetip + `"
				}
			},
			"disk":{
				"type":"` + *disk + `" ,
				"status":"success",
				"parameter":{
					"duration":"` + *disk_duration + `",
					"dir":"` + *diskdir + `"
				}
			},
			"proc":{
				"type":"` + *proc + `" ,
				"status":"success" ,
				"parameter":{
					"signal":"` + *proc_signal + `",
					"duration":"` + *proc_duration + `"
				}
			},
			"syscpu":{
				"type":"` + *syscpu + `" ,
				"status":"success",
				"parameter":{
					"duration":"` + *syscpu_duration + `"
				}
			},
			"sysmem":{
				"type":"` + *sysmem + `" ,
				"status":"success",
				"parameter":{
					"duration":"` + *sysmem_duration + `"
				}
			}
		}`

	req_host := fmt.Sprintf("http://%s:8091/task", *host)

	req, err := request.NewRequest("POST", req_host, strings.NewReader(str))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Accept-Encoding", "")

	//resp, err := req.Do()
	_, err = req.Do()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		if "ok" == string(body) {
			fmt.Println("abnormal request has been send to agent\n")
		}
	*/
	fmt.Println("abnormal request has been send to agent\n")
}
