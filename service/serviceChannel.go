package service

import "fmt"

var (
	serviceNames = make(map[string]Service)
)

type Service interface {
	SetChannels(chan<- string, <-chan string)
	GetOutChannel() chan<- string
	SetOutChannel(chan<- string)
	GetName() string
	Start(int)
}

// register the new service, and start it with number of times to go thru loop
func RegisterService(service Service, runloop int) {
	fmt.Printf("REGISTERING -- [%s]\n", service.GetName())
	serviceNames[service.GetName()] = service
	go service.Start(runloop)
}

// flip channels between both services
func SwitchChannel(service1, service2 string) {

	s1 := serviceNames[service1]
	s2 := serviceNames[service2]

	tmpChannel := s1.GetOutChannel()
	s1.SetOutChannel(s2.GetOutChannel())
	s2.SetOutChannel(tmpChannel)

}
