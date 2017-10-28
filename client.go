package main

import (
	"fmt"
	"net"
)

type Client struct {
	nick string
	channels []string
	conn net.Conn
}

func send(conn net.Conn, msg string) {
	fmt.Fprintf(conn, msg + "\r\n")
	fmt.Printf("-> %v\n", msg)
}

func (cli *Client) ChangeNick(nick string) { // TODO: add error handling for unavailable nicks
	cli.nick = nick
	send(cli.conn, "NICK " + nick)
}

func (cli *Client) Pong(target string) {
	send(cli.conn, "PONG " + target)
}