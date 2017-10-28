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

func analyzeCommand(client *Client, cmd string) {
	if strings.Index(cmd, "!join") == 0 {
		client.Join(cmd[6:])
	} else if strings.Index(cmd, "!nick") == 0 {
		client.ChangeNick(cmd[6:])
	} else if strings.Index(cmd, "!msg") == 0 {
		cmd := cmd[5:]
		split := strings.Split(cmd, " ")
		target := split[0]
		msg := cmd[len(target) + 1:]
		fmt.Printf("DEBUG: %s ~ %s\n", target, msg)
		client.Privmsg(target, msg)
	}
}

func (cli *Client) Connect(network string) {
	var err error
	cli.conn, err = net.Dial("tcp", SERVER)
	if err != nil {
		panic(err)
	}
	cli.ChangeNick(NICK)
	cli.Auth(NICK + " 8 * :Kuha")
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

func (cli *Client) Join(channel string) {
	send(cli.conn, "JOIN " + channel)
}

func (cli *Client) Privmsg(channel string, msg string) {
	send(cli.conn, "PRIVMSG " + channel + " :" + msg)
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
					analyzeCommand(cli, msg)
				}
			}
		}
	}	
}