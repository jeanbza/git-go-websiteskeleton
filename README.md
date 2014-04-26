git-go-websiteskeleton
===============

A basic website skeleton in Go (or golang, if you prefer) that comes with the Gorilla Multiplexer, bootstrap, and jQuery.

### You Will Need ###
1. A go environment http://golang.org/doc/install#install
2. Gorilla mux: go get github.com/gorilla/mux (http://www.gorillatoolkit.org/pkg/mux)
3. Some basic knowledge of Go's httpd package. See the excellent gowiki tutorial at http://golang.org/doc/articles/wiki/

### Installation ###
1. cd $GOPATH/src
2. git clone https://github.com/jadekler/git-go-websiteskeleton.git
3. cd git-go-websiteskeleton
4. go build main.go
5. ./main
6. Navigate to http://localhost:8080

### Access Logging ###
Access logging is turned on by default, and will (also) by default write to git-go-websiteskeleton/access_log.txt. You can turn access logging on/off in main.go. You can change the location of this log at app/common/access_log.go.
