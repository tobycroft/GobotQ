package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type WsClient struct {
	Conn       *websocket.Conn
	retryTime  int
	retryDelay time.Duration
	retry      bool
	err        error
	Recv       chan []byte
	Send       chan []byte
}

func (ws *WsClient) SetRetry(retry bool) *WsClient {
	ws.retry = retry
	return ws
}

func (ws *WsClient) SetRetryTime(retryTime int) *WsClient {
	ws.retryTime = retryTime
	return ws
}

func (ws *WsClient) SetRetryDelay(retryDelaySec time.Duration) *WsClient {
	ws.retryDelay = retryDelaySec * time.Second
	return ws
}

func (ws *WsClient) NewConnect(url string) error {
	if ws.retry {
		if ws.retryDelay.Seconds() < 1 {
			ws.retryDelay = 5 * time.Second
		}
	}
	if conn, _, err := websocket.DefaultDialer.Dial(url, nil); err != nil {
		return err
	} else {
		ws.Conn = conn
		ws.Recv = make(chan []byte, 1)
		ws.Send = make(chan []byte, 1)
		go ws.recv_data()
		ws.send_data()
		if ws.Conn != nil {
			ws.Conn.Close()
		}
		if ws.retry {
			time.Sleep(ws.retryDelay)
			return ws.NewConnect(url)
		}
	}
	return ws.err
}

func (ws *WsClient) recv_data() {
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.err = err
			log.Println("read:", err)
			return
		}
		ws.Recv <- message
	}
}

func (ws *WsClient) send_data() {
	for c := range ws.Send {
		err := ws.Conn.WriteMessage(websocket.TextMessage, c)
		if err != nil {
			ws.err = err
			log.Println("write:", err)
			return
		}
	}
}
