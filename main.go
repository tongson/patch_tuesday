package main

import (
	"embed"
	"os"

	"github.com/tongson/LadyLua"
	"github.com/yuin/gopher-lua"
)

//go:embed main/*
var mainSrc embed.FS

//go:embed src/*
var luaSrc embed.FS

func main() {
	L := lua.NewState()
	defer L.Close()

	// See for available modules -> https://github.com/tongson/LadyLua#modules
	// Also:
	// https://github.com/tongson/LadyLua/blob/main/docs/go-loader.adoc
	// https://github.com/tongson/LadyLua/blob/main/docs/go-helper.adoc

	// Load `http` and `json` modules
	ll.PreloadGo(L, "http")
	ll.PreloadGo(L, "json")
	ll.LoadGlobalGo(L, "extend")

	// Allow loading(require) Lua code from LadyLua; found in `LadyLua/internal/lua`
	ll.Preload(L)

	// Load Lua source from `src`; for `require("cvrf")`
	// Usually modules specific to a project or program
	// Depends on the go:embed directive, any directory or filename works
	ll.PreloadModule(L, "cvrf", ll.ReadFile(luaSrc, "src/cvrf.lua"))

	// Capture command line arguments
	ll.FillArg(L, os.Args)

	// Load Lua source from `main`; the entrypoint Lua code or so-called main()
	// Depends on the go:embed directive, any directory or filename works
	ll.Main(L, ll.ReadFile(mainSrc, "main/main.lua"))

	// If all goes well; Lua code can call `os.exit` to override exit code
	os.Exit(0)
}
