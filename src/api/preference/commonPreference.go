// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package preference

import "gorani/api/preference/section"

func onSetCommonPreference() {
	inst := GetInstance()

	// host
	_findToString(section.KEY_SERVER_HOST, &inst.tServer.Host)

	// debug
	_findToBool(section.KEY_SERVER_DEBUG, &inst.tServer.Debug)

	// print config
	printConfig(section.SECTION_SERVER, inst.tServer)
}
