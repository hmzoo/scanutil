
package webserver

import (
	"bytes"
	"log"
	"net/http"
	"time"
  "github.com/google/uuid"
	"github.com/gorilla/websocket"
  "encoding/json"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ItemMsg struct {
  Id string `json:"id"`
  Tp string `json:"tp"`
  Item Item `json:"item"`
}

type DataMsg struct {
  Id string `json:"id"`
  Tp string `json:"tp"`
  Data *Data `json:"data"`
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

  id string

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	data *Data
}

func (c *Client) process(databytes []byte) {
  msg:= ItemMsg{}
  err := json.Unmarshal(databytes, &msg)
  if(err != nil) {
    return
  }
  msg.Id=c.id
  log.Println(msg)

 if(msg.Tp=="ADD"){
	 c.data.AddItem(msg.Item)
	 c.data.SaveCSV()
 }

  b, err := json.Marshal(msg);
  if err!=nil {
  log.Println(err)
  return
 }

  c.hub.broadcast <- b
}

func (c *Client) sendData() {
	msg := DataMsg{Id :c.id,Tp:"DATA",Data:c.data}
	b, err := json.Marshal(msg);

	if err!=nil {
	log.Println(err)
	return
 }

	if err := c.conn.WriteMessage(websocket.TextMessage, b); err != nil {
        log.Println(err)
        return
    }



}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

    c.process(message)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
      /*
			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
      */
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(data *Data,hub *Hub, w http.ResponseWriter, r *http.Request) {
  upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{ data: data ,hub: hub, id:NewID(), conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	client.sendData()

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func NewID() string{
  u := uuid.New()
  return u.String()
}
