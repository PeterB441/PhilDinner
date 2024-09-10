package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var fork1 = make(chan bool, 1)
	var fork2 = make(chan bool, 1)
	var fork3 = make(chan bool, 1)
	var fork4 = make(chan bool, 1)
	var fork5 = make(chan bool, 1)

	go fork(fork1)
	go fork(fork2)
	go fork(fork3)
	go fork(fork4)
	go fork(fork5)
	go philos(1, 998*time.Millisecond, fork1, fork2)
	go philos(2, 999*time.Millisecond, fork2, fork3)
	go philos(3, 1000*time.Millisecond, fork3, fork4)
	go philos(4, 1001*time.Millisecond, fork4, fork5)
	go philos(5, 1002*time.Millisecond, fork5, fork1)
	for {
		/*
			var f1 = <-fork1
			fork1 <- f1
			var f2 = <-fork2
			fork2 <- f2
			var f3 = <-fork3
			fork3 <- f3
			var f4 = <-fork4
			fork4 <- f4
			var f5 = <-fork5
			fork5 <- f5
			fmt.Println(f1, f2, f3, f4, f5)
			time.Sleep(1000 * time.Millisecond)
		*/
	}
}

func philos(id int, delay time.Duration, forkL chan bool, forkR chan bool) {
	var timesEaten = 0
	for {
		// Every time the philosopher tries to take an action, a coinflip is made to determine if it will eat or think
		// This is done to prevent deadlocks, by making sure that all philosiphers will probably take different actions than their neightbours
		// And if they do all somehow try to eat at the same time, it will result in them all flipping a new coin in the next cycle
		if rand.Intn(2) == 0 && timesEaten < 5 {
			// Decided to try to eat
			var hasEaten = eat(delay, forkL, forkR)
			if hasEaten { // Has seuccesfully eaten food
				timesEaten++
				fmt.Println("Philosopher", id, "has eaten", timesEaten, "times")
			} else { // Failed to eat, due to a lack of forks
				fmt.Println("Philosopher", id, " fail to eat, an is now thinking...")
				think(delay)
			}
			if timesEaten == 5 {
				fmt.Println("Philosopher", id, "is done eating and will now only think")
			}
		} else {
			// Decided to think
			if timesEaten != 5 { // To stop it from printing it thinking, after being done eating
				fmt.Println("Philosopher", id, "is thinking...")
			}
			think(delay)
		}
	}
}

// true represents a fork that is not available to be picked up
// false represents a fork that is available to be picked up
func eat(delay time.Duration, forkL chan bool, forkR chan bool) bool {
	var pickedUp = <-forkR
	if pickedUp { // The fork on the right is already taken
		forkR <- true
		return false
	} else { //The fork on the right is now picked up
		forkR <- true
	}
	pickedUp = <-forkL
	if pickedUp { // The fork on the left is already taken
		forkL <- true
		putForkDown(forkR)
		return false
	} else { // The fork on the left is now picked up
		forkL <- true
	}
	time.Sleep(delay) // Mmmmm, eating time
	putForkDown(forkR)
	putForkDown(forkL)
	return true
}

func think(delay time.Duration) {
	time.Sleep(delay)
}

func putForkDown(fork chan bool) {
	<-fork
	fork <- false
}

func fork(fork chan bool) {
	var channel = fork
	channel <- false
	for {

	}
}
