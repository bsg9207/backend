// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

import (
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
)

const (
	WIDTH_SOURCE   = 102
	WIDTH_STACK    = WIDTH_SOURCE - 4
	WIDTH_DEFAULT  = WIDTH_SOURCE
	PADDING_STRING = "  "
)

var (
	gopath = os.Getenv("GOPATH")

	DefaultLoggerPrintFunc = func(format string, data ...interface{}) {
		htFormat := Highlighted(format)
		fmt.Printf(htFormat+"\n", data...)
	}

	DefaultLogger = &TErrorLogger{
		config: &TErrorLoggerConfig{
			PrintFunc:               DefaultLoggerPrintFunc,
			LinesOfBefore:           5,
			LinesOfAfter:            5,
			PrintStack:              true,
			PrintSource:             true,
			PrintError:              true,
			ExitOnDebugSuccess:      false,
			DisableStackIndentation: false,
		},
	}
)

func LogError(uErr error) bool {
	DefaultLogger.Overload(1)
	return DefaultLogger.LogError(uErr)
}

type IErrorLogger interface {
	LogError(err error) bool
	SetConfig(config *TErrorLoggerConfig)
	Config() *TErrorLoggerConfig
}

type TErrorLoggerConfig struct {
	PrintFunc               func(Format string, data ...interface{})
	LinesOfBefore           int
	LinesOfAfter            int
	PrintStack              bool
	PrintSource             bool
	PrintError              bool
	ExitOnDebugSuccess      bool
	DisableStackIndentation bool
}

type TPrintSourceOptions struct {
	FuncLine    int
	StartLine   int
	EndLine     int
	Highlighted map[int][]int
}

type TErrorLogger struct {
	config             *TErrorLoggerConfig
	stackDepthOverload int
}

func (l *TErrorLogger) Printf(_format string, _data ...interface{}) {
	l.config.PrintFunc(_format, _data...)
}

func (l *TErrorLogger) Overload(_amount int) {
	l.stackDepthOverload += _amount
}

func (l *TErrorLogger) LogError(_uErr error) bool {
	if _uErr == nil {
		return false
	}

	stackTraceLines := parseStackTrace(1 + l.stackDepthOverload)
	if stackTraceLines == nil || len(stackTraceLines) < 1 {
		l.Printf("Error: '%s'", _uErr)
		l.Printf("The stack trace is empty.")
		return true
	}

	if l.config.PrintError {
		l.Printf("Error in '%s()': [%s]", stackTraceLines[0].CallingObject, _uErr.Error())
	}

	if l.config.PrintSource {
		l.PrintTraceback(stackTraceLines, l.config.PrintStack)
	}

	if l.config.ExitOnDebugSuccess {
		os.Exit(1)
	}

	l.stackDepthOverload = 0

	return true
}

func (l *TErrorLogger) makeLogForSource(lines []string, opts TPrintSourceOptions) []string {
	var strSource []string
	if opts.FuncLine != -1 && opts.FuncLine < opts.StartLine {
		tmp := fmt.Sprintf("%s%03d  %s", PADDING_STRING, opts.FuncLine+1, lines[opts.FuncLine])
		strSource = append(strSource, tmp)
		if opts.FuncLine < opts.StartLine-1 {
			strSource = append(strSource, "    ...")
		}
	}

	for i := opts.StartLine; i < opts.EndLine; i++ {
		if _, ok := opts.Highlighted[i]; !ok || len(opts.Highlighted[i]) != 2 {
			tmp := fmt.Sprintf("%s%03d  %s", PADDING_STRING, i+1, lines[i])
			strSource = append(strSource, tmp)
			continue
		}

		hlStart := max(opts.Highlighted[i][0], 0)
		hlEnd := min(opts.Highlighted[i][1], len(lines)-1)
		tmp := fmt.Sprintf("%s%03d  %s%s%s", "❱ ", i+1, lines[i][:hlStart],
			lines[i][hlStart:hlEnd+1], lines[i][hlEnd+1:])
		strSource = append(strSource, tmp)
	}

	return strSource
}

