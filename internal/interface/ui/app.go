package ui

import (
	"email-client/internal/domain/model"
	"email-client/internal/interface/controller"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

func WithBackButton(app *tview.Application, backTo tview.Primitive, content tview.Primitive) tview.Primitive {
	backButton := tview.NewButton("← Press esc").
		SetSelectedFunc(func() {
			app.SetRoot(backTo, true)
		})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(backButton, 1, 0, false).
		AddItem(content, 0, 1, true)

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

	detail := tview.NewTextView().SetDynamicColors(true).SetWrap(true)
	list := tview.NewList()

	setupForm(app, handler, list, detail)
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			app.SetRoot(flex, true)
			return nil
		}
		return event
	})

	menu := tview.NewList().
		AddItem("Inbox", "View your inbox", 'i', func() {
			inbox := handler.GetInbox()
			list.Clear()
			for _, email := range inbox {
				email := email // capture the loop variable
				list.AddItem(fmt.Sprintf("[%s] %s", email.From, email.Subject), "", 0, func() {
					detail.SetText(fmt.Sprintf("[::b]From:[-:-] %s\n[::b]To:[-:-] %s\n[::b]Subject:[-:-] %s\n\n%s",
						email.From, email.To, email.Subject, email.Body))
					app.SetRoot(WithBackButton(app, list, detail), true)
				})
			}
			app.SetRoot(list, true)
		}).
		AddItem("Compose", "Write a new email", 'c', func() {
			newForm := setupForm(app, handler, list, detail)
			app.SetRoot(WithBackButton(app, flex, newForm), true)
		}).AddItem("🗑️ Delete", "Choose an email to delete", 'd', func() {
		inbox := handler.GetInbox()
		deleteList := tview.NewList()

		for _, email := range inbox {
			email := email
			deleteList.AddItem(fmt.Sprintf("[%s] %s", email.From, email.Subject), "", 0, func() {
				handler.Delete(email.ID)

				inbox := handler.GetInbox()
				list.Clear()
				for _, e := range inbox {
					e := e
					list.AddItem(fmt.Sprintf("[%s] %s", e.From, e.Subject), "", 0, func() {
						detail.SetText(fmt.Sprintf("[::b]From:[-:-] %s\n[::b]To:[-:-] %s\n[::b]Subject:[-:-] %s\n\n%s",
							e.From, e.To, e.Subject, e.Body))
						app.SetRoot(WithBackButton(app, list, detail), true)
					})
				}

				modal := tview.NewModal().
					SetText("✅ Email deleted successfully!").
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, label string) {
						app.SetRoot(flex, true)
					})
				app.SetRoot(modal, true)
			})
		}

		deleteList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEsc {
				app.SetRoot(flex, true)
				return nil
			}
			return event
		})

		app.SetRoot(WithBackButton(app, flex, deleteList), true)
	}).
		AddItem("Quit", "Exit app", 'q', func() {
			app.Stop()
		})

	flex.AddItem(menu, 30, 1, true)

	app.SetRoot(flex, true)
	app.Run()
}

func setupForm(app *tview.Application, handler *controller.Handler, list *tview.List, detail *tview.TextView) *tview.Form {
	form := tview.NewForm()
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown:
			return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
		case tcell.KeyUp:
			return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModShift)
		}
		return event
	})
	form.
		AddInputField("To", "", 40, nil, nil).
		AddInputField("Subject", "", 40, nil, nil).
		AddInputField("Body", "", 40, nil, nil).
		AddButton("Send", func() {
			to := form.GetFormItemByLabel("To").(*tview.InputField).GetText()
			subject := form.GetFormItemByLabel("Subject").(*tview.InputField).GetText()
			body := form.GetFormItemByLabel("Body").(*tview.InputField).GetText()

			if to == "" || subject == "" || body == "" {
				detail.SetText("[red]All fields are required.")
				app.SetRoot(WithBackButton(app, list, detail), true)
				return
			}
			if !strings.Contains(to, "@") || !strings.Contains(to, ".") {
				detail.SetText("[red]Invalid email address format.")
				app.SetRoot(WithBackButton(app, list, detail), true)
				return
			}

			handler.Send(model.Email{
				From:    "me@example.com",
				To:      to,
				Subject: subject,
				Body:    body,
			})

			form.GetFormItemByLabel("To").(*tview.InputField).SetText("")
			form.GetFormItemByLabel("Subject").(*tview.InputField).SetText("")
			form.GetFormItemByLabel("Body").(*tview.InputField).SetText("")

			inbox := handler.GetInbox()
			list.Clear()
			for _, email := range inbox {
				email := email
				list.AddItem(fmt.Sprintf("[%s] %s", email.From, email.Subject), "", 0, func() {
					detail.SetText(fmt.Sprintf("[::b]From:[-:-] %s\n[::b]To:[-:-] %s\n[::b]Subject:[-:-] %s\n\n%s",
						email.From, email.To, email.Subject, email.Body))
					app.SetRoot(WithBackButton(app, list, detail), true)
				})
			}
			app.SetRoot(list, true)
		}).
		AddButton("Cancel", func() {
			form.GetFormItemByLabel("To").(*tview.InputField).SetText("")
			form.GetFormItemByLabel("Subject").(*tview.InputField).SetText("")
			form.GetFormItemByLabel("Body").(*tview.InputField).SetText("")
			app.SetRoot(list, true)
		})

	return form
}
