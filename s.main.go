//BUG - After the client exit, the number in center will not reduce.
//TODO - Add a onclose event in frontend.
package main

import "fmt"
import "net/http"
import "strconv"
import "github.com/gorilla/websocket"

type Center struct {
	num_onliner int
}

func newCenter() *Center {
	return new(Center)
}

//return the number of people online:
func (c *Center) GetOnliner() []byte {
	//convert c.num_online, which is int, to a byte array:
	return []byte(strconv.Itoa(c.num_onliner))
}

func (c *Center) AddOnliner() {
	c.num_onliner += 1
}

func ServeWs(upgrader websocket.Upgrader, center *Center, w http.ResponseWriter, r *http.Request) {
	//Initialization:
	//create the connection:
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	fmt.Println("/ws - accessing....")
	center.AddOnliner()
	//Polling:
	for {
		//Read message from the client:
		//code will be blocked here until it received msg:
		msg_type, msg_cx, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Fatal - conn--readmsg")
			return
		}
		fmt.Println(msg_type, msg_cx)
		//Write the messages(how many people onlines) to the client:
		err = conn.WriteMessage(websocket.TextMessage, center.GetOnliner())
		if err != nil {
			fmt.Println("Fatal - conn--response")
			return
		}
		fmt.Println("conn--response....!")
	}
}

func main() {
	fmt.Println("http://127.0.0.1:9999")
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	var center = newCenter()
	//To serve the webpages to the client:
	http.Handle("/", http.FileServer(http.Dir(".")))
	//To handle the websocket request:
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(upgrader, center, w, r)
	})
	http.ListenAndServe(":9999", nil)
}
