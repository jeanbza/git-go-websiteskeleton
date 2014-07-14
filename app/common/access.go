// Note: inspiration for this from https://gist.github.com/cespare/3985516
package common

import (
    "time"
    "strings"
    "strconv"
    "net/http"

    "github.com/golang/glog"
)

const (
    logFile = "access_log.txt"
)

type accessLog struct {
    ip, method, uri, protocol, host     string
    status                              int
    elapsedTime                         time.Duration
}

func LogAccess(w http.ResponseWriter, req *http.Request, duration time.Duration, status int) {
    clientIP := req.RemoteAddr

    if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
        clientIP = clientIP[:colon]
    }

    record := &accessLog{
        ip:             clientIP,
        method:         req.Method,
        uri:            req.RequestURI,
        protocol:       req.Proto,
        host:           req.Host,
        status:         status,
        elapsedTime:    duration,
    }

    writeAccessLog(record)
}

func writeAccessLog(record *accessLog) {
    logRecord := ""+record.ip+" "+record.protocol+" "+record.method+": "+record.uri+", host: "+record.host+" (load time: "+strconv.FormatFloat(record.elapsedTime.Seconds(), 'f', 5, 64)+" seconds)"
    glog.Infoln(logRecord)
}