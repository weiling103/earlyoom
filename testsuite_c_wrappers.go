package earlyoom_testsuite

import (
	"fmt"
	"os"
)

// #cgo CFLAGS: -std=gnu99 -DCGO
// #include "meminfo.h"
// #include "kill.h"
// #include "msg.h"
import "C"

func sanitize(s string) string {
	cs := C.CString(s)
	C.sanitize(cs)
	return C.GoString(cs)
}

func parse_term_kill_tuple(optarg string, upper_limit int) (error, int, int) {
	cs := C.CString(optarg)
	tuple := C.parse_term_kill_tuple(cs, C.long(upper_limit))
	errmsg := C.GoString(&(tuple.err[0]))
	if len(errmsg) > 0 {
		return fmt.Errorf(errmsg), 0, 0
	}
	return nil, int(tuple.term), int(tuple.kill)
}

func is_alive(pid int) bool {
	res := C.is_alive(C.int(pid))
	return bool(res)
}

func fix_truncated_utf8(str string) string {
	cstr := C.CString(str)
	C.fix_truncated_utf8(cstr)
	return C.GoString(cstr)
}

func get_process_stats(pid int) C.struct_procinfo {
	oldwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = os.Chdir("/proc")
	if err != nil {
		panic(err)
	}
	defer os.Chdir(oldwd)

	return C.get_process_stats(C.int(pid))
}