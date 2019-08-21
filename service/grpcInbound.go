package service

import (
	"fmt"
	"time"
)

// mimic the update service
type Grpc struct {
	Name        string
	outChannel  chan<- string
	quitChannel <-chan string
}

// fully satisfy the service
func (g *Grpc) SetChannels(out chan<- string, quit <-chan string) {
	g.outChannel = out
	g.quitChannel = quit
}

func (g Grpc) GetOutChannel() chan<- string {
	return g.outChannel
}

func (g *Grpc) SetOutChannel(out chan<- string) {
	g.outChannel = out
}

// this mimics the updater
func (g *Grpc) Start(timesToWrite int) {

	fmt.Println("----- starting grpc service ------")
Q:
	for i := 0; i < 10; i++ {
		select {

		case <-g.quitChannel:
			break Q

		default:
			message := fmt.Sprintf("grpc %d", i)
			g.outChannel <- message
			time.Sleep(time.Second * 1)
		}

	}

	fmt.Println("leaving grpc")
}

func (g Grpc) GetName() string {
	return g.Name
}
