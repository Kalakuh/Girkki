package main

import (
	"fmt"
	"net"
	"strings"
	"bufio"
)

type Client struct {
	nick string
	channels []string
	conn net.Conn
	logger *Logger
}

func send(conn net.Conn, msg string) {
	fmt.Fprintf(conn, msg + "\r\n")
	fmt.Printf("-> %v\n", msg)
}

func (cli *Client) ChangeNick(nick string) { // TODO: add error handling for unavailable nicks
	cli.nick = nick
	send(cli.conn, "NICK " + nick)
}

func (cli *Client) Auth(text string) {
	send(cli.conn, "USER " + text)
}

func (cli *Client) Pong(target string) {
	send(cli.conn, "PONG " + target)
}

func (cli *Client) Run() {
	reader := bufio.NewReader(cli.conn)
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
			cli.Pong(result[5:])
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
					cli.logger.Log(receiver + " " + nick + " " + msg)
				}
			}
		}
	}	
}