package main

import (
	"os"
	"os/signal"
	"time"
)

func main() {
	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt)

	go func () {
		<- ch
		go func () {
			for {
				<- ch
				println("2")
			}
		}()
		println("1")
	}()

	println("Before sleep")
	time.Sleep(5 * time.Second)
	println("After sleep")


}