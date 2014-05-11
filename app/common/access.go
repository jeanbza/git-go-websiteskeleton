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
    ip                    string
    time                  time.Time
    method, uri, protocol string
    status                int
    responseBytes         int64
    elapsedTime           time.Duration
}

func LogAccess(w http.ResponseWriter, req *http.Request, duration time.Duration) {
    clientIP := req.RemoteAddr

    if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
        clientIP = clientIP[:colon]
    }

    record := &accessLog{
        ip:             clientIP,
        time:           time.Time{},
        method:         req.Method,
        uri:            req.RequestURI,
        protocol:       req.Proto,
        status:         http.StatusOK,
        elapsedTime:    duration,
    }

    writeAccessLog(record)
}

func writeAccessLog(record *accessLog) {
    logRecord := "["+record.time.Format("02/Jan/2006 03:04:05")+"] "+record.ip+" "+record.protocol+" "+record.method+": "+record.uri+" (load time: "+strconv.FormatFloat(record.elapsedTime.Seconds(), 'f', 5, 64)+" seconds)"
    glog.Infoln(logRecord)
}