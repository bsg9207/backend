// Created by Seunggwan, Back on 2024.04.23
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package main

import (
	"gorani/api/cmd"
	"gorani/api/logHandler"
	"gorani/api/preference"
	"gorani/api/routers"
	"gorani/api/worker"
)

func main() {
	// get build mode
	build := cmd.GetBuildFromCommand()

	// log handler 초기화
	logHandler.Initialize(build)
	logHandler.Write("trace", 0, "initialize log handler")

	// 설정 초기화
	preference.Initialize(build)
	logHandler.Write("trace", 0, "initialize preference")

	// api server 시작
	logHandler.Write("trace", 0, "start gorani api server")
	worker.Run(routers.DefaultAPITable())
}
