package main

import (
	"fmt"
	"time"

	"github.com/dfense/channelswap/service"
)

// small demo to show channel swap out
func main() {

	realChannel := make(chan string)
	fmt.Printf("Pointer1: %p\n", realChannel)
	quitUpdater := make(chan string)
	quitGrpc := make(chan string)
	quitCoreReader := make(chan interface{})

	// start reader
	go service.CoreReader(realChannel, quitCoreReader)

	// grpc service
	grpc := &service.Grpc{Name: "grpc"}
	grpc.SetChannels(realChannel, quitGrpc)
	service.RegisterService(grpc, 10)

	// updater service
	updater := &service.Updater{Name: "updater"}
	updater.SetChannels(nil, quitUpdater)
	service.RegisterService(updater, 3)

	time.Sleep(time.Second * 5)

	quitGrpc <- "stop"
	time.Sleep(time.Second * 1)
	quitCoreReader <- "stop"
	time.Sleep(time.Second * 1)

	fmt.Println("----- leaving program ---------")
}
