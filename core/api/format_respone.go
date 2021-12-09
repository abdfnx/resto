package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tidwall/pretty"
	"github.com/yosssi/gohtml"
)

var (
	statusTitle     = "Status"
	statusCodeTitle = "Status Code"
	statusRowHeader = table.Row{statusTitle, statusCodeTitle}
)

// formatResponse formats the Response with Indents and Colors
func formatResponse(resp *http.Response) (string, string, string, error) {
	heads := fmt.Sprint(resp.Header)
	toReturn := ""

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	str := string(body)

	statusTable := table.NewWriter()
	statusTable.AppendHeader(statusRowHeader)
	statusTable.AppendRow(table.Row{resp.Status, resp.StatusCode})
	statusTable.SetStyle(table.StyleRounded)

	headersTable := table.NewWriter()

	for key, value := range resp.Header {
		v1 := strings.ReplaceAll(value[0], "[", "")
		v2 := strings.ReplaceAll(v1, "]", "")

		if len(v2) > 80 {
			v2 = v2[:80] + "..."
		}

		headersTable.AppendRow(table.Row{key, v2})
		headersTable.SetStyle(table.StyleRounded)
		headersTable.Style().Options.DrawBorder = true
		headersTable.Style().Options.SeparateRows = true
	}

	if strings.Contains(heads, "json") {
		toReturn = string(pretty.Pretty([]byte(str)))
	} else if strings.Contains(heads, "xml") || strings.Contains(heads, "html") || strings.Contains(heads, "plain") {
		var s string

		if strings.Contains(heads, "plain") {
			s = str
		} else {
			s = gohtml.Format(str)
		}

		toReturn = s
	}

	return toReturn, statusTable.Render(), headersTable.Render(), nil
}
