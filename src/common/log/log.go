// Created by Seunggwan, Back on 2024.04.18
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"gorani/common/utils"
	"gorani/common/utils/file"
)

var lockMutex sync.Mutex

const (
	LOG_TYPE_ROTATE = 1
	LOG_TYPE_STDOUT = 2
	LOG_TYPE_DEBUG  = 4
)

type TLogObject struct {
	strName  string
	strPath  string
	iLogType int
	i32Level int32

	strModulePath  string
	iModulePathLen int

	i64Created      int64
	i32RotationCnt  int32
	i64RotationFreq int64

	i32CallStack int
}

func (obj *TLogObject) GetLogName() string {
	return obj.strName
}

func (obj *TLogObject) SetLog(_strName string, _strPath string, _iLogType int, _i32Level int32, _i32RotationCnt int32, _i64RotationFreq int64) {
	obj.strName = _strName
	obj.strPath = _strPath
	obj.i32Level = _i32Level
	obj.iLogType = _iLogType
	obj.i32RotationCnt = _i32RotationCnt
	obj.i64RotationFreq = _i64RotationFreq
	obj.strModulePath, _ = os.Getwd()
	obj.iModulePathLen = len(obj.strModulePath)

	obj.i32CallStack = 2

	file.Mkdir(_strPath)
}

func (obj *TLogObject) SetCallStack(_iCall int) {
	obj.i32CallStack = _iCall
}

func (obj *TLogObject) Write(_i32Level int32, _strLog string) {

	var i64TmCurrent int64

	if obj.i32Level < _i32Level {
		return
	}

	i64TmCurrent = time.Now().Unix()
	if obj._tryRotate(i64TmCurrent) == false {
	}

	strText := obj._compileLogstring(_strLog)

	strFilePath := fmt.Sprintf("%v/%v.txt", obj.strPath, obj.strName)
	if (obj.iLogType & LOG_TYPE_ROTATE) == LOG_TYPE_ROTATE {
		_writeFile(strFilePath, strText)
	} else if (obj.iLogType & LOG_TYPE_STDOUT) == LOG_TYPE_STDOUT {
		_writeConsole(strText)
	}
}

func (obj *TLogObject) _tryRotate(_i64TmCurrent int64) bool {

	if obj.i32RotationCnt == 0 {
		return false
	}

	strMetaFilePath := fmt.Sprintf("%v/%v.meta", obj.strPath, obj.strName)

	if obj.i64Created == 0 {
		if file.ExistPath(strMetaFilePath) == true {
			strCTime, _ := _readFromMetaFile(strMetaFilePath)
			obj.i64Created = utils.StringToI64(strCTime, 10)
		} else {
			obj.i64Created = _i64TmCurrent
			_writeToMetaFile(utils.I64ToString(obj.i64Created, 10), strMetaFilePath)
		}
	}

	if obj.i64Created > 0 {
		i64TempCTime := (obj.i64Created / obj.i64RotationFreq) * obj.i64RotationFreq
		i64TempCurTime := (_i64TmCurrent / obj.i64RotationFreq) * obj.i64RotationFreq

		if i64TempCTime < i64TempCurTime {
			//Try Rotate
			strLastFilePath := fmt.Sprintf("%v/%v.txt.%v", obj.strPath, obj.strName, obj.i32RotationCnt)

			if file.ExistPath(strLastFilePath) == true {
				file.Remove(strLastFilePath)
			}

			var i32 int32
			for i32 = obj.i32RotationCnt - 1; i32 >= 0; i32-- {
				var strNextFilePath string
				var strPreFilePath string
				if i32 == 0 {
					strPreFilePath = fmt.Sprintf("%v/%v.txt", obj.strPath, obj.strName)
					strNextFilePath = fmt.Sprintf("%v/%v.txt.%v", obj.strPath, obj.strName, i32+1)
					//Metafile 갱신
					obj.i64Created = _i64TmCurrent
					_writeToMetaFile(fmt.Sprintf("%v", obj.i64Created), strMetaFilePath)
				} else {
					strPreFilePath = fmt.Sprintf("%v/%v.txt.%v", obj.strPath, obj.strName, i32)
					strNextFilePath = fmt.Sprintf("%v/%v.txt.%v", obj.strPath, obj.strName, i32+1)
				}

				if file.ExistPath(strPreFilePath) == true {
					err := file.Rename(strPreFilePath, strNextFilePath)
					if err != nil {
						//TODO
					}
				}
			}
			return true
		}
	} else {
		//Crate New Log file
		fmt.Println("New LogFile.....")
	}
	return false
}

// CompileLogstring : Compile Log string from args
func (obj *TLogObject) _compileLogstring(_strLog string) string {
	var strText string
	localTime := time.Now().Local()

	strText += localTime.Format("2006-01-02 15:04:05") + "\t"
	if (obj.iLogType & LOG_TYPE_DEBUG) == 4 {
		_, strFile, iLine, _ := runtime.Caller(obj.i32CallStack)
		strText += fmt.Sprintf("/%s %d\t", filepath.Base(strFile), iLine)
	}
	return strText + _strLog + "\r\n"
}

// WriteToMetaFile : write to meta file
func _writeToMetaFile(_strText string, _strLogFile string) {
	lockMutex.Lock()
	defer lockMutex.Unlock()
	f, e := os.OpenFile(_strLogFile, os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()
	f.WriteString(_strText)
}

// ReadFromMetaFile : read from meta file
func _readFromMetaFile(_strLogMetaFile string) (string, error) {
	lockMutex.Lock()
	defer lockMutex.Unlock()
	f, e := os.OpenFile(_strLogMetaFile, os.O_CREATE|os.O_RDWR, 0644)
	if e != nil {
		fmt.Println(e)
		return "", e
	}
	defer f.Close()
	buffer := make([]byte, 10)
	_, e = f.Read(buffer)
	if e != nil {
		return "", e
	}
	return string(buffer), nil
}

// WriteFile : write log at file
func _writeFile(_strLogFile string, _strText string) {
	lockMutex.Lock()
	defer lockMutex.Unlock()
	file.WriteFileFromString(_strLogFile, _strText)
}

// WriteConsole : write log at stanadard output buffer
func _writeConsole(_text string) {
	fmt.Println(_text)
}
