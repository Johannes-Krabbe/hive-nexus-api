package chat

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type StartChatRequestBody struct{}

func (h handler) StartChat(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		fmt.Println("Could not upgrade connection")
		fmt.Println(err.Error())
	}

	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				fmt.Println("Could not read client data")
				fmt.Println(err.Error())
				// conn.Close()
				break
			}
			fmt.Println(">>>", string(msg))
			var newMsg string = "message from server, hello"
			err = wsutil.WriteServerMessage(conn, op, []byte(newMsg))
			if err != nil {
				fmt.Println("Could not send message to client")
				fmt.Println(err.Error())
			}
		}
	}()
}
