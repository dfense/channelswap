package service

import (
	"fmt"
	"time"
)

// mimic the update service
type Updater struct {
	Name        string
	outChannel  chan<- string
	quitChannel <-chan string
}

// fully satisfy the service
func (up *Updater) SetChannels(out chan<- string, quit <-chan string) {
	up.outChannel = out
	up.quitChannel = quit
}

func (up Updater) GetOutChannel() chan<- string {
	return up.outChannel
}

func (up *Updater) SetOutChannel(out chan<- string) {
	up.outChannel = out
}

// this mimics the updater
// create a devnull channel reader, and then assign that to the current grpc service.
// steal the existnig grpc service channel, and write all the updater messages.
// once update is done, politely replace the original grpc channel, and it is happy again.
// cleanup by killing the dev null reader and deleting it's channel
func (up *Updater) Start(timesToWrite int) {

	// create a swap channel for grpc
	devNullChannel := make(chan string)
	quitdevNullReader := make(chan interface{})
	go devNullReader(devNullChannel, quitdevNullReader)

	up.outChannel = devNullChannel
	// swap channels with grpc
	SwitchChannel(up.GetName(), "grpc")

Q:
	for i := 0; i < timesToWrite; i++ {
		select {

		case <-up.quitChannel:

			break Q

		default:
			message := fmt.Sprintf("updater %d", i)
			up.outChannel <- message
			time.Sleep(time.Second * 1)
		}

	}

	// flip channels back
	SwitchChannel(up.GetName(), "grpc")

	// kill the devNullReader
	if quitdevNullReader != nil {
		quitdevNullReader <- "quit"
		close(devNullChannel)
		close(quitdevNullReader)
	}

	fmt.Println("leaving Updater")
}

func (up Updater) GetName() string {
	return up.Name
}
