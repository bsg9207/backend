// Created by Seunggwan, Back on 2024.04.23
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package main

import (
	"gorani/common/cmd"
)

// go run .\main.go -mode=dev
// dev
func main() {
	arg := cmd.GetArg(cmd.STRING, "mode", "nothing", "to work")
	argStr := arg.(*string)

	println(*argStr)
}
