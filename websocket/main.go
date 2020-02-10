package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

var rtmWSUrl = ""

func main() {
	ws, err := websocket.Dial(rtmWSUrl, "", "https://api.slack.com/")
	if err != nil {
		panic(err)
	}

	var event interface{}
	for {
		if err := websocket.JSON.Receive(ws, &event); err != nil {
			panic(err)
		}

		fmt.Println(event)
	}
}
