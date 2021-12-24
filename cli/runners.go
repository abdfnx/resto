package cli

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "strings"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/abdfnx/resto/core/api"
	"github.com/abdfnx/resto/core/editor"
	"github.com/abdfnx/resto/core/editor/runtime"
	"github.com/abdfnx/resto/core/options"
	"github.com/abdfnx/resto/tools"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tidwall/gjson"
)

func runBasic(opts *options.CLIOptions, method string) error {
	if opts.Method.AuthType.BasicAuthUsername != "" && opts.Method.AuthType.BasicAuthPassword != "" {
		opts.Method.AuthType.Type = "basic"
	} else if opts.Method.AuthType.TokenAuth != "" {
		opts.Method.AuthType.Type = "bearer"
	}

	respone, status, requestHeaders, err := api.BasicGet(
		opts.URL,
		method,
		opts.Method.AuthType.Type,
		opts.Method.AuthType.TokenAuth,
		opts.Method.AuthType.BasicAuthUsername,
		opts.Method.AuthType.BasicAuthPassword,
		true,
		0,
		nil,
	)

	if err != nil {
		return err
	}

	if opts.Method.JustShowHeaders {
		fmt.Println(requestHeaders)
		fmt.Println("")
		fmt.Println(status)
	} else if opts.Method.JustShowBody {
		fmt.Println("")
		fmt.Println(respone)
	} else if opts.Method.SaveFile != "" {
		respone, _, _, _ := api.BasicGet(
			opts.URL,
			method,
			opts.Method.AuthType.Type,
			opts.Method.AuthType.TokenAuth,
			opts.Method.AuthType.BasicAuthUsername,
			opts.Method.AuthType.BasicAuthPassword,
			false,
			0,
			nil,
		)

		data := []byte(respone)

		err := os.WriteFile(opts.Method.SaveFile, data, 0644)

		if err != nil {
			return err
		}
	} else {
		fmt.Println(requestHeaders)
		fmt.Println("")
		fmt.Println(status)
		fmt.Println("")
		fmt.Println(string(respone))
	}

	return nil
}

func runWithBody(opts *options.CLIOptions, method string) error {
	if opts.Method.AuthType.BasicAuthUsername != "" && opts.Method.AuthType.BasicAuthPassword != "" {
		opts.Method.AuthType.Type = "basic"
	} else if opts.Method.AuthType.TokenAuth != "" {
		opts.Method.AuthType.Type = "bearer"
	}

	fn := tools.CLIRequestFile("txt")
	by := string(opts.Method.Body)
	cType := ""

	if opts.Method.ContentType != "" {
		if opts.Method.ContentType == "application/json" || opts.Method.ContentType == "json" {
			fn = tools.CLIRequestFile("json")
			cType = "application/json"
		} else if opts.Method.ContentType == "application/graphql" || opts.Method.ContentType == "graphql" {
			fn = tools.CLIRequestFile("graphql")
			cType = "application/graphql"
		} else if opts.Method.ContentType == "application/xml" || opts.Method.ContentType == "xml" {
			fn = tools.CLIRequestFile("xml")
			cType = "application/xml"
		} else if opts.Method.ContentType == "text/html" || opts.Method.ContentType == "html" {
			fn = tools.CLIRequestFile("html")
			cType = "text/html"
		} else {
			fn = tools.CLIRequestFile("txt")
			cType = "text/plain"
		}
	}

	if opts.Method.OpenEditor {
		content, err := ioutil.ReadFile(fn)
		buffer := editor.NewBufferFromString(string(content), fn)
		if err != nil {
			log.Fatalf("could not read %v: %v", fn, err)
		}

		var colorscheme editor.Colorscheme
		if railscast := runtime.Files.FindFile(editor.RTColorscheme, "railscast"); railscast != nil {
			if data, err := railscast.Data(); err == nil {
				colorscheme = editor.ParseColorscheme(string(data))
			}
		}

		bodyEditor := editor.NewView(buffer)
		bodyEditor.SetRuntimeFiles(runtime.Files)
		bodyEditor.SetColorscheme(colorscheme)

		app := tview.NewApplication()
		bodyEditor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
				case tcell.KeyCtrlS:
					tools.SaveBuffer(buffer, fn)
					app.Stop()
					return nil
			}

			return event
		})

		app.SetRoot(bodyEditor, true)

		if err := app.Run(); err != nil {
			log.Fatalf("%v", err)
		}

		b, e := os.Open(fn)

		if e != nil {
			fmt.Println(e)
		}

		defer b.Close()

		body, err := ioutil.ReadAll(b)

		if err != nil {
			panic(err)
		}

		by = string(body)
	}

	if opts.Method.IsBodyStdin {
		std, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		by = string(std)
	}

	respone, status, requestHeaders, err := api.BasicRequestWithBody(
		opts.URL,
		method,
		cType,
		by,
		opts.Method.AuthType.Type,
		opts.Method.AuthType.TokenAuth,
		opts.Method.AuthType.BasicAuthUsername,
		opts.Method.AuthType.BasicAuthPassword,
		true,
		0,
		nil,
	)

	if err != nil {
		return err
	}

	if opts.Method.JustShowHeaders {
		fmt.Println(requestHeaders)
		fmt.Println("")
		fmt.Println(status)
	} else if opts.Method.JustShowBody {
		fmt.Println("")
		fmt.Println(respone)
	} else if opts.Method.SaveFile != "" {
		respone, _, _, _ := api.BasicRequestWithBody(
			opts.URL,
			method,
			cType,
			by,
			opts.Method.AuthType.Type,
			opts.Method.AuthType.TokenAuth,
			opts.Method.AuthType.BasicAuthUsername,
			opts.Method.AuthType.BasicAuthPassword,
			false,
			0,
			nil,
		)

		data := []byte(respone)

		err := os.WriteFile(opts.Method.SaveFile, data, 0644)

		if err != nil {
			return err
		}
	} else {
		fmt.Println(requestHeaders)
		fmt.Println("")
		fmt.Println(status)
		fmt.Println("")
		fmt.Println(string(respone))
	}

	return nil
}

func runGetLatest(opts *options.GetLatestCommandOptions) error {
	registry := "github.com"

	if opts.Registry != "" {
		registry = opts.Registry
	}

	url := "https://api.github.com/repos/" + opts.Repo + "/releases/latest"

	if registry == "gitlab.com" {
		url = "https://gitlab.com/api/v4/projects/" + opts.Repo + "/repository/tags"
	} else if registry == "bitbucket.org" {
		url = "https://api.bitbucket.org/2.0/repositories/" + opts.Repo + "/refs"
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("Error creating request: %s", err.Error())
	}

	if err != nil {
		return err
	}
		
	if opts.Token != "" {
		if opts.Registry == "gitlab.com" {
			req.Header.Add("PRIVATE-TOKEN", opts.Token)
		} else {
			req.Header.Add("Authorization", "Bearer " + opts.Token)
		}
	}

	client := httpClient.HttpClient()
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("Error sending request: %s", err.Error())
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	body := string(b)

	v := gjson.Get(body, "tag_name")

	if registry == "gitlab.com" {
		value := gjson.Get(body, "#.name")
		v = gjson.Get(value.String(), "0")
	} else if registry == "bitbucket.org" {
		value := gjson.Get(body, "values")
		value2 := gjson.Get(value.String(), "0")
		v = gjson.Get(value2.String(), "name")
	}

	if v.Exists() {
		fmt.Println(v.String())
	} else {
		fmt.Println("no releases found")
	}

	return nil
}
