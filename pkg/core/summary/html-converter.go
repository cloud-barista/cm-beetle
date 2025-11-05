/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package summary provides infrastructure summary HTML generation
package summary

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	emoji "github.com/yuin/goldmark-emoji"
)

// ConvertMarkdownToHTML converts markdown content to HTML
// It uses goldmark library with the following configurations:
// - Extensions: GFM (GitHub Flavored Markdown), Table, Strikethrough, Linkify, TaskList, Emoji
// - Parser: AutoHeadingID enabled (automatic heading ID generation)
// - Renderer: Unsafe mode (allow raw HTML), XHTML mode
func ConvertMarkdownToHTML(md []byte) []byte {
	var buf bytes.Buffer

	// Configure goldmark with GitHub Flavored Markdown support
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,           // GitHub Flavored Markdown
			extension.Table,         // Table support
			extension.Strikethrough, // Strikethrough support
			extension.Linkify,       // Auto-link URLs
			extension.TaskList,      // Task list support
			emoji.Emoji,             // Emoji support (:smile:, :tada:, etc.)
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Automatically assign IDs to headings
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allow raw HTML in markdown
			html.WithXHTML(),  // Render as XHTML
		),
	)

	// Convert markdown to HTML
	if err := markdown.Convert(md, &buf); err != nil {
		return []byte(fmt.Sprintf("<p>Error converting markdown to HTML: %v</p>", err))
	}

	// Wrap with complete HTML structure
	htmlHeader := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Infrastructure Summary</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        h1, h2, h3, h4, h5, h6 {
            color: #2c3e50;
            margin-top: 24px;
            margin-bottom: 16px;
            font-weight: 600;
            line-height: 1.25;
        }
        h1 { font-size: 2em; border-bottom: 2px solid #eaecef; padding-bottom: 0.3em; }
        h2 { font-size: 1.5em; border-bottom: 1px solid #eaecef; padding-bottom: 0.3em; }
        h3 { font-size: 1.25em; }
        table {
            border-collapse: collapse;
            width: 100%;
            margin: 16px 0;
            background-color: white;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
        }
        th, td {
            border: 1px solid #dfe2e5;
            padding: 12px;
            text-align: left;
        }
        th {
            background-color: #f6f8fa;
            font-weight: 600;
        }
        tr:nth-child(even) {
            background-color: #f9f9f9;
        }
        code {
            background-color: #f6f8fa;
            border-radius: 3px;
            padding: 2px 6px;
            font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
            font-size: 85%;
        }
        pre {
            background-color: #f6f8fa;
            border-radius: 6px;
            padding: 16px;
            overflow: auto;
            line-height: 1.45;
        }
        pre code {
            background-color: transparent;
            padding: 0;
        }
        a {
            color: #0366d6;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        hr {
            border: 0;
            border-top: 1px solid #eaecef;
            margin: 24px 0;
        }
        blockquote {
            border-left: 4px solid #dfe2e5;
            padding: 0 16px;
            color: #6a737d;
            margin: 0;
        }
        ul, ol {
            padding-left: 2em;
            margin: 16px 0;
        }
        li {
            margin: 4px 0;
        }
    </style>
</head>
<body>
`
	htmlFooter := `
</body>
</html>`

	// Combine HTML parts
	var result bytes.Buffer
	result.WriteString(htmlHeader)
	result.Write(buf.Bytes())
	result.WriteString(htmlFooter)

	return result.Bytes()
}
