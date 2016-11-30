# golang-plugin
## first issue
Using a non-main package which uses stdlib package plugin doesn't work as expected with tip at the moment.

## What is the problem?
There is an error like

    16/11/30 21:49:52 plugin.Open: plugin was built with a different version of package runtime

## second issue
Using a main package as c-archive which uses stdlib package plugin doesn't seem able to load plugins.

## What is the problem?
Panic in go code like
    fatal error: runtime: no plugin module data
    
    goroutine 17 [running, locked to thread]:
    runtime.throw(0x4a9ba3, 0x1e)
            /home/erwo/go-devel/src/runtime/panic.go:596 +0x97 fp=0xc42003ca70 sp=0xc42003ca50
    plugin.lastmoduleinit(0x11a2420, 0xc42000e028, 0x11a3460, 0x26, 0x766f00)
            /home/erwo/go-devel/src/runtime/plugin.go:13 +0xc27 fp=0xc42003cba0 sp=0xc42003ca70
    plugin.open(0x4a6a5f, 0xd, 0x0, 0x0, 0x0)
            /home/erwo/go-devel/src/plugin/plugin_dlopen.go:72 +0x25f fp=0xc42003cde0 sp=0xc42003cba0
    plugin.Open(0x4a6a5f, 0xd, 0x0, 0x2, 0x3)
            /home/erwo/go-devel/src/plugin/plugin.go:30 +0x37 fp=0xc42003ce18 sp=0xc42003cde0
    main.RunGoPlugin()
            /home/erwo/golang-plugin/src/indirect_with_c/indirect_with_c.go:12 +0x52 fp=0xc42003cea8 sp=0xc42003ce18
    main._cgoexpwrap_8c4bdecea134_RunGoPlugin()
            indirect_with_c/_obj/_cgo_gotypes.go:45 +0x16 fp=0xc42003ceb0 sp=0xc42003cea8
    runtime.call32(0x0, 0x7fff5a248e28, 0x7fff5a248ecf, 0x0)
            /home/erwo/go-devel/src/runtime/asm_amd64.s:501 +0x4a fp=0xc42003cee0 sp=0xc42003ceb0
    runtime.cgocallbackg1(0x0)
            /home/erwo/go-devel/src/runtime/cgocall.go:297 +0x1a1 fp=0xc42003cf58 sp=0xc42003cee0
    runtime.cgocallbackg(0x0)
            /home/erwo/go-devel/src/runtime/cgocall.go:184 +0x86 fp=0xc42003cfc0 sp=0xc42003cf58
    runtime.cgocallback_gofunc(0x0, 0x0, 0x0, 0x0)
            /home/erwo/go-devel/src/runtime/asm_amd64.s:754 +0x71 fp=0xc42003cfe0 sp=0xc42003cfc0
    runtime.goexit()
            /home/erwo/go-devel/src/runtime/asm_amd64.s:2184 +0x1 fp=0xc42003cfe8 sp=0xc42003cfe0
    
    goroutine 18 [syscall, locked to thread]:
    runtime.goexit()
            /home/erwo/go-devel/src/runtime/asm_amd64.s:2184 +0x1
