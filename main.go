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
	ll.ModuleLoader(L, "cvrf", ll.ReadFile(luaSrc, "src/cvrf.lua"))
	ll.FillArg(L, os.Args)
	ll.MainLoader(L, ll.ReadFile(mainSrc, "main/main.lua"))
	os.Exit(0)
}
