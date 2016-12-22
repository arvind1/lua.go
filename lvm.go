// Copyright (c) 2016 Arvind Singh
// All rights reserved.
// Use of this source code is governed by a MIT License
// that can be found in LICENSE file

package lua

/*
#cgo pkg-config: lua53

#include <stdlib.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type LuaVM C.lua_State

func New() *LuaVM {
	return (*LuaVM)(C.luaL_newstate())
}

func (L *LuaVM) OpenLibs() {
	C.luaL_openlibs(L)
}

func (L *LuaVM) Register(name string, fn unsafe.Pointer) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.lua_pushcclosure(L, C.lua_CFunction(fn), 0)
	C.lua_setglobal(L, cstr)
}

func (L *LuaVM) DoString(code string) error {
	cstr := C.CString(code)
	defer C.free(unsafe.Pointer(cstr))

	if err := C.luaL_loadstring(L, cstr); err != C.LUA_OK {
		C.lua_settop(L, -(1)-1)
		return errors.New(fmt.Sprint(err))
	}

	if err := C.lua_pcallk(L, 0, 0, 0, 0, nil); err != C.LUA_OK {
		return errors.New(fmt.Sprint(err))
	}

	return nil
}
