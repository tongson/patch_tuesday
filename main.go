package main

import (
	"embed"
	"fmt"
	"github.com/tongson/LadyLua/external/gluahttp"
	"github.com/yuin/gopher-lua"
	"github.com/tongson/LadyLua/external/gopher-json"
	"github.com/tongson/LadyLua/src"
	"os"
	"runtime"
)

//go:embed src/*
var mainSrc embed.FS

func main() {
	runtime.MemProfileRate = 0
	defer ll.RecoverPanic()
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("http", gluahttp.Xloader)
	L.PreloadModule("ll_json", json.Loader)
	ll.PatchLoader(L, "table")
	ll.PatchLoader(L, "string")
	preload := L.GetField(L.GetField(L.Get(lua.EnvironIndex), "package"), "preload")
	L.SetField(preload, "fmt", ll.LuaLoader(L, "fmt"))
	L.SetField(preload, "json", ll.LuaLoader(L, "json"))
	L.SetField(preload, "argparse", ll.LuaLoader(L, "argparse"))
	argtb := L.NewTable()
	for i := 0; i < len(os.Args); i++ {
		L.RawSet(argtb, lua.LNumber(i), lua.LString(os.Args[i]))
	}
	L.SetGlobal("arg", argtb)
	src, _ := mainSrc.ReadFile("src/main.lua")
	if err := L.DoString(string(src)); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
