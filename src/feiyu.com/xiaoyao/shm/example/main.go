package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_new(char *name) {
    shm_unlink(name);
    return shm_open(name, O_RDWR|O_CREAT|O_EXCL, FILE_MODE);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	SHM_NAME = "myshm"
	SHM_SIZE = 4 * 1024 * 1024
)

type MyData struct {
	Col1, Col2, Col3 int
}

func main() {
	_, err := C.my_shm_new(C.CString(SHM_NAME))
	if err != nil {
		fmt.Printf("C.my_shm_new failed : %v", err)
	}

	return
}
