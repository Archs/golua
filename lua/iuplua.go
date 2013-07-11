package lua

/*
#cgo CFLAGS: -Ilua
#cgo LDFLAGS: -laiox

#include "lua.h"
#include "iup.h"
#include "iuplua.h"
*/
import "C"

func (L *State) OpenIUP() {
	C.iuplua_open(L.s)
	C.iupkey_open(L.s)
}

// // add iuplua
// //   other libs
// iuplua_open(L);
// iupcontrolslua_open(L);
// iupkey_open(L);
// // iuplua_close(L);
// // iup_pplotlua_open(L);
// // cd
// cdlua_open(L);
// cdluaiup_open(L);
// // im
// iupimlua_open(L);
// imlua_open(L);
// imlua_open_process(L);
// // gl
// iupolelua_open(L);
// iupgllua_open(L);
// // end other libs
