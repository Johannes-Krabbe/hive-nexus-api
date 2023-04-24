package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	var wg sync.WaitGroup

	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://localhost:8080/chat/start")
	if err != nil {
		fmt.Println("failed to open connection to server")
		fmt.Println(err.Error())
	}
	wg.Add(1)

	go func() {
		defer conn.Close()
		defer wg.Done()
		var msg string = "hello"
		wsutil.WriteClientMessage(conn, ws.OpText, []byte(msg))

		for {

		}

	}()

	wg.Wait()
}
