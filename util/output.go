package util

import (
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
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

func MarkdownToText(md string) string {
	return string(markdown.Render(md, 80, 6))
}
