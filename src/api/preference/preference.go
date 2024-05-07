// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package preference

import (
	"fmt"
	"os"
	"sync"

	"gorani/api/preference/section"
	"gorani/common/errors"
	"gorani/common/utils"
	YamlLoader "gorani/common/yaml"
)

type TPreference struct {
	strModulePath  string
	iModulePathLen int

	//yaml
	mapYaml map[string]interface{}

	//lock
	lockMutex sync.Mutex

	//attributes
	tServer section.TServer
}

var instance *TPreference
var once sync.Once

func GetInstance() *TPreference {
	once.Do(func() {
		instance = &TPreference{}
	})
	return instance
}

func Initialize(_build string) {
	defer func() {
		s := recover()
		if s != nil {
			recovered, ok := s.(errors.TError)
			if ok {
				errors.LogError(recovered)
			} else {
				err := fmt.Errorf("%v", s)
				errors.LogError(err)
			}
		}
	}()
	// load preference
	LoadPreference(_build)

	// set preference
	SetPreference()
}

func LoadPreference(_build string) {
	inst := GetInstance()

	// set module path
	inst.strModulePath, _ = os.Getwd()
	inst.iModulePathLen = len(inst.strModulePath)

	// load yaml
	inst.lockMutex.Lock()
	defer inst.lockMutex.Unlock()

	inst.mapYaml = nil
	inst.mapYaml = make(map[string]interface{})

	for {
		path := "./config/" + _build + "/preference.yaml"
		e := YamlLoader.LoadFromFile(path, &inst.mapYaml)
		if error(e) != nil {
			panic(e)
		}
		break
	}

}

func SetPreference() {
	// print preference
	fmt.Printf(
		"\x1b[32m%s\033[0m \n",
		"#################### Preference ####################",
	)

	// set common preference
	onSetCommonPreference()
}

func printConfig(_sectionName string, _config interface{}) {
	if _config != nil {
		fmt.Printf("\x1b[32m%s\033[0m ", _sectionName)
		fmt.Println(utils.InterfaceToJsonString(_config, true))
	}
}
