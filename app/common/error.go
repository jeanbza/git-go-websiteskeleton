package common

import (
    "log"
    "runtime"

    "github.com/golang/glog"
)

// 0 = info
// 1 = warning
// 2 = error - should be most common
// 3 = fatal
func CheckError(err error, level int) {
    if err != nil {
        var stack [4096]byte
        runtime.Stack(stack[:], false)
        log.Printf("%q\n%s\n", err, stack[:])

        switch level {
        case 0:
            glog.Infoln("%q\n%s\n", err)
        case 1:
            glog.Warningln("%q\n%s\n", err)
        case 2:
            glog.Errorln("%q\n%s\n", err)
        case 3:
            glog.Fatalln("%q\n%s\n", err)
        }

        glog.Flush()
    }
}