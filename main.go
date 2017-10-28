package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

const NICK string = "quha"
const SERVER string = "chat.freenode.net:8000"
const CHANNEL string = "#datatahti"

const LOG_PATH string = "girkki.log" 

func main() {
	conn, dialErr := net.Dial("tcp", SERVER)
	if dialErr != nil {
		panic(dialErr)
	}

	client := Client{
		NICK,
		[]string{CHANNEL},
		conn,
	}

	var logger Logger
	logger.Open(LOG_PATH)
	defer logger.Close()

	
	//send(conn, "NICK " + NICK)
	client.ChangeNick(NICK)
	send(conn, "USER " + NICK + " 8 * :Kuha")
	reader := bufio.NewReader(conn)
	for {
		result, err := reader.ReadString('\n')
		for result[len(result) - 1] == '\r' || result[len(result) - 1] == '\n' {
			result = result[:len(result) - 1]
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("<- %v\n", result)
		if result[:4] == "PING" {
			client.Pong(result[5:])
			//send(conn, "PONG " + result[5:len(result) - 1])
		} else {
			spl := strings.Split(result, " ")
			command := spl[1]
			if command == "PRIVMSG" { // We're not currently interested in commands other than PRIVMSGs
				sender := strings.Split(spl[0], "!")
				nick := sender[0][1:]
				receiver := spl[2]
				indexOfColon := strings.Index(result[1:], ":")
				if indexOfColon != -1 {
					msg := result[indexOfColon + 2:]
					logger.Log(SERVER + " " + receiver + " " + nick + " " + msg)
				}
			}
		}
	}
}