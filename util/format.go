package util

import (
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/fatih/color"
)

func Indent(text, indent string) string {
	if text[len(text)-1:] == "\n" {
		result := ""
		for _, j := range strings.Split(text[:len(text)-1], "\n") {
			result += indent + j + "\n"
		}
		return result
	}
	result := ""
	for _, j := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		result += indent + j + "\n"
	}
	return result[:len(result)-1]
}

var Green = color.New(color.FgGreen)
var Red = color.New(color.FgRed, color.Bold)
var Yellow = color.New(color.FgYellow)
var Blue = color.New(color.FgBlue)
var White = color.New(color.FgWhite)

func MarkdownToText(md string) string {
	return string(markdown.Render(md, 80, 6))
}
