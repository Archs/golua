package main

import (
	"fmt"
	"github.com/Archs/go-iup/iup"
	"time"
)

func autoHide(h *iup.Handle) {
	shown := false
	i := 0
	go func() {
		println("begin autohide of", h)
		for {
			time.Sleep(2 * time.Second)
			if shown {
				h.Hide()
				shown = false
			} else {
				h.Show()
				shown = true
			}
			println(h, i)
			i = i + 1
		}
	}()
}

func autoAppend(txt *iup.Handle) {
	i := 0
	go func() {
		for {
			s := fmt.Sprintln(i)
			txt.SetAttribute("INSERT", s)
			println("inserting ", s, "into:", txt)
			i = i + 1
			time.Sleep(time.Second)
		}
	}()
}

func newDialog() *iup.Handle {
	btn := iup.Button("TITLE=clickme")
	txt := iup.Text("MULTILINE=YES", "EXPAND=YES")
	dlg := iup.Dialog("TITLE=newDialog", "SIZE=halfxhalf",
		iup.Vbox(
			btn,
			txt,
		),
		"SHRINK=YES",
	)
	btn.SetCallback(func(arg *iup.ButtonAction) {
		println(dlg, btn, "clicked")
	})
	autoAppend(txt)
	return dlg
}

func main() {
	iup.Open()
	dlg := iup.Dialog(
		iup.Vbox(
			iup.Frame(
				"TITLE=BTN GROUD",
				iup.Hbox(
					iup.Button(
						"TITLE=btn",
						func(arg *iup.ButtonAction) {
							println("btn clicked")
						},
					),
					iup.Button(
						"TITLE=&New Dialog",
						func(arg *iup.ButtonAction) {
							d := newDialog()
							d.Show()
						},
					),
				),
			),
		),
		"SIZE=200x200",
		"SHRINK=YES",
	)
	dlg.Show()
	iup.MainLoop()
}
