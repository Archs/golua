package main

import (
	"flag"
	"github.com/Archs/go-iup/iup"
	"github.com/Archs/golua/lua"
)

func adder(L *lua.State) int {
	a := L.ToInteger(1)
	b := L.ToInteger(2)
	L.PushInteger(int64(a + b))
	return 1 // number of return values
}

func goprint(L *lua.State) int {
	a := L.ToString(1)
	println(a)
	return 0 // number of return values
}

func goinit(L *lua.State) int {
	for i, v := range iup.GetAllNames() {
		println(i, v)
	}
	mh := iup.GetHandle("menu")
	if mh == nil {
		println("Unable to get handle")
	} else {
		println("Parent Name:", iup.GetName(mh.(*iup.Handle).GetParent()))
		mh.(*iup.Handle).Popup(iup.MOUSEPOS, iup.MOUSEPOS)
	}
	btn := iup.GetHandle("btn").(*iup.Handle)
	if btn == nil {
		println("btn not found")
		return 0
	}
	btn.SetCallback(func(arg *iup.ButtonAction) {
		println("btn called (with callback set in go)")
	})
	return 0 // number of return values
}

func main() {
	flag.Parse()
	println("begin")
	l := lua.NewState()
	defer l.Close()
	l.OpenLibs()
	l.OpenIUP()
	l.Register("adder", adder)
	l.Register("goinit", goinit)
	l.Register("goprint", goprint)
	l.DoString("print(adder(2, 2))")
	l.DoString("goprint(adder(3, 3))")
	// l.OpenIUP()
	println("opening:", flag.Arg(0))
	err := l.DoFile(flag.Arg(0))
	if err != nil {
		panic(err.Error())
	}
}
