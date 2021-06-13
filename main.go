package main

import (
	"embed"
	"github.com/tongson/LadyLua/external/gluahttp"
	"github.com/yuin/gopher-lua"
	"github.com/tongson/LadyLua/external/gopher-json"
	"github.com/tongson/LadyLua/src"
	"os"
	"runtime"
)

//go:embed main/*
var mainSrc embed.FS

//go:embed src/*
var luaSrc embed.FS

func main() {
	runtime.MemProfileRate = 0
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("http", gluahttp.Xloader)
	L.PreloadModule("ll_json", json.Loader)
	ll.PatchLoader(L, "table")
	ll.PatchLoader(L, "string")
	ll.EmbedLoader(L)
	cvrf, _ := luaSrc.ReadFile("src/cvrf.lua")
	ll.ModuleLoader(L, "cvrf", string(cvrf))
	argtb := L.NewTable()
	for i := 0; i < len(os.Args); i++ {
		L.RawSet(argtb, lua.LNumber(i), lua.LString(os.Args[i]))
	}
	L.SetGlobal("arg", argtb)
	src, _ := mainSrc.ReadFile("main/main.lua")
	ll.MainLoader(L, src)
	os.Exit(0)
}
