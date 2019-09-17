package main

import (
    "flag"
    "fmt"
    "syscall"
    "unsafe"
)

// flag variable
var (
    help  bool   // help info
    def   bool   // default mode ,print \a
    box   bool   // message box mode
    mybox string // custom message box info
)

// syscall variable
var (
    libUser32       = syscall.NewLazyDLL("user32.dll")
    procMessageBoxW = libUser32.NewProc("MessageBoxW")
)

const (
    MB_OK = 0x00000000
)

// init function to parse flag
func init() {
    flag.BoolVar(&def, "d", true, "default mode,print BEL.")
    flag.BoolVar(&help, "h", false, "help info.")
    flag.BoolVar(&box, "b", false, "box mode,popup a message box.")
    flag.StringVar(&mybox, "m", "", "custom mode,you can custom message box text.")
}

func main() {
    flag.Parse()
    switch {
    case help:
        flag.Usage()
    case mybox != "":
        customMode(mybox)
    case box:
        boxMode()
    case def:
        defaultMode()
    }
}

// defaultMode
func defaultMode() {
    fmt.Printf("%v\a", "ticker complete.")
}

// boxMode
func boxMode() {
    MessageBoxW("ticker", "ticker complete", MB_OK)
}

// customMode
func customMode(text string) {
    MessageBoxW("ticker", text, MB_OK)
}

// system call MessageBoxW(windows api)
func MessageBoxW(caption, text string, style uintptr) (result uint32) {
    text_, _ := syscall.UTF16PtrFromString(text)
    caption_, _ := syscall.UTF16PtrFromString(caption)
    ret, _, _ := syscall.Syscall6(procMessageBoxW.Addr(), 4, 0, uintptr(unsafe.Pointer(text_)), uintptr(unsafe.Pointer(caption_)), style, 0, 0)
    return uint32(ret)
}
