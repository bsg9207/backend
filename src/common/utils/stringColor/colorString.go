// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package stringColor

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
	CodeExpr     = `\033\[[\d;?]+m`
)

var (
	CodeRegex = regexp.MustCompile(CodeExpr)
)

type TColor uint8

const (
	FgBase uint8 = 30
	BgBase uint8 = 40
)

const (
	FgBlack TColor = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// FgDefault revert default FG
	FgDefault TColor = 39
)

func (c TColor) Code() string {
	return strconv.Itoa(int(c))
}

func (c TColor) Paint(args ...interface{}) string {
	var str string
	if length := len(args); length == 0 {
		return ""
	}

	str = fmt.Sprint(args...)
	code := c.Code()
	if len(code) == 0 {
		return str
	}

	return fmt.Sprintf(FullColorTpl, code, str)
}

func ClearColorCode(str string) string {
	return CodeRegex.ReplaceAllString(str, "")
}
