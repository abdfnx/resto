package cli

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"

	"github.com/abdfnx/resto/tools"
	"github.com/abdfnx/resto/core/api"
	"github.com/abdfnx/resto/core/editor"
	"github.com/abdfnx/resto/core/editor/runtime"
	"github.com/abdfnx/resto/core/options"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
