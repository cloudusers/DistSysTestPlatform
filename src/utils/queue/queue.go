package queue

import (
	//"fmt"
	//"time"
	config "DistSysTestPlatform/src/utils/config"
)

type TimeoutQueue struct {
	NetQueue   TQueue
	NetQueueCh chan TQueueItem

	DiskQueue   TQueue
	DiskQueueCh chan TQueueItem

	SysCpuQueue   TQueue
	SysCpuQueueCh chan TQueueItem

	SysMemQueue   TQueue
	SysMemQueueCh chan TQueueItem

	ProcQueue   TQueue
	ProcQueueCh chan TQueueItem
}

func (this *TimeoutQueue) Initialize() bool {

	this.NetQueue = NewBasicTimeoutQueue()
	this.NetQueueCh = this.NetQueue.GetDequeueChan()

	this.DiskQueue = NewBasicTimeoutQueue()
	this.DiskQueueCh = this.DiskQueue.GetDequeueChan()

	this.SysCpuQueue = NewBasicTimeoutQueue()
	this.SysCpuQueueCh = this.SysCpuQueue.GetDequeueChan()

	this.SysMemQueue = NewBasicTimeoutQueue()
	this.SysMemQueueCh = this.SysMemQueue.GetDequeueChan()

	this.ProcQueue = NewBasicTimeoutQueue()
	this.ProcQueueCh = this.ProcQueue.GetDequeueChan()

	return true
}

func (this *TimeoutQueue) NetEnqueue(
	element_json_config config.EthernetConfig,
	duration int64 /*second*/) error {
	//fmt.Println(element_json_config)

	item, err := NewBasicTimeoutQueueItem(element_json_config, int(duration))
	if err != nil {
		return err
	}

	err = this.NetQueue.Enqueue(item)
	if err != nil {
		return err
	}

	return nil
}

func (this *TimeoutQueue) DiskEnqueue(
	element_json_config config.DiskConfig,
	duration int64 /*second*/) error {
	//fmt.Println(element_json_config)

	item, err := NewBasicTimeoutQueueItem(element_json_config, int(duration))
	if err != nil {
		return err
	}

	err = this.DiskQueue.Enqueue(item)
	if err != nil {
		return err
	}

	return nil
}

func (this *TimeoutQueue) SysCpuEnqueue(
	element_json_config config.SysCpuConfig,
	duration int64 /*second*/) error {

	item, err := NewBasicTimeoutQueueItem(element_json_config, int(duration))
	if err != nil {
		return err
	}

	err = this.SysCpuQueue.Enqueue(item)
	if err != nil {
		return err
	}

	return nil
}

func (this *TimeoutQueue) SysMemEnqueue(
	element_json_config config.SysMemConfig,
	duration int64 /*second*/) error {

	item, err := NewBasicTimeoutQueueItem(element_json_config, int(duration))
	if err != nil {
		return err
	}

	err = this.SysMemQueue.Enqueue(item)
	if err != nil {
		return err
	}

	return nil
}

func (this *TimeoutQueue) ProcEnqueue(
	element_json_config config.ProcConfig,
	duration int64 /*second*/) error {
	//fmt.Println(element_json_config)

	item, err := NewBasicTimeoutQueueItem(element_json_config, int(duration))
	if err != nil {
		return err
	}

	err = this.ProcQueue.Enqueue(item)
	if err != nil {
		return err
	}

	return nil
}

func NewQueueSingleton() *TimeoutQueue {
	return &TimeoutQueue{}
}

var QueueFactory *TimeoutQueue

func init() {
	QueueFactory = NewQueueSingleton()
	QueueFactory.Initialize()
}
