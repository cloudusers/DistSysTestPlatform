package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
	mux "github.com/gorilla/mux"

	logic "DistSysTestPlatform/src/logic"
	ethernet "DistSysTestPlatform/src/plugs/ethernet"
	fs "DistSysTestPlatform/src/plugs/fs"
	proc "DistSysTestPlatform/src/plugs/proc"
	sys "DistSysTestPlatform/src/plugs/sys"
	config "DistSysTestPlatform/src/utils/config"
	convert "DistSysTestPlatform/src/utils/convert"
	queue "DistSysTestPlatform/src/utils/queue"
	request "DistSysTestPlatform/src/utils/requests"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		io.WriteString(w, config.FAIL)
		return
	}

	trans_body := string(body)
	cfg, err := config.SingletonParaConfig.Parse(trans_body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}

	log.Debug("(-------------tasker para conf)===================================", cfg)

	//start abnormal simulate threads
	//setup timer wait timeout(arrival duration)
	total_duration := convert.String2Int64(cfg.Duration)
	interval := convert.String2Int64(cfg.Interval)
	if total_duration < interval {
		x := total_duration
		total_duration = interval
		interval = x
	}

	//resume ethernet|disk
	logic.TearDown(cfg)

	io.WriteString(w, config.OK)
	/////////begin
	t_start := time.Now().Unix()
	timer := time.NewTicker(time.Duration(interval) * time.Second)
	is_exist_hack := false
	is_error_hack := false
	for {
		//ethernet
		if logic.CheckValidity(config.TYPE_ETHERNET, cfg) {
			is_exist_hack = true
			if !Dispatcher(cfg, config.SCHEDULER_STATUS_FINISH_ETHERNET,
				"dispatch_ethernet", trans_body) {
				is_error_hack = true
				break
			}

			//timer
			<-timer.C
			if IsTimeout(t_start, total_duration-interval) {
				timer.Stop()
				break
			}
		}

		//disk
		if logic.CheckValidity(config.TYPE_DISK, cfg) {
			is_exist_hack = true
			if !Dispatcher(cfg, config.SCHEDULER_STATUS_FINISH_DISK,
				"dispatch_disk", trans_body) {
				is_error_hack = true
				break
			}

			//timer
			<-timer.C
			if IsTimeout(t_start, total_duration-interval) {
				timer.Stop()
				break
			}
		}

		//proc
		if logic.CheckValidity(config.TYPE_PROC, cfg) {
			is_exist_hack = true
			if !Dispatcher(cfg, config.SCHEDULER_STATUS_FINISH_PROC,
				"dispatch_proc", trans_body) {
				is_error_hack = true
				break
			}

			//timer
			<-timer.C
			if IsTimeout(t_start, total_duration-interval) {
				timer.Stop()
				break
			}
		}

		//syscpu
		if logic.CheckValidity(config.TYPE_SYSCPU, cfg) {
			is_exist_hack = true
			if !Dispatcher(cfg, config.SCHEDULER_STATUS_FINISH_SYSCPU,
				"dispatch_syscpu", trans_body) {
				is_error_hack = true
				break
			}

			//timer
			<-timer.C
			if IsTimeout(t_start, total_duration-interval) {
				timer.Stop()
				break
			}
		}

		//sysmem
		if logic.CheckValidity(config.TYPE_SYSMEM, cfg) {
			is_exist_hack = true
			if !Dispatcher(cfg, config.SCHEDULER_STATUS_FINISH_SYSMEM,
				"dispatch_sysmem", trans_body) {
				is_error_hack = true
				break
			}

			//timer
			<-timer.C
			if IsTimeout(t_start, total_duration-interval) {
				timer.Stop()
				break
			}
		}

		if !is_exist_hack {
			break
		}
	}

	//resume ethernet|disk again
	logic.TearDown(cfg)

	if is_error_hack {
		io.WriteString(w, config.FAIL)
		return
	}
	/////////////end

	log.Debug("task finish---------------", trans_body)
}

func IsTimeout(t int64, thresh int64) bool {
	t_stop := time.Now().Unix()
	if t_stop-t > thresh {
		return true
	}

	return false
}

func Dispatcher(cfg config.ParaConfig, status string,
	dispatch string, body string) bool {

	go func() {
		request.PostRequest("127.0.0.1",
			config.SingletonXmlConfig.Base.AgentListen,
			dispatch,
			body)
	}()

	return true
}

func DispatchEthernetHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}

	trans_body := string(body)
	log.Debug("(------------ethernet input body)===================================", trans_body)
	cfg, err := config.SingletonParaConfig.Parse(trans_body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}
	log.Debug("(-------------ethernet para conf)===================================", cfg)

	duration := convert.String2Int64(cfg.Ethernet.Parameter.Duration)
	err = queue.QueueFactory.NetEnqueue(cfg.Ethernet, duration)
	if nil != err {
		io.WriteString(w, config.FAIL)
	} else {
		io.WriteString(w, config.OK)
	}
}

func DispatchDiskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}

	trans_body := string(body)
	log.Debug("(------------disk input body)===================================", trans_body)
	cfg, err := config.SingletonParaConfig.Parse(trans_body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}
	log.Debug("(-------------disk para conf)===================================", cfg)

	duration := convert.String2Int64(cfg.Disk.Parameter.Duration)
	err = queue.QueueFactory.DiskEnqueue(cfg.Disk, duration)
	if nil != err {
		io.WriteString(w, config.FAIL)
	} else {
		io.WriteString(w, config.OK)
	}
}

func DispatchSysCpuHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}

	trans_body := string(body)
	log.Debug("(------------syscpu input body)===================================", trans_body)
	cfg, err := config.SingletonParaConfig.Parse(trans_body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}
	log.Debug("(-------------syscpu para conf)===================================", cfg)

	duration := convert.String2Int64(cfg.SysCpu.Parameter.Duration)
	err = queue.QueueFactory.SysCpuEnqueue(cfg.SysCpu, duration)
	if nil != err {
		io.WriteString(w, config.FAIL)
	} else {
		io.WriteString(w, config.OK)
	}
}

func DispatchSysMemHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}

	trans_body := string(body)
	log.Debug("(------------sysmem input body)===================================", trans_body)
	cfg, err := config.SingletonParaConfig.Parse(trans_body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}
	log.Debug("(-------------sysmem para conf)===================================", cfg)

	duration := convert.String2Int64(cfg.SysMem.Parameter.Duration)
	err = queue.QueueFactory.SysMemEnqueue(cfg.SysMem, duration)
	if nil != err {
		io.WriteString(w, config.FAIL)
	} else {
		io.WriteString(w, config.OK)
	}
}

func DispatchProcHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}

	trans_body := string(body)
	log.Debug("(------------proc input body)===================================", trans_body)
	cfg, err := config.SingletonParaConfig.Parse(trans_body)
	if err != nil {
		io.WriteString(w, config.FAIL)
		return
	}
	log.Debug("(-------------proc para conf)===================================", cfg)

	duration := convert.String2Int64(cfg.Proc.Parameter.Duration)
	err = queue.QueueFactory.ProcEnqueue(cfg.Proc, duration)
	if nil != err {
		io.WriteString(w, config.FAIL)
	} else {
		io.WriteString(w, config.OK)
	}
}

func StartInspector() {
	go ethernet.RunTc()
	go fs.Run()
	go sys.RunCpu()
	go sys.RunMem()
	go proc.Run()
}

func StartHttp() {
	//start httpserver
	r := mux.NewRouter()
	r.HandleFunc("/task", TaskHandler)
	r.HandleFunc("/dispatch_ethernet", DispatchEthernetHandler)
	r.HandleFunc("/dispatch_disk", DispatchDiskHandler)
	r.HandleFunc("/dispatch_syscpu", DispatchSysCpuHandler)
	r.HandleFunc("/dispatch_sysmem", DispatchSysMemHandler)
	r.HandleFunc("/dispatch_proc", DispatchProcHandler)
	//...
	http.Handle("/", r)

	listen := ":" + config.SingletonXmlConfig.Base.Listen
	log.Critical(http.ListenAndServe(listen, r))
	log.Info("Http Server Stoped(see Critical Log)")
}

func Start() {
	//start inspectors
	go func() {
		StartInspector()
	}()

	//start http
	StartHttp()
}

/*
* main
 */
func main() {
	//config
	config.SingletonXmlConfig.Parse("config.agent.xml")

	//logger
	logger, err := log.LoggerFromConfigAsFile(config.SingletonXmlConfig.Base.Logpath)
	if err != nil {
		fmt.Println("init logger fail!")
		return
	}
	log.ReplaceLogger(logger)
	log.Flush()

	//start inspector and httpserver
	Start()
}
