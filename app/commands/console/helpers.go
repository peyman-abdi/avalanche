package console

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

func (c *ConsoleAppImpl) Ask(question string, callback func(bool)) {
	modal := tview.NewModal()
	modal.SetBorder(true)
	modal.SetText(question)
	modal.SetBorderPadding(5, 5, 5, 5)
	modal.AddButtons([]string{"Yes", "No"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		c.Back()
		callback(buttonIndex == 0)
	})
	c.SetPage("ask", modal, false)
}

func (c *ConsoleAppImpl) Info(content string, callback func()) {
	c.dialog(content, tcell.ColorCadetBlue, callback)
}
func (c *ConsoleAppImpl) Warn(content string, callback func()) {
	c.dialog(content, tcell.ColorOrange, callback)
}
func (c *ConsoleAppImpl) Error(content string, callback func()) {
	c.dialog(content, tcell.ColorOrangeRed, callback)
}
func (c *ConsoleAppImpl) Success(content string, callback func()) {
	c.dialog(content, tcell.ColorDarkOliveGreen, callback)
}

func (c *ConsoleAppImpl) dialog(content string, backColor tcell.Color, callback func()) {
	modal := tview.NewModal()
	modal.SetBorder(true)
	modal.SetText(content)
	modal.SetBackgroundColor(backColor)
	modal.SetBorderPadding(5, 5, 5, 5)
	modal.AddButtons([]string{"OK"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		c.Back()
		callback()
	})
	c.SetPage("dialog", modal, false)
}