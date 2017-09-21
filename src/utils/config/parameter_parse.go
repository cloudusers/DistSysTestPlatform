package config

import (
	"encoding/json"
	"sync"
	"unsafe"

	log "github.com/cihub/seelog"
)

func UnmarshalJson(jstr string) (map[string]interface{}, error) {

	var data map[string]interface{}

	err := json.Unmarshal([]byte(jstr), &data)
	if err != nil {
		log.Error(err.Error())
		//panic("config.UnmarshalJson")
	}

	return data, nil
}

func MarshalJson(kv map[string]interface{}) (string, error) {

	b, err := json.Marshal(kv)
	if err != nil {
		log.Error(err.Error())
		//panic("config.MarshalJson")
		var xx string
		return xx, err
	}

	return *(*string)(unsafe.Pointer(&b)), nil
}

func MarshalsJson(kv map[string]map[string][]string) (string, error) {

	b, err := json.Marshal(kv)
	if err != nil {
		log.Error(err.Error())
		var xx string
		return xx, err
	}

	return string(b), nil
}

func Json2Array(jstr string) ([]string, error) {
	var data []string

	err := json.Unmarshal([]byte(jstr), &data)
	if err != nil {
		log.Error(err.Error())
		//panic("config.Json2Array")
	}

	return data, nil
}

func Array2Json(data []string) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		//panic("config.Array2Json")
	}

	return string(b), nil
}

///////////////////////////////////////////////////

//struct to jsonstr
func Struct2Json(cfg ParaConfig) (string, error) {
	b, err := json.Marshal(&cfg)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return string(b), nil
}

//jsonstr to struct
func Json2Struct(jstr string) (ParaConfig, error) {
	var cfg ParaConfig
	cfg = ParaConfig{}

	err := json.Unmarshal([]byte(jstr), &cfg)
	if err != nil {
		log.Error(err.Error())
		return ParaConfig{}, err
	}

	return cfg, nil
}

//
func (this *ParaConfig) Parse(json_str string) (ParaConfig, error) {

	mutex.Lock()
	defer mutex.Unlock()

	err := json.Unmarshal([]byte(json_str), this)
	if err != nil {
		log.Error(err.Error())
		return ParaConfig{}, err
	}

	var cfg ParaConfig
	cfg = *this

	return cfg, nil
}

//new config
func NewParaConfigSingleton() *ParaConfig {
	return &ParaConfig{}
}

var SingletonParaConfig = NewParaConfigSingleton()

var mutex sync.Mutex
