package main

import "log"

func main() {

	TestAMPQ()
	forever := make(chan bool)
	log.Println("[*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
