package main

import (
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

var searchResultsAlertTemplate *template.Template

func init() {
	var err error

	if searchResultsAlertTemplate, err = parseTemplate(searchResultsAlertTemplateContent); err != nil {
		// This shouldn't fail, since we control the template content via a
		// constant below.
		panic(err)
	}
}

type searchResultsAlert struct {
	Title           string
	Description     string
	ProposedQueries []struct {
		Description string
		Query       string
	}
}

func (alert *searchResultsAlert) HasAlert() bool {
	return alert.Title != ""
}

func (alert *searchResultsAlert) Render() (string, error) {
	b := &strings.Builder{}
	if err := searchResultsAlertTemplate.Execute(b, alert); err != nil {
		return "", errors.Wrap(err, "rendering alert template")
	}
	return b.String(), nil
}

const searchResultsAlertTemplateContent = `
{{- color "search-alert-title"}}❗{{.Title}}{{color "nc" -}}{{"\n"}}
{{- if gt (len .Description) 0 -}}
	{{- color "search-alert-description"}}{{.Description}}{{"\n"}}{{color "nc" -}}
{{- end -}}
{{- if gt (len .ProposedQueries) 0 -}}
	{{- color "search-alert-proposed-title"}}Did you mean:{{color "nc" -}}
	{{- "\n" -}}
	{{- range .ProposedQueries -}}
		{{- color "search-alert-proposed-query"}}{{.Query}}{{color "nc" -}}
		{{- " - " -}}
		{{- color "search-alert-proposed-description"}}{{.Description}}{{color "nc" -}}
		{{- "\n" -}}
	{{- end -}}
{{- end -}}
`
