module patch_tuesday

go 1.16

require (
	github.com/tongson/LadyLua v0.0.0-20210613112916-1737fa6d707d
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9
)

replace github.com/yuin/gopher-lua => github.com/tongson/gopher-lua v0.0.0-20210610051759-53ab9600e09f