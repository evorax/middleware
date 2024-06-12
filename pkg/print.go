package middle

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	infoColor    = color.New(color.FgGreen).SprintFunc()
	warningColor = color.New(color.FgYellow).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
)

func formatMessage(level, message string, colorFunc func(a ...interface{}) string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s [%s] %s", timestamp, colorFunc(level), message)
}

func Info(args ...any) {
	var message string
	for _, arg := range args {
		message += fmt.Sprint(arg, " ")
	}
	fmt.Println(formatMessage("INFO", message, infoColor))
}

func Warning(args ...any) {
	var message string
	for _, arg := range args {
		message += fmt.Sprint(arg, " ")
	}
	fmt.Println(formatMessage("WARNING", message, warningColor))
}

func Error(args ...any) {
	var message string
	for _, arg := range args {
		message += fmt.Sprint(arg, " ")
	}
	fmt.Println(formatMessage("ERROR", message, errorColor))
}
