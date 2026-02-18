package main

//Go has powerful tools for concurrent programming that make it easy to work with multiple simultaneous threads. Most imp is Go routines.
//A go routine is a light weight thread that you can launch with very simple syntax. go routines run in the same memory space as an application's main thread.
//they dont block the main thread, and they are very easy to create and manage.
//user go routines whenever you have to execute a tasl that would otherwise block the main thread: working with local files, making web services api requests.

import (
	"fmt"
	"time"
)
//to run a function as a go routine, call it by starting with the keyword "go"
//But anonymous function has to be called just after declaration by using paranthesis and passing the value. it isnt reusable and can work if you only need it once.
func main(){
	go say("Hello from go routines")
	fmt.Println("Hello from main")

	//anony go routine
	go func(message string){
		fmt.Println(message)
	}("Hello from the anonymous function")

	/*if we run the code like this it wont work, i will only get output from main htread, because i have told the function to wait.
	but the main function keeps on working. and when main is done executed then the whole application stops working. so we have to add a bit of delay in the main thread.
	*/
	time.Sleep(2*time.Second)
	fmt.Println("All done")
}

func say(message string){
	time.Sleep(1*time.Second)
	fmt.Println(message)
}

//you can also create anonymous go routines by creating a function without a name that immediately follows go keyword.

/*Channels in Go are a fundamental feature for communication between goroutines. They provide a way for goroutines to send and receive values, enabling safe concurrent programming.
ch := make(chan int)        // unbuffered channel
ch := make(chan int, 5)     // buffered channel with capacity 5
ch <- 42        // send value to channel
value := <-ch   // receive value from channel
Channel Types:
Unbuffered: Synchronous - sender blocks until receiver is ready
Buffered: Asynchronous - sender only blocks when buffer is full

CHANNELS ARE QUEUES
If you try to receive more values than what's available in the channel, your program will deadlock and hang forever.
You need a channel because goroutines run independently and you need a way to collect their results back to the main function.
The channel acts like a collection box where all goroutines drop their results, and you pick them up one by one in the loop.
Goroutines can't directly return values - they need to send results somewhere
runtime.NumCPU() returns the number of logical CPU cores available to the current process.
	Returns an integer representing CPU cores (including hyperthreaded cores)
	On a 4-core CPU with hyperthreading: returns 8
	On a 6-core CPU without hyperthreading: returns 6

Why use it: Optimal parallelism: Creates one goroutine per CPU core
			Avoids oversubscription: Too many goroutines can cause context switching overhead
			Hardware adaptive: Automatically scales with different machines*/
