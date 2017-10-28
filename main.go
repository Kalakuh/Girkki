package main

const NICK string = "Quha"
const SERVER string = "chat.freenode.net:8000"
const LOG_PATH string = "girkki.log" 

func main() {
	var logger Logger
	logger.Open(LOG_PATH)
	defer logger.Close()

	client := Client{
		NICK,
		[]string{},
		nil,
		&logger,
	}

	go func() {
		client.Connect(SERVER)
		client.Run()
	}()

	for {}
}