package main

import (
	"net"
)

const NICK string = "quha"
const SERVER string = "chat.freenode.net:8000"

const LOG_PATH string = "girkki.log" 

func main() {
	var logger Logger
	logger.Open(LOG_PATH)
	defer logger.Close()

	conn, dialErr := net.Dial("tcp", SERVER)
	if dialErr != nil {
		panic(dialErr)
	}

	client := Client{
		NICK,
		[]string{},
		conn,
		&logger,
	}

	client.ChangeNick(NICK)
	client.Auth(NICK + " 8 * :Kuha")
	client.Run()
}