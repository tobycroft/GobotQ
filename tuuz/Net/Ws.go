package Net

import "github.com/gorilla/websocket"

type WebsocketClient struct {
	Conn      *websocket.Conn
	retryTime int
	retry     bool
	Recv      chan any
	Send      chan any
}

func (s *WebsocketClient) SetRetry(retry bool) *WebsocketClient {
	s.retry = retry
	return s
}

func (s *WebsocketClient) SetRetryTime(retryTime int) *WebsocketClient {
	s.retryTime = retryTime
	return s
}

func (ws *WebsocketClient) NewWebSocketClient(url string) *WebsocketClient {
	if conn, _, err := websocket.DefaultDialer.Dial(url, nil); err != nil {
		panic(err.Error())
	} else {
		ws.Conn = conn
	}
	return ws
}
