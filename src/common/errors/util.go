// Created by Seunggwan, Back on 2024.04.18
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"gorani/common/utils/stringColor"
)

func deleteBlankLinesFromRange(lines []string, start, end *int) {
	//clean from out of range values
	(*start) = max(*start, 0)
	(*end) = min(*end, len(lines)-1)

	//clean leading blank lines
	for (*start) <= (*end) {
		if strings.Trim(lines[(*start)], " \n\t") != "" {
			break
		}
		(*start)++
	}

	//clean trailing blank lines
	for (*end) >= (*start) {
		if strings.Trim(lines[(*end)], " \n\t") != "" {
			break
		}
		(*end)--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func replaceTabToSpace(strOrigin string) string {
	tabString := makeTabWithSpace()
	outString := strings.ReplaceAll(strOrigin, "\t", tabString)
	return outString
}

func removeCarrageReturn(strOrigin string) string {
	outString := strings.ReplaceAll(strOrigin, "\r", "")
	return outString
}

func makeTabWithSpace() string {
	var out string
	for i := 0; i < TAB_SIZE; i++ {
		out += " "
	}
	return out
}

// boxed string
const (
	TAB_SIZE  = 4
	MAX_WIDTH = 202
)

// B for beginning
// M for middle
// E for end
type TRowStyle struct {
	B string
	M string
	E string
}

// T for top
// M for Middle
// B for botton
type TBorderStyle struct {
	T TRowStyle
	M TRowStyle
	B TRowStyle
}

var (
	NormalBorder = TBorderStyle{
		T: TRowStyle{B: "┌─", M: "─", E: "─┐"},
		M: TRowStyle{B: "│ ", M: " ", E: " │"},
		B: TRowStyle{B: "└─", M: "─", E: "─┘"},
	}

	BoldedBorder = TBorderStyle{
		T: TRowStyle{B: "┏━", M: "━", E: "━┓"},
		M: TRowStyle{B: "┃ ", M: " ", E: " ┃"},
		B: TRowStyle{B: "┗━", M: "━", E: "━┛"},
	}

	RoundedBorder = TBorderStyle{
		T: TRowStyle{B: "╭─", M: "─", E: "─╮"},
		M: TRowStyle{B: "│ ", M: " ", E: " │"},
		B: TRowStyle{B: "╰─", M: "─", E: "─╯"},
	}
)

func row(rowStyle TRowStyle, length int) string {
	rowString := rowStyle.B
	width := min(length, MAX_WIDTH)
	rowString = fmt.Sprintf("%s%s", rowString, strings.Repeat(rowStyle.M, width))
	rowString = fmt.Sprintf("%s%s", rowString, rowStyle.E)
	return rowString
}

func (s *TBorderStyle) titleRow(strTitle string, titlePos int, length int) string {
	// tab은 모두 space로 변경
	title := replaceTabToSpace(strTitle)

	// title 양쪽에 가독성을 위한 여백 추가
	title = " " + title + " "

	// step beginning
	out := s.T.B

	// step title
	out = fmt.Sprintf("%s%s", out, strings.Repeat(s.T.M, titlePos))

	// color code가 들어오면 문자열 길이가 정확히 측정이 안됨
	// color code를 제거하여 길이를 측정하도록 한다.
	countTitle := utf8.RuneCountInString(stringColor.ClearColorCode(title))

	width := min(length, MAX_WIDTH)
	if width < countTitle {
		// width보다 길면 생략
		out += title[:(width-3)] + "..."
	} else {
		out += title
	}

	// step middle
	countTitle += titlePos
	if countTitle < width {
		out = fmt.Sprintf("%s%s", out, strings.Repeat(s.T.M, width-countTitle))
	}

	// step end
	out += s.T.E
	return out
}

func (s *TBorderStyle) topRow(length int) string {
	return row(s.T, length)
}

func (s *TBorderStyle) bottonRow(length int) string {
	return row(s.B, length)
}

func (s *TBorderStyle) middleRow(strOrigin string, length int) string {
	// tab을 모두 space로 변경
	ori := replaceTabToSpace(strOrigin)

	// carrage return 제거
	ori = removeCarrageReturn(ori)

	// color code가 들어오면 문자열 길이가 정확히 측정이 안됨
	// color code를 제거하여 길이를 측정하도록 한다.
	countOri := utf8.RuneCountInString(stringColor.ClearColorCode(ori))

	// step beginning
	out := s.M.B

	// step middle
	width := min(length, MAX_WIDTH)
	if width < countOri {
		// width보다 길면 생략
		out = fmt.Sprintf("%s%s", out, (ori[:(width-7)] + "...    "))
	} else {
		out = fmt.Sprintf("%s%s", out, ori)
	}

	if countOri < width {
		out = fmt.Sprintf("%s%s", out, strings.Repeat(s.M.M, width-countOri))
	}

	// step end
	out += s.M.E
	return out
}

const (
	PADDING_STEP     = 2
	MAX_STRING_WIDTH = MAX_WIDTH - (2 * PADDING_STEP)
	TITLE_POSITION   = 1
)

func makeTitleBoxString(strTitle string, normalString []string, width int) []string {
	var boxedString []string
	boxedString = append(boxedString, RoundedBorder.titleRow(strTitle, TITLE_POSITION, width))
	boxedString = append(boxedString, RoundedBorder.middleRow(" ", width)) // padding
	for _, str := range normalString {
		boxedString = append(boxedString, RoundedBorder.middleRow(str, width))
	}
	boxedString = append(boxedString, RoundedBorder.bottonRow(width))

	return boxedString
}

// input []string
// output []string -> box로 감싸진 string 배열
func makeBoxedString(normalString []string, width int, needPadding bool) []string {
	var boxedString []string
	boxedString = append(boxedString, RoundedBorder.topRow(width))

	if needPadding {
		boxedString = append(boxedString, RoundedBorder.middleRow(" ", width))
	}

	for _, str := range normalString {
		boxedString = append(boxedString, RoundedBorder.middleRow(str, width))
	}

	if needPadding {
		boxedString = append(boxedString, RoundedBorder.middleRow(" ", width))
	}

	boxedString = append(boxedString, RoundedBorder.bottonRow(width))
	return boxedString
}