func (l *TErrorLogger) makeLogForStack(stackTraceLines []TStackTraceItem) []string {
	lengthLines := len(stackTraceLines)
	var stackStrings []string
	for i := lengthLines - 1; i >= 0; i-- {
		padding := ""
		if !l.config.DisableStackIndentation {
			for j := 0; j < lengthLines-1-i; j++ {
				padding += "  "
			}
		}

		tmp := fmt.Sprintf("%s%s() (%s:%d)", padding, stackTraceLines[i].CallingObject,
			stackTraceLines[i].SourcePathRef, stackTraceLines[i].SourceLineRef)
		stackStrings = append(stackStrings, tmp)
	}

	return stackStrings
}

func (l *TErrorLogger) PrintTraceback(stackTraceLines []TStackTraceItem, bPrintStack bool) {
	filePath := stackTraceLines[0].SourcePathRef
	shortFilePath := filePath
	if gopath != "" {
		shortFilePath = strings.Replace(filePath, gopath+"/src/", "", -1)
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		l.Printf("Can't read file %s : %s.", filePath, err.Error())
		return
	}

	lines := strings.Split(string(b), "\n")
	debugLineNumber := stackTraceLines[0].SourceLineRef

	minLine := debugLineNumber - l.config.LinesOfBefore
	maxLine := debugLineNumber + l.config.LinesOfAfter

	deleteBlankLinesFromRange(lines, &minLine, &maxLine)

	lines = lines[:maxLine+1]

	funcLine := findFuncLine(lines, debugLineNumber)
	if funcLine > minLine {
		minLine = funcLine + 1
	}

	failingLineIndex, columnStart, columnEnd := findFailingLine(lines, funcLine, debugLineNumber)

	var srcRef string
	if failingLineIndex != -1 {
		srcRef = fmt.Sprintf("'%s:%d' in %s", shortFilePath, failingLineIndex+1,
			stackTraceLines[0].CallingObject)
	} else {
		srcRef = fmt.Sprintf("error in [%s] (not found failing line, func call is at line [%d])",
			shortFilePath, debugLineNumber)
	}

	srcpart := l.makeLogForSource(lines, TPrintSourceOptions{
		FuncLine: funcLine,
		Highlighted: map[int][]int{
			failingLineIndex: {columnStart, columnEnd},
		},
		StartLine: minLine,
		EndLine:   maxLine,
	})

	var source []string
	source = append(source, srcRef)
	source = append(source, " ") //padding
	source = append(source, srcpart...)

	if bPrintStack {
		stack := l.makeLogForStack(stackTraceLines)
		boxedStack := makeTitleBoxString("Stack Trace", stack, WIDTH_STACK)

		var traceback []string
		traceback = append(traceback, source...)
		traceback = append(traceback, " ") // padding
		traceback = append(traceback, boxedStack...)

		boxedTraceback := makeBoxedString(traceback, WIDTH_SOURCE, false)
		for _, tb := range boxedTraceback {
			l.Printf(tb)
		}
	} else {
		boxedSource := makeBoxedString(source, WIDTH_SOURCE, false)
		for _, src := range boxedSource {
			l.Printf(src)
		}
	}

}

// stack trace item
type TStackTraceItem struct {
	CallingObject string
	Args          []string
	SourcePathRef string
	SourceLineRef int
}

const (
	Stack            = `((?:(?:[a-zA-Z._-]+)[/])*(?:[*a-zA-Z0-9_]*\.)+[a-zA-Z0-9_]+)\(((?:(?:0x[0-9a-f]+)|(?:...)[,\s]*)+)*\)[\s]+([/:\-a-zA-Z0-9\._]+)[:]([0-9]+)[\s](?:\+0x([0-9a-f]+))*`
	HexNumber        = `0x[0-9a-f]+`
	FuncLine         = `^func[\s][a-zA-Z0-9]+[(](.*)[)][\s]*{`
	ErrorLineFunc    = `[\.]LogError[\(](.*)[/)]`
	ErrorLineVarName = `LogError[\(](.*)\)`
	VarDefinition    = `%s[\s\:]*={1}([\s]*[a-zA-Z0-9\._]+)`
)

