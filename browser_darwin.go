package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/sqweek/dialog"
)

type browser struct {
	keepAlive bool
	url       string
}

var urlCh chan browser

func init() {
	urlCh = make(chan browser, 1)
}
func openUI(keepAlive bool, url string) (func(), error) {
	urlCh <- browser{
		keepAlive: keepAlive,
		url:       url,
	}
	return func() {}, nil
}
func errorMessage(message string) {
	dialog.Message("error[%s]", message).Title("Error")
}
func main() {
	startDarwin()
}

//export startDarwin
func startDarwin() unsafe.Pointer {
	go start()
	v := <-urlCh
	l := len(v.url)
	addr := C.malloc(C.ulong(l + 1))
	C.memset(addr, 0, C.ulong(l+1))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&v.url))
	C.memcpy(addr, unsafe.Pointer(sh.Data), C.ulong(sh.Len))
	return addr
}

//export freePointer
func freePointer(p unsafe.Pointer) {
	C.free(p)
}
