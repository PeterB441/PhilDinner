package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var fork1 = make(chan bool)
	var fork2 = make(chan bool)
	var fork3 = make(chan bool)
	var fork4 = make(chan bool)
	var fork5 = make(chan bool)

	go philos(1000*time.Millisecond, fork1, fork2)
	go philos(1000*time.Millisecond, fork2, fork3)
	go philos(1000*time.Millisecond, fork3, fork4)
	go philos(1000*time.Millisecond, fork4, fork5)
	go philos(1000*time.Millisecond, fork5, fork1)
	go fork(fork1)
	go fork(fork2)
	go fork(fork3)
	go fork(fork4)
	go fork(fork5)
	for {

	}
}

func philos(delay time.Duration, forkL chan bool, forkR chan bool) {
	for {
		if(rand.Intn(2) == 0){
			// Decided to try to eat
		}else{
			// Decided to think
			fmt.Println("This philosopher is thinking...")
		}
		time.Sleep(delay)
	}
}

func fork(fork chan bool) {
	chan bool channel = fork
	for {

	}
}
