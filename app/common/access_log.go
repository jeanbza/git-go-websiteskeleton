package common

import (
    "os"
    "io"
    "time"
    "strings"
    "strconv"
    "net/http"
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
    logRecord := "["+record.time.Format("02/Jan/2006 03:04:05")+"] "+record.ip+" "+record.protocol+" "+record.method+": "+record.uri+" (load time: "+strconv.FormatFloat(record.elapsedTime.Seconds(), 'f', 5, 64)+" seconds)\n"

    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
    CheckError(err)
    io.WriteString(file, logRecord)
    file.Close()
}