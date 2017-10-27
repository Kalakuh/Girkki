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

func send(conn net.Conn, msg string) {
	fmt.Fprintf(conn, msg + "\r\n")
	fmt.Printf("-> %v\n", msg)
}

func main() {
	conn, dialErr := net.Dial("tcp", SERVER)
	if dialErr != nil {
		fmt.Println(dialErr)
		return
	}

	var logger Logger
	logger.Open(LOG_PATH)
	defer logger.Close()

	send(conn, "NICK " + NICK)
	send(conn, "USER " + NICK + " 8 * :Kuha")
	reader := bufio.NewReader(conn)
	for {
		result, err := reader.ReadString('\n')
		for result[len(result) - 1] == '\r' || result[len(result) - 1] == '\n' {
			result = result[:len(result) - 1]
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("<- %v", result)
		if result[:4] == "PING" {
			send(conn, "PONG " + result[5:len(result) - 1])
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