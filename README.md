git-go-websiteskeleton
===============

A basic website skeleton in Go (or golang, if you prefer) that comes with the Gorilla Multiplexer, bootstrap, and jQuery.

### You Will Need ###
1. A go environment http://golang.org/doc/install#install
2. Gorilla mux: go get github.com/gorilla/mux (http://www.gorillatoolkit.org/pkg/mux)
3. glog: go get github.com/golang/glog
4. Some basic knowledge of Go's httpd package. See the excellent gowiki tutorial at http://golang.org/doc/articles/wiki/

### Installation ###
1. cd $GOPATH/src
2. git clone https://github.com/jadekler/git-go-websiteskeleton.git
3. cd git-go-websiteskeleton
4. go build main.go
5. ./main
6. Navigate to http://localhost:8080

### Access Logging ###
Access logging captures each request serve, logging the requestor's IP, the request method, uri, and protocol, the time
requested, and the page load time. Logs are written to $TMP. If you leave $TMP unset, logs will be written to a temporary
folder.

To set $TMP in a Unix environment, simply type export TMP=/path/to/dir or edit your bash profile and add the aforementioned
line.