package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/abdfnx/resto/validation"

	"github.com/abdfnx/resto/core/graphql"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tidwall/pretty"
	"github.com/rivo/tview"
)

var (
	status     string
	statusCode string
)

// BasicRequestWithBody sends put|patch|post|delete requests
func BasicRequestWithBody(
		httpURL,
		method,
		contentType,
		reqBody,
		authType,
		baererToken,
		basicAuthUsername,
		basicAuthPassword string,
		isCommand bool,
		headersCount int,
		headersForm *tview.Form,
	) (string, string, string, error) {
	url, err := validation.CheckURL(httpURL)

	if err != nil {
		return "", "", "", err
	}

	if contentType == "application/graphql" {
		// create a client (safe to share across requests)
		httpclient := &http.Client{}
		client := graphql.NewClient(url, graphql.WithHTTPClient(httpclient))

		// make a request
		req := graphql.NewRequest(reqBody)

		if authType == "bearer" {
			req.Header.Set("Authorization", "Bearer " + baererToken)
		} else if authType == "basic" {
			req.Header.Set("Authorization", "Basic " + basicAuth(basicAuthUsername, basicAuthPassword))
		}

		for i := 0; i < headersCount; i++ {
			key := headersForm.GetFormItem(i).(*tview.InputField).GetLabel()
			value := headersForm.GetFormItem(i).(*tview.InputField).GetText()
			
			req.Header.Set(key, value)
		}

		// define a Context for the request
		ctx := context.Background()

		// run it and capture the response
		var respData map[string]interface{}

		client.Run(ctx, req, &respData)

		jsonString, err := json.Marshal(respData)
		data := `{ "data": ` + string(jsonString) + `}`
		prettyData := string(pretty.Pretty([]byte(data)))

		if err != nil {
			panic(err)
		}

		statusTable := table.NewWriter()

		if string(jsonString) != "null" {
			status = "200 OK"
			statusCode = "200"
		} else if string(jsonString) == "null" {
			status = "404 Not Found"
			statusCode = "404"
		} else {
			status = "500 Internal Server Error"
			statusCode = "500"
		}

		statusTable.AppendHeader(statusRowHeader)
		statusTable.AppendRow(table.Row{status, statusCode})
		statusTable.SetStyle(table.StyleRounded)

		if string(jsonString) != "null" {
			if isCommand {
				colored := string(pretty.Color([]byte(prettyData), nil))
				return colored, statusTable.Render(), " ", err
			} else {
				return prettyData, statusTable.Render(), " ", err
			}
		} else if string(jsonString) == "null" {
			return string(jsonString), statusTable.Render(), "", err
		} else {
			return string(jsonString), statusTable.Render(), " ", err
		}
	} else {
		var jsonStr = []byte(reqBody)
		req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

		if err != nil {
			return "", "", "", fmt.Errorf("Error creating request: %s", err.Error())
		}

		req.Header.Set("Content-Type", contentType)
		
		if authType == "bearer" {
			req.Header.Set("Authorization", "Bearer " + baererToken)
		} else if authType == "basic" {
			req.Header.Set("Authorization", "Basic " + basicAuth(basicAuthUsername, basicAuthPassword))
		}

		if headersForm != nil {
			for i := 0; i < headersCount; i++ {
				key := headersForm.GetFormItem(i).(*tview.InputField).GetLabel()
				value := headersForm.GetFormItem(i).(*tview.InputField).GetText()
				
				req.Header.Set(key, value)
			}
		}

		client := httpClient.HttpClient()

		res, err := client.Do(req)

		if err != nil {
			return "", "", "", fmt.Errorf("Error sending request: %s", err.Error())
		}

		defer res.Body.Close()

		return formatResponse(res, isCommand)
	}
}
