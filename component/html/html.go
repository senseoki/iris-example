package html

import (
	"html/template"
	"time"
)

// MakeTagCSS ...
func MakeTagCSS(tags []string) template.HTML {
	t := time.Now()
	ts := t.Format("20060102150405")
	var result string
	for _, value := range tags {
		result += `<link rel="stylesheet" type="text/css" href="/static/` + value + `?t=` + ts + `" />`
	}
	return template.HTML(result)
}

// MakeTagExternalCSS ...
func MakeTagExternalCSS(tags []string) template.HTML {
	var result string
	for _, value := range tags {
		result += `<link rel="stylesheet" type="text/css" href="` + value + `" />`
	}
	return template.HTML(result)
}

// MakeTagJavascript function gets the internal JavaScript library.
func MakeTagJavascript(tags []string) template.HTML {
	t := time.Now()
	ts := t.Format("20060102150405")
	var result string
	for _, value := range tags {
		result += `<script type="text/javascript" src="/static/` + value + `?t=` + ts + `"></script>`
	}
	return template.HTML(result)
}

// MakeTagExternalJavascript function imports an external JavaScript library.
func MakeTagExternalJavascript(tags []string) template.HTML {
	var result string
	for _, value := range tags {
		result += `<script type="text/javascript" src="` + value + `"></script>`
	}
	return template.HTML(result)
}
