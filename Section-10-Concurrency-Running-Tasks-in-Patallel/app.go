package main

// Concurrency and Goroutines
import (
	"fmt"
	"time"
)

// func greet(phrase string) {
// 	fmt.Println("Hello!", phrase)
// }

// func slowGreet(phrase string) {
// 	time.Sleep(3 * time.Second) // Simulating a slow long-running task
// 	fmt.Println("Hello!", phrase)
// }

// normal function calls without concurrency
// func main() {
// 	greet("nice to meet you")
// 	greet("how are you")
// 	slowGreet("how...are...you...")
// 	greet("i hope you are liking the course")
// }

// // Concurrency with Goroutines
// func main() {
// 	go greet("nice to meet you")
// 	go greet("how are you")
// 	go slowGreet("how...are...you...")
// 	go greet("i hope you are liking the course")
// }

// func slowGreet(phrase string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second) // Simulating a slow long-running task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

// // Concurrency with Goroutines
// func main() {
// 	// go greet("nice to meet you")
// 	// go greet("how are you")
// 	done := make(chan bool)
// 	go slowGreet("how...are...you...", done)
// 	// go greet("i hope you are liking the course")
// 	fmt.Println(<-done)
// }

// Adding channels to all goroutines ----------------------------
func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // Simulating a slow long-running task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

// Concurrency with Goroutines
func main() {
	dones := make([]chan bool, 4)
	// done := make(chan bool)

	dones[0] = make(chan bool)
	go greet("nice to meet you", dones[0])

	dones[1] = make(chan bool)
	go greet("how are you", dones[1])

	dones[2] = make(chan bool)
	go slowGreet("how...are...you...", dones[2])

	dones[3] = make(chan bool)
	go greet("i hope you are liking the course", dones[3])

	for _, doneChan := range dones {
		<-doneChan
	}
}
