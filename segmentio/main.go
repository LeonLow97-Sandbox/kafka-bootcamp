package main

func main() {
	producer()
	go consumer()

	select{}
}