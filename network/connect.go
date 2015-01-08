package network

import (
	"github.com/fhbzyc/c_game/libs/log"
	"github.com/fhbzyc/c_game/protocol"
	"github.com/gorilla/websocket"
)

const ()

type Connect struct {
	Uid    int
	AreaId int
	RoleId int
	Conn   *websocket.Conn
	Chan   chan []byte
}

func (c *Connect) Send(s []byte) error {
	c.Chan <- s
	return nil
}

func (c *Connect) Write() {
	c.Chan = make(chan []byte, 10)
	go func() {
		for s := range c.Chan {
			if err := c.Conn.WriteMessage(websocket.TextMessage, s); err != nil {
				log.Logger.Warn("Send Message Error: ", err)
			} else {
				log.Logger.Info("Send Success")
			}
		}
	}()
}

func (c *Connect) Close() {

	playerMap.Delete(c.AreaId, c.RoleId)
	c.Conn.Close()
	close(c.Chan)
}

func (c *Connect) InMap() {
	conn := playerMap.Get(c.AreaId, c.RoleId)
	if conn != nil {
		errMsg := new(protocol.Error)
		errMsg.Error.Code = protocol.ERROR_TOKEN
		errMsg.Error.Message = "连接失效,请重连"
		conn.Conn.WriteMessage(websocket.TextMessage, protocol.MarshalError(errMsg))
		conn.Conn.Close()
	}
	playerMap.Set(c.AreaId, c.RoleId, c)
}
