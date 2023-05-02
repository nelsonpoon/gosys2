package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	//"strconv"
)

var a fyne.App

func createApp() {
	//a := app.New()
	a=app.New()
}

// Hello returns a greeting for the named person.
func WelcomeWin(a fyne.App) {
    // Return a greeting that embeds the name in a message.
    //message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//a := app.New()
	w := a.NewWindow("Welcome Windows")

	hello := widget.NewLabel("Welcome to gosys2!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome to gosys2!")
		}),
	))

	w.ShowAndRun()

}
/*
func LoginWin(a *fyne.App) {
    // Return a greeting that embeds the name in a message.
    //message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//a := app.New()
	w := a.NewWindow("Welcome to Login Window")

	hello := widget.NewLabel("Welcome to login!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome to login!")
		}),
	))

	w.ShowAndRun()

}*/