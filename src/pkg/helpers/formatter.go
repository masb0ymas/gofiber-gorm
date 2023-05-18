package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"unicode"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

type Options struct {
	Label string
}

// Print Log
func PrintLog(title string, message string, options ...Options) string {
	var labelType string

	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	logName := green("[server]:")

	if options == nil {
		labelType = blue(title)
	}

	// check object struct
	for _, opt := range options {
		if opt.Label == "success" {
			labelType = blue(title)
		} else if opt.Label == "warning" {
			labelType = yellow(title)
		} else if opt.Label == "error" {
			labelType = red(title)
		} else {
			labelType = blue(title)
		}
	}

	newMessage := cyan(message)
	result := fmt.Sprintf("%s %s %s", logName, labelType, newMessage)

	return result
}

// Pretty JSON
func PrettyJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}

// Parse UUID
func ParseUUID(value any) uuid.UUID {
	byteUID, _ := json.Marshal(value)
	newUID := uuid.Must(uuid.ParseBytes(byteUID))

	return newUID
}

// check uuid is valid
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// check value is digit
func IsDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// check value is number
func IsNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsNumber(c) {
			return false
		}
	}
	return true
}
