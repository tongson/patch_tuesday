package main

import (
	"embed"
	"github.com/yuin/gopher-lua"
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
	// Load `http` and `json` modules
	ll.GoLoader(L, "http")
	ll.GoLoader(L, "json")
	// Load lua code to patch `table` and `string`; found in `LadyLua/src/lua`
	ll.PatchLoader(L, "table")
	ll.PatchLoader(L, "string")
	// Load all plain Lua code from LadyLua; found in `LadyLua/src/lua`
	ll.EmbedLoader(L)
	// Load Lua source from `src`; for `require("cvrf")`
	ll.ModuleLoader(L, "cvrf", ll.ReadFile(luaSrc, "src/cvrf.lua"))
	// Capture command line arguments
	ll.FillArg(L, os.Args)
	// Load Lua source from `main`; the main() Lua code
	ll.MainLoader(L, ll.ReadFile(mainSrc, "main/main.lua"))
	os.Exit(0)
}
