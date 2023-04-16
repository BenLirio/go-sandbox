package main

func main() {
	ch := make(chan string, 1)

	go func () {
		<- ch
	}()
}