// Created by Seunggwan, Back on 2024.04.18
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package file

import (
	"fmt"
	"gorani/common/errors"
	"os"
)

func Mkdir(_strPath string) error {
	if _, serr := os.Stat(_strPath); serr != nil {
		merr := os.MkdirAll(_strPath, os.ModePerm)
		if merr != nil {
			return errors.ERROR_MKDIR(merr.Error())
		}
	}
	return nil
}

func WriteFileFromString(_strPath string, _strText string) {
	f, e := os.OpenFile(_strPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()
	f.WriteString(_strText)
}

func ExistPath(_strPath string) bool {
	if _, err := os.Stat(_strPath); err != nil {
		return false
	}
	return true
}

func Remove(_strPath string) error {
	return os.Remove(_strPath)
}

func Rename(_strFrom string, _strTo string) error {
	return os.Rename(_strFrom, _strTo)
}
