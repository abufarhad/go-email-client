package ui

import (
	"email-client/internal/controller"
	"email-client/internal/model"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func WithBackButton(app *tview.Application, backTo tview.Primitive, content tview.Primitive) tview.Primitive {
	backButton := tview.NewButton("← Press esc").
		SetSelectedFunc(func() {
			app.SetRoot(backTo, true)
		})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(backButton, 1, 0, false).
		AddItem(content, 0, 1, true)

	// Optional: handle Esc key as well
	layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			app.SetRoot(backTo, true)
			return nil
		}
		return event
	})

	return layout
}

func StartApp(handler *controller.Handler) {
	app := tview.NewApplication()

	list := tview.NewList()
	detail := tview.NewTextView().SetDynamicColors(true).SetWrap(true)
	form := tview.NewForm()

	// Prepare inbox
	inbox := handler.GetInbox()
	for _, email := range inbox {
		list.AddItem(fmt.Sprintf("[%s] %s", email.From, email.Subject), "", 0, nil)
	}

	var flex *tview.Flex

	// Menu (uses flex in callback)
	menu := tview.NewList().
		AddItem("Inbox", "View your inbox", 'i', func() {
			app.SetRoot(list, true)
		}).
		AddItem("Compose", "Write a new email", 'c', func() {
			app.SetRoot(WithBackButton(app, flex, form), true)
		}).
		AddItem("Quit", "Exit app", 'q', func() {
			app.Stop()
		})

	flex = tview.NewFlex().
		AddItem(menu, 30, 1, true).
		AddItem(list, 0, 2, false)

	// View email details
	list.SetSelectedFunc(func(i int, mainText, secondaryText string, shortcut rune) {
		email := inbox[i]
		detail.SetText(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
			email.From, email.To, email.Subject, email.Body))
		app.SetRoot(WithBackButton(app, flex, detail), true)
	})

	// Compose form
	form.
		AddInputField("To", "", 40, nil, nil).
		AddInputField("Subject", "", 40, nil, nil).
		AddInputField("Body", "", 40, nil, nil).
		AddButton("Send", func() {
			to := form.GetFormItemByLabel("To").(*tview.InputField).GetText()
			subject := form.GetFormItemByLabel("Subject").(*tview.InputField).GetText()
			body := form.GetFormItemByLabel("Body").(*tview.InputField).GetText()
			email := model.Email{
				From:    "me@example.com",
				To:      to,
				Subject: subject,
				Body:    body,
			}
			handler.Send(email)
			app.SetRoot(flex, true)
		})

	// Start app with main layout
	app.SetRoot(flex, true)
	app.Run()
}
