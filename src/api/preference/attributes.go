// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package preference

// //////////////////////////////////////////////////////////////////////////////
// section - server
func GetHost() string {
	inst := GetInstance()
	return inst.tServer.Host
}

func IsDebugMode() bool {
	inst := GetInstance()

	return inst.tServer.Debug
}
