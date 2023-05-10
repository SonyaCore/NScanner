package main

import (
	"Scanner/pkg/logger"
	"Scanner/pkg/net"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"strings"
)

var PORT = ":8000"

func VersionStatement() string {
	return strings.Join([]string{
		"Port Scanner Service", " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
	}, "")
}

func main() {
	logger.Init()
	fmt.Println(VersionStatement())

	net.RegisterRoutes()
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
