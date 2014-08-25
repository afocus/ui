// 18 august 2014

package ui

import (
	"unsafe"
)

// #include "winapi_windows.h"
import "C"

func openFile() string {
	name := C.openFile()
	if name == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(name))
	return wstrToString(name)
}
