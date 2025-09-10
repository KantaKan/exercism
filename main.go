package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)



func openConnection(done chan bool){
	fmt.Println("Attempting connection ....")

	if rand.Intn(100) > 50 {
		fmt.Println("Connection Failed ")
		time.Sleep(1000000 * time.Hour)
	}else{
		fmt.Println("Connection Established")
	}
	done <- true
}

func openConnectionWithTimeout(){
	ctx , cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan bool)
	go openConnection(done)
	select {
	case <- done:
		fmt.Println("Operation Completed")
	case <- ctx.Done():
		fmt.Println("Operation Timed Out")
	}
}


func main(){
	openConnectionWithTimeout()
}