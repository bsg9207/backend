// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

import (
	"regexp"

	"gorani/common/utils/stringColor"
)

const (
	// boxed text border
	NormalBorder_Top = "┌[─]{1,}.{0,}[─]{1,}┐"
	NormalBorder_Bot = "└[─]{2,}┘"

	BoldedBorder_Top = "┏[━]{1,}.{0,}[─]{1,}┓"
	BoldedBorder_Bot = "┗[━]{2,}┛"

	RoundedBorder_Top = "╭[─]{1,}.{0,}[─]{1,}╮"
	RoundedBorder_Bot = "╰[─]{2,}╯"

	VerticalBoder = "^│.*│$"

	// error arrow
	ErrorArrow = "❱"

	// src file path and line number
	SrcPathRef = "((\\/.{0,}){0,}:(\\d+))"
)

var (
	// border
	NBTRegex = regexp.MustCompile(NormalBorder_Top)
	NBBRegex = regexp.MustCompile(NormalBorder_Bot)

	BBTRegex = regexp.MustCompile(BoldedBorder_Top)
	BBBRegex = regexp.MustCompile(BoldedBorder_Bot)

	RBTRegex = regexp.MustCompile(RoundedBorder_Top)
	RBBRegex = regexp.MustCompile(RoundedBorder_Bot)

	VBRegex = regexp.MustCompile(VerticalBoder)

	// error arrow
	EARegex = regexp.MustCompile(ErrorArrow)

	// src file path and line number
	SPRRegex = regexp.MustCompile(SrcPathRef)

	green = stringColor.FgGreen.Paint
	red   = stringColor.FgRed.Paint
	cyan  = stringColor.FgCyan.Paint
)

func Highlighted(format string) string {
	strFormat := format

	// border -> green
	strFormat = borderHighlighted(strFormat)

	// vertical border -> green
	strFormat = verticalBoderHighlighted(strFormat)

	// error arrow -> red
	strFormat = errorArrowHighlighted(strFormat)

	// src file path -> green
	strFormat = sourcePathHighlighted(strFormat)

	return strFormat
}

func borderHighlighted(fmt string) string {
	strFmt := fmt

	if NBTRegex.MatchString(strFmt) {
		for _, str := range NBTRegex.FindAllString(strFmt, -1) {
			strFmt = NBTRegex.ReplaceAllString(strFmt, green(str))
		}
	} else if NBBRegex.MatchString(strFmt) {
		for _, str := range NBBRegex.FindAllString(strFmt, -1) {
			strFmt = NBBRegex.ReplaceAllString(strFmt, green(str))
		}
	} else if BBTRegex.MatchString(strFmt) {
		for _, str := range BBTRegex.FindAllString(strFmt, -1) {
			strFmt = BBTRegex.ReplaceAllString(strFmt, green(str))
		}
	} else if BBBRegex.MatchString(strFmt) {
		for _, str := range BBBRegex.FindAllString(strFmt, -1) {
			strFmt = BBBRegex.ReplaceAllString(strFmt, green(str))
		}
	} else if RBTRegex.MatchString(strFmt) {
		for _, str := range RBTRegex.FindAllString(strFmt, -1) {
			strFmt = RBTRegex.ReplaceAllString(strFmt, green(str))
		}
	} else if RBBRegex.MatchString(strFmt) {
		for _, str := range RBBRegex.FindAllString(strFmt, -1) {
			strFmt = RBBRegex.ReplaceAllString(strFmt, green(str))
		}
	}

	return strFmt
}

func verticalBoderHighlighted(fmt string) string {
	strFmt := fmt
	if VBRegex.MatchString(strFmt) {
		tmpregex := regexp.MustCompile("[│]")
		strFmt = tmpregex.ReplaceAllString(strFmt, green("│"))
	}

	return strFmt
}

func errorArrowHighlighted(fmt string) string {
	strFmt := fmt
	if EARegex.MatchString(strFmt) {
		tmpregex := regexp.MustCompile("[❱]")
		strFmt = tmpregex.ReplaceAllString(strFmt, red("❱"))
	}

	return strFmt
}

func sourcePathHighlighted(fmt string) string {
	strFmt := fmt
	if SPRRegex.MatchString(strFmt) {
		for _, str := range SPRRegex.FindAllString(strFmt, -1) {
			strFmt = SPRRegex.ReplaceAllString(strFmt, cyan(str))
		}
	}

	return strFmt
}
