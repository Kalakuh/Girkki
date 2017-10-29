package main

import "sync"

const NICK string = "Quha"
const SERVER string = "chat.freenode.net:8000"
const LOG_PATH string = "girkki.log" 

/**
TODO: Make sure that client actually restarts after an error
TODO: Add an goroutine to each client that checks if the connection has timed out
*/

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