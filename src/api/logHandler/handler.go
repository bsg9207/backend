// Created by Seunggwan, Back on 2024.04.18
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package logHandler

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gorani/common/log"

	"gopkg.in/yaml.v2"
)

type TLogSettings struct {
	Path     string `yaml:"path"`
	Debug    bool   `yaml:"debug"`
	Level    int32  `yaml:"level"`
	RotaCnt  int32  `yaml:"rotacnt"`
	RotaFreq int64  `yaml:"rotafreq"`
}

type TLogHandler struct {
	sliceLogObj []log.TLogObject
}

// using singleton
var instance *TLogHandler
var once sync.Once

func GetInstance() *TLogHandler {
	once.Do(func() {
		instance = &TLogHandler{}
	})
	return instance
}

func _getLog(_strLogName string) *log.TLogObject {
	inst := GetInstance()
	for i := range inst.sliceLogObj {
		if inst.sliceLogObj[i].GetLogName() == _strLogName {
			return &inst.sliceLogObj[i]
		}
	}
	return nil
}

func Initialize(
	_build string,
) error {
	inst := GetInstance()

	logPath := "./config/" + _build + "/log.yaml"
	byteData, err := ioutil.ReadFile(logPath)

	if err != nil {
		return err
	}

	myLog := TLogSettings{}
	err = yaml.Unmarshal(byteData, &myLog)
	if err != nil {
		return err
	}

	// LOG SETTING
	traceLogObj := log.TLogObject{}
	var iLogType int
	if myLog.Debug {
		iLogType = log.LOG_TYPE_ROTATE | log.LOG_TYPE_DEBUG
	} else {
		iLogType = log.LOG_TYPE_ROTATE
	}
	traceLogObj.SetLog(
		"trace",
		myLog.Path,
		iLogType,
		myLog.Level,
		myLog.RotaCnt,
		myLog.RotaFreq,
	)
	traceLogObj.SetCallStack(3)
	inst.sliceLogObj = append(inst.sliceLogObj, traceLogObj)

	byteData = nil
	return err
}

func Write(_strLogName string, _i32Level int32, _args ...interface{}) bool {
	logObj := _getLog(_strLogName)

	if logObj == nil {
		return false
	}

	var strTemp string
	for i := range _args {
		strTemp += fmt.Sprintf("%v", _args[i])
	}
	logObj.Write(_i32Level, strTemp)

	return true
}
