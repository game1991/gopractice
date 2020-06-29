package phpbbencrypt

//#include <stdlib.h>
//#include "dencrypt.h"
import "C"

import (
	"unsafe"
)

func Encrypt(password string) string {
	cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(cpassword))
	arr := string(make([]byte, 1024))
	carr := C.CString(arr)
//	defer C.free(unsafe.Pointer(carr))
	C.encrypt(carr, cpassword)
	return C.GoString(carr)
}

func Verify(password string, hash string) bool {
	cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(cpassword))
	chash := C.CString(hash)
	defer C.free(unsafe.Pointer(chash))
	ret := C.verify(cpassword, chash)
	return 1 == ret
}
