package merkle



/*
#cgo LDFLAGS: -L${SRCDIR}/shared -lrustlib
#include <stdint.h>

extern uint64_t PoseidonHashGo(data1 *const u8, len1 usize, data2 *const u8, len2 usize, out_len *mut usize);
*/


import (
	"C"
	"fmt"
	"unsafe"
)


func PoseidonHashGo(input []byte) {
	
	ptr := (*C.uint8_t)(unsafe.Pointer(&input[0]))
	length := C.size_t(len(input))
	result := C.PoseidonHash(ptr, length)
	fmt.Println("Rust returned: ", uint64(result))


}
