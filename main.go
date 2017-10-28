package main

import "sync"
import "fmt"

const NICK string = "Quha"
const SERVER string = "chat.freenode.net:8000"
const LOG_PATH string = "girkki.log" 

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	var logger Logger
	logger.Open(LOG_PATH)
	defer logger.Close()

	client := Client{
		NICK,
		[]string{},
		nil,
		&logger,
		false,
	}

	go func() {
		defer wg.Done()
		for !client.exit {
			client.Connect(SERVER)
			client.Run()
		}
	}()
	wg.Wait()
}