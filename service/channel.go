package service

import (
	"fmt"
)

// this mimics the core reader
func CoreReader(reader <-chan string, quit <-chan interface{}) {

	fmt.Println("------- reading real channel ---------")
Z:
	for {
		select {

		case code := <-reader:
			fmt.Printf("coreReader: %s\n", code)

		case <-quit:
			break Z
		}
	}

	fmt.Println("leaving CoreReader")
}

// this mimics the core reader
func devNullReader(reader <-chan string, quit <-chan interface{}) {

	fmt.Println("------- reading fake channel ---------")
Z:
	for {
		select {

		case code := <-reader:
			fmt.Printf("devNullReader: %s\n", code)

		case <-quit:
			break Z

		}
	}

	fmt.Println("leaving devNullReader")
}
