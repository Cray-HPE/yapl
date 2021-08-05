package util

import (
	markdown "github.com/MichaelMure/go-term-markdown"
)

func MarkdownToText(md string) string {
	return string(markdown.Render(md, 80, 6))
}
