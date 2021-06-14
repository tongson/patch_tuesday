package main

import (
	"embed"
	"github.com/tongson/LadyLua"
	"github.com/yuin/gopher-lua"
	"os"
)

//go:embed main/*
var mainSrc embed.FS

//go:embed src/*
var luaSrc embed.FS

func main() {
	L := lua.NewState()
	defer L.Close()

	// See for available modules -> https://github.com/tongson/LadyLua#modules
	// Load `http` and `json` modules
	ll.GoLoader(L, "http")
	ll.GoLoader(L, "json")

	// Load lua code to patch `table` and `string`; found in `LadyLua/src/lua`
	ll.PatchLoader(L, "table")
	ll.PatchLoader(L, "string")

	// Allow loading(require) Lua code from LadyLua; found in `LadyLua/src/lua`
	ll.EmbedLoader(L)

	// Load Lua source from `src`; for `require("cvrf")`
	// Usually modules specific to a project or program
	// Depends on the go:embed directive, any directory or filename works
	ll.ModuleLoader(L, "cvrf", ll.ReadFile(luaSrc, "src/cvrf.lua"))

	// Capture command line arguments
	ll.FillArg(L, os.Args)

	// Load Lua source from `main`; the entrypoint Lua code or so-called main()
	// Depends on the go:embed directive, any directory or filename works
	ll.MainLoader(L, ll.ReadFile(mainSrc, "main/main.lua"))

	// If all goes well; Lua code can call `os.exit` to override exit code
	os.Exit(0)
}
