git-go-websiteskeleton
===============

A basic website skeleton in Go (or golang, if you prefer) that comes with the Gorilla Multiplexer for routing, glog for access and error logging, as well as jquery, bootstrap, and font-awesome for making excellent web apps.

### You Will Need ###
1. A go environment http://golang.org/doc/install#install
2. Gorilla mux: go get github.com/gorilla/mux (http://www.gorillatoolkit.org/pkg/mux)
3. glog: go get github.com/golang/glog
4. httpscerts: github.com/kabukky/httpscerts
5. secureheader: go get github.com/kr/secureheader
6. Some basic knowledge of Go's httpd package. See the excellent gowiki tutorial at http://golang.org/doc/articles/wiki/

### Installation ###
1. cd $GOPATH/src
2. git clone https://github.com/jadekler/git-go-websiteskeleton.git
3. cd git-go-websiteskeleton
4. go run main.go
5. Navigate to https://localhost:8080

Note: To make it your own project, simply remove the .git file from the project and upload it into your own repository.

### Logging ###
This project contains access and error logging. By default, these logs will be placed in your system's temp folder (on most *nix machines, this is /tmp; on mac, it's /private/var/gobbledegook). To change the log dir, simply run main with -log_dir="/path/to/dir", as in ./main -log_dir="/path/to/dir". If you decide to change the log location (which I recommend), you are in charge of cleaning your log dir periodically. The program's log library (glog) will handle max file sizes, but multiple logs will not be cleaned automatically.

Some more info on what gets captured:
- Access logging captures each request served, logging the requestor's IP, the request method, uri, and protocol, the time
requested, and the page load time.
- Error logging captures the time, error, and a stack trace.

Note: on *nix systems, you can also set the TMPDIR environment variable to specify default output location.

## TLS connection ###
Currently this skeleton is created for TLS connection and using auto-created certificate. Just create/replace key.pem and cert.pem  with your own if you want to use real certificate. These files are not ander Git control.
