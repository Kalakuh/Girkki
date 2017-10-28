package main

const NICK string = "quha"
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

	client.Connect(SERVER)
	client.ChangeNick(NICK)
	client.Auth(NICK + " 8 * :Kuha")
	client.Run()
}