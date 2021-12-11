package layout

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"

	"github.com/abdfnx/resto/tools"
	"github.com/abdfnx/resto/core/api"
	"github.com/abdfnx/resto/core/editor"
	"github.com/abdfnx/resto/core/editor/runtime"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tidwall/pretty"
)

var (
	// request
	method          string
	httpURL         string
	cType           string
	fn              string = tools.RequestFile()

	// respone
	body    	    string
	respone 	    string
	status  	    string

	// auth
	authType 	    string
	requestHeaders  string
)

func Layout() {
	app := tview.NewApplication()
	flex := tview.NewFlex()
	helpPage := tview.NewPages()
	helpText := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	help := `
		Welcome to Resto!

		resto is a cli app can send pretty HTTP & API requests from your terminal.

		Shortcuts:
			- Ctrl+P: Open Renio Panel
			- Ctrl+H: Open Help Guide
			- Ctrl+S: Save Request Body
			- Ctrl+Q: Quit			
	`

	fmt.Fprintf(helpText, "%s ", help)

	helpPage.AddAndSwitchToPage("help", tview.NewGrid().
		SetColumns(30, 0, 30).
		SetRows(3, 0, 3).
		AddItem(helpText, 1, 1, 1, 1, 0, 0, true), true).
	ShowPage("main")

	// forms
	authForm := tview.NewForm()
	panelForm := tview.NewForm()
	requestForm := tview.NewForm()

	// request inputs
	urlField := tview.NewInputField().
		SetLabel("URL").
		SetFieldWidth(32).
		SetPlaceholder("URL")
	
	requestMethods := tview.NewDropDown().
		SetLabel("Request Method").
		SetOptions([]string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"HEAD",
		}, func(option string, optionIndex int) {
			method = option
		}).SetCurrentOption(0)

	contentType := tview.NewDropDown().
		SetLabel("Content Type").
		SetOptions([]string{
			"none",
			"application/json",
			"application/graphql",
			"application/xml",
			"text/html",
			"text/plain",
		}, func(option string, optionIndex int) {
			cType = option
		}).SetCurrentOption(0)

	// request body
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
	bodyEditor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyCtrlS:
				tools.SaveBuffer(buffer, fn)
				app.SetRoot(flex, true).SetFocus(requestForm)
				return nil
		}

		return event
	})

	// response outputs
	responseView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	statusView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	// headers inputs
	headers := tview.NewTextView()

	// auth inputs
	token := tview.NewInputField().
		SetLabel("Token").
		SetFieldWidth(20) 

	username := tview.NewInputField().
		SetLabel("Username").
		SetFieldWidth(20)

	password := tview.NewInputField().
		SetLabel("Password").
		SetFieldWidth(20)

	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(requestForm, 0, 1, false).
		AddItem(authForm, 20, 1, false).
		AddItem(panelForm, 15, 1, false), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(responseView, 0, 3, false).
		AddItem(statusView, 7, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 0, false)
	
	var send = func() {
		responseView.Clear()
		statusView.Clear()

		httpURL = urlField.GetText()

		b, e := os.Open(fn)

		if e != nil {
			fmt.Println(e)
		}

		defer b.Close()

		currentBody, err := ioutil.ReadAll(b)

		if err != nil {
			panic(err)
		}

		if cType == "application/json" {
			fn = "reqBody.json"
		} else if cType == "application/graphql" {
			fn = "reqBody.gql"
		} else if cType == "application/xml" {
			fn = "reqBody.xml"
		} else if cType == "text/html" {
			fn = "reqBody.html"
		} else {
			fn = "reqBody.txt"
		}

		if cType == "application/json" {
			var r map[string]interface{}
			json.Unmarshal([]byte(currentBody), &r)
			body = string(pretty.Pretty([]byte(currentBody)))
		} else {
			body = string(currentBody)
		}

		if method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE" {
			respone, status, requestHeaders, _ =
				api.BasicRequestWithBody(
					httpURL,
					method,
					cType,
					body,
					authType,
					token.GetText(),
					username.GetText(),
					password.GetText(),
					false,
				)
		} else {
			body = ""

			respone, status, requestHeaders, _ =
				api.BasicGet(
					httpURL,
					method,
					authType,
					token.GetText(),
					username.GetText(),
					password.GetText(),
					false,
				)
		}

		headers.Clear()
		requestHeaders += "\n\nTo Exit Press 'Esc' Key"

		fmt.Fprintf(responseView, "%s ", respone)
		fmt.Fprintf(statusView, "%s ", status)
		fmt.Fprintf(headers, "%s", requestHeaders)
	}

	requestForm.AddFormItem(requestMethods).
		AddFormItem(urlField).
		AddFormItem(contentType).
		AddButton("Panel", func() {
			app.SetRoot(flex, true).SetFocus(panelForm)
		}).
		AddButton("Body", func() {
			app.SetRoot(bodyEditor, true).SetFocus(bodyEditor).Run()
		}).
		AddButton("Authorization", func() {
			app.SetRoot(flex, true).SetFocus(authForm)
		}).
		AddButton("Send", func() {
			send()

			app.SetRoot(flex, true).SetFocus(responseView)
		})

	responseView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			app.SetRoot(flex, true).SetFocus(requestForm)
		}
	})

	headers.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			app.SetRoot(flex, true).SetFocus(requestForm)
		}
	})

	var Input = func(text, label string, width int, doneFunc func(text string)) {
		fileNameInput := tview.NewPages()

		input := tview.NewInputField().SetText(text)
		input.SetBorder(true)
		input.SetLabel(label).SetLabelWidth(width).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				doneFunc(input.GetText())
				fileNameInput.RemovePage("input")
			}
		})

		input.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEsc {
				app.SetRoot(flex, true).SetFocus(requestForm)
			}
		})

		fileNameInput.AddAndSwitchToPage("input", tview.NewGrid().
			SetColumns(0, 0, 0).
			SetRows(0, 3, 0).
			AddItem(input, 1, 1, 1, 1, 0, 0, true), true).ShowPage("main")
		
		app.SetRoot(fileNameInput, true).SetFocus(input)
	}

	panelModel := tview.NewModal().
		SetText("What task do you want to do?").
		AddButtons([]string{
			"Request Form",
			"Send Request",
			"Body",
			"Authorization",
			"Show Response Headers",
			"Save Response in File",
			"Return",
			"Quit From App",
		}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
				case "Request Form":
					app.SetRoot(flex, true).SetFocus(requestForm)

				case "Send Request":
					send()
					app.SetRoot(flex, true).SetFocus(requestForm)

				case "Body":
					app.SetRoot(bodyEditor, true).SetFocus(bodyEditor)
				
				case "Authorization":
					app.SetRoot(flex, true).SetFocus(authForm)

				case "Show Response Headers":
					app.SetRoot(headers, true).SetFocus(headers)
				
				case "Save Response in File":
					data := []byte(respone)

					Input("response.json", "file name", 5, func(fn string) {
						err := os.WriteFile(fn, data, 0644)

						if err != nil {
							panic(err)
						}

						app.SetRoot(flex, true).SetFocus(requestForm)
					})
				
				case "Return":
					app.SetRoot(flex, true).SetFocus(requestForm)
				
				case "Quit From App":
					app.Stop()
			}
		})

	panelForm.AddButton("Show Headers", func() {
		app.SetRoot(headers, true).SetFocus(headers)
	}).AddButton("Open Panel", func() {
		app.SetRoot(panelModel, true).SetFocus(panelModel)
	}).AddButton("Exit", func() {
		app.Stop()
	})

	authForm.AddDropDown("Authentication Type", []string{"none", "basic auth", "bearer token"}, 0, func(option string, optionIndex int) {
		tokenIndex := authForm.GetFormItemIndex("Token")
		usernameIndex := authForm.GetFormItemIndex("Username")
		passwordIndex := authForm.GetFormItemIndex("Password")

		if option == "basic auth" {
			if tokenIndex != -1 {
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Token"))
			} else if usernameIndex != -1 && passwordIndex != -1 {
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Username"))
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Password"))
			}

			authForm.AddFormItem(username)
			authForm.AddFormItem(password)

			authType = "basic"
		} else if option == "bearer token" {
			if usernameIndex != -1 && passwordIndex != -1 {
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Username"))
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Password"))
			} else if tokenIndex != -1 {
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Token"))
			}

			authForm.AddFormItem(token)

			authType = "bearer"
		} else {
			if usernameIndex != -1 && passwordIndex != -1 {
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Username"))
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Password"))
			}

			if tokenIndex != -1 {
				authForm.RemoveFormItem(authForm.GetFormItemIndex("Token"))

			}

			token.SetText("")
			username.SetText("")
			password.SetText("")
		}
	})

	authForm.AddButton("Panel", func() {
		app.SetRoot(flex, true).SetFocus(panelForm)
	}).AddButton("Request", func() {
		app.SetRoot(flex, true).SetFocus(requestForm)
	})

	helpText.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			app.SetRoot(flex, true).SetFocus(requestForm)
		}
	})

	// set borders
	authForm.SetBorder(true)
	panelForm.SetBorder(true)
	requestForm.SetBorder(true)
	responseView.SetBorder(true)
	statusView.SetBorder(true)

	if err := app.
		EnableMouse(true).
		SetRoot(flex, true).
		SetFocus(requestForm).
		Sync().
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
				case tcell.KeyCtrlP:
					app.SetRoot(panelModel, true).SetFocus(panelModel)
					return nil

				case tcell.KeyCtrlH:
					app.SetRoot(helpPage, true).SetFocus(helpPage)
					return nil

				case tcell.KeyCtrlQ:
					app.Stop()
					return nil
			}

			return event
		}).
		Run();
	err != nil {
		panic(err)
	}
}