var (
	regexpParseStack                 = regexp.MustCompile(Stack)
	regexpHexNumber                  = regexp.MustCompile(HexNumber)
	regexpFuncLine                   = regexp.MustCompile(FuncLine)
	regexpParseDebugLineFunc         = regexp.MustCompile(ErrorLineFunc)
	regexpParseDebugLineParseVarName = regexp.MustCompile(ErrorLineVarName)
	regexpFindVarDefinition          = func(varName string) *regexp.Regexp {
		return regexp.MustCompile(fmt.Sprintf(VarDefinition, varName))
	}
)

func parseStackTrace(deltaDepth int) []TStackTraceItem {
	// fmt.Println(string(debug.Stack()))
	return parseAnyStackTrace(string(debug.Stack()), deltaDepth)
}

func parseAnyStackTrace(stackString string, deltaDepth int) []TStackTraceItem {
	stackArray := strings.Split(stackString, "\n")
	if len(stackArray) < 2*(2+deltaDepth) {
		return nil
	}

	stack := strings.Join(stackArray[2*(2+deltaDepth):], "\n")
	parseResult := regexpParseStack.FindAllStringSubmatch(stack, -1)

	sti := make([]TStackTraceItem, len(parseResult))
	for i := range parseResult {
		args := regexpHexNumber.FindAllString(parseResult[i][2], -1)
		srcLine, err := strconv.Atoi(parseResult[i][4])
		if LogError(err) {
			srcLine = -1
		}

		sti[i] = TStackTraceItem{
			CallingObject: parseResult[i][1],
			Args:          args,
			SourcePathRef: parseResult[i][3],
			SourceLineRef: srcLine,
		}
	}

	return sti
}

func findFuncLine(lines []string, lineNumber int) int {
	for i := lineNumber; i > 0; i-- {
		if regexpFuncLine.Match([]byte(lines[i])) {
			return i
		}
	}

	return -1
}

func findFailingLine(lines []string, funcLine int, debugLine int) (failingLineIndex, columnStart, columnEnd int) {
	failingLineIndex = -1

	reMatches := regexpParseDebugLineParseVarName.FindStringSubmatch(lines[debugLine-1])
	if len(reMatches) < 2 {
		return
	}
	varName := reMatches[1]

	reFindVar := regexpFindVarDefinition(varName)

	for i := debugLine; i >= funcLine && i > 0; i-- {
		// fmt.Printf("%d: %s", i, lines[i])

		if strings.Trim(lines[i], "\n\t") == "" {
			// fmt.Printf("%d: ignoring blank line", i)
			continue
		} else if len(lines[i]) >= 2 && lines[i][:2] == "//" {
			// fmt.Printf("ignoring comment line", i)
			continue
		}

		index := reFindVar.FindStringSubmatchIndex(lines[i])
		if index == nil {
			// fmt.Printf("%d: var definition not found for '%s'", i, varName)
			continue
		}

		failingLineIndex = i
		columnStart = index[0]

		openedBrackets, closedBrackets := 0, 0
		for j := index[1]; j < len(lines[i]); j++ {
			if lines[i][j] == '(' {
				openedBrackets++
			} else if lines[i][j] == ')' {
				closedBrackets++
			}

			if openedBrackets == closedBrackets {
				columnEnd = j
				return
			}
		}

		if columnEnd == 0 {
			// fmt.Printf("Fixing value of columnEnd(0). Defaulting to end of failing line")
			columnEnd = len(lines[i]) - 1
		}

		return
	}

	return
}
