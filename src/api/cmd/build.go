// Created by Seunggwan, Back on 2024.04.23
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package cmd

import (
	commonCmd "gorani/common/cmd"
)

type TBuildMode string

const (
	BUILD_DEV   TBuildMode = "dev"
	BUILD_LIVE  TBuildMode = "live"
	BUILD_LOCAL TBuildMode = "local"
)

func (m TBuildMode) String() string {
	return string(m)
}

// get build mode from command arguments
// [dev|live|live-a|live-b|live-qa|local|qa]
// dev : gorani server
// live :
// local : developer local pc
func GetBuildFromCommand() string {
	arg := commonCmd.GetArg(
		commonCmd.STRING,             // arguments type
		"build",                      // arguments name
		"local",                      // default value
		"build mode[dev|live|local]", // desc
	)

	mode := *(arg.(*string))
	switch TBuildMode(mode) {
	case BUILD_DEV:
	case BUILD_LIVE:
	case BUILD_LOCAL:
	default:
		return BUILD_LOCAL.String()
	}

	return mode
}
