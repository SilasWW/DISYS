package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var activePhilosophers sync.WaitGroup

func philosopher(philName string, strongHand chan string, strongHandReturn chan string, weakHand chan string, weakHandReturn chan string) {

	//random amount of time slept
	var sleepTime = rand.Intn(75)

	// philosopher eats three times
	for eaten := 0; eaten < 3; {

		//sends pickup "request" to mainhand fork
		strongHand <- "pickup"
		x := <-strongHandReturn

		//sends pickup "request" to weakhand fork
		weakHand <- "pickup"
		y := <-weakHandReturn

		//checks that both forks accept the pickup
		if x == "accept" && y == "accept" {
			fmt.Printf("%s is now eating \n", philName)
			time.Sleep(time.Duration(sleepTime))

			eaten++
		}

		//throws both forks away
		strongHand <- "throw"
		weakHand <- "throw"

		fmt.Printf("%s is now thinking \n", philName)
		time.Sleep(time.Duration(sleepTime))
	}

	// update activePhilosopher to keep track of who is done eating
	activePhilosophers.Done()
	fmt.Printf("%s is now thinking indefinitely \n", philName)
}

func fork(name string, sendA chan string, receiveA chan string, sendB chan string, receiveB chan string) {
	//while loop makes forks available until end of program run
	for true {
		//select takes first message from either chan b or chan a
		select {
		//checks that first message is pickup, sends accept message back to philosopher,
		//waits and checks that last message is throw
		case ARes := <-receiveA:
			if ARes == "pickup" {
				sendA <- "accept"

				ARes = <-receiveA

				if ARes != "throw" {
					fmt.Printf("Fork: %s Received: %s Instead of: throw \n", name, ARes)
				}
			} else {
				fmt.Printf("Fork: %s Received %s Instead of: pickup \n", name, ARes)
			}

		//same as case a but for b channels
		case BRes := <-receiveB:
			if BRes == "pickup" {
				sendB <- "accept"

				BRes = <-receiveB

				if BRes != "throw" {
					fmt.Printf("Fork: %s Received: %s Instead of: throw \n", name, BRes)
				}
			} else {
				fmt.Printf("Fork: %s Received %s Instead of: pickup \n", name, BRes)
			}
		}
	}
}

func main() {
	activePhilosophers.Add(5)

	//make channels
	var edge1 = make(chan string)
	var edge2 = make(chan string)
	var edge3 = make(chan string)
	var edge4 = make(chan string)
	var edge5 = make(chan string)
	var edge6 = make(chan string)
	var edge7 = make(chan string)
	var edge8 = make(chan string)
	var edge9 = make(chan string)
	var edge10 = make(chan string)

	//make returns
	var return1 = make(chan string)
	var return2 = make(chan string)
	var return3 = make(chan string)
	var return4 = make(chan string)
	var return5 = make(chan string)
	var return6 = make(chan string)
	var return7 = make(chan string)
	var return8 = make(chan string)
	var return9 = make(chan string)
	var return10 = make(chan string)

	//make forks
	go fork("fork1", return1, edge1, return2, edge2)
	go fork("fork2", return3, edge3, return4, edge4)
	go fork("fork3", return5, edge5, return6, edge6)
	go fork("fork4", return7, edge7, return8, edge8)
	go fork("fork5", return9, edge9, return10, edge10)

	// start go routine for each philosopher. The selection of each philosophers forks is chosen so their main
	// hand faces another philosopher as often as possible. This helps to break any loops that could be formed.
	// may not actually be beneficial

	go philosopher("ph1", edge10, return10, edge1, return1)
	go philosopher("ph2", edge3, return3, edge2, return2)
	go philosopher("ph3", edge4, return4, edge5, return5)
	go philosopher("ph4", edge7, return7, edge6, return6)
	go philosopher("ph5", edge8, return8, edge9, return9)
	// waits for everyone to finish
	activePhilosophers.Wait()
	fmt.Println("All Philosophers have finished eating!")
}
