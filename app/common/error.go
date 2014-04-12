package common

import (
    "log"
    "runtime"
)

func CheckError(err error) {
    if err != nil {
        var stack [4096]byte
        runtime.Stack(stack[:], false)
        log.Printf("%q\n%s\n", err, stack[:])
    }
}